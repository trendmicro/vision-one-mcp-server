package tools

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

type paramKind int

const (
	kindString paramKind = iota
	kindNumber
	kindBoolean
	kindArray
	kindObject
)

func (k paramKind) String() string {
	switch k {
	case kindNumber:
		return "number"
	case kindBoolean:
		return "boolean"
	case kindArray:
		return "array"
	case kindObject:
		return "object"
	default:
		return "string"
	}
}

type bodyKind int

const (
	bodyNone bodyKind = iota
	// bodyObject sends the BodyFields as a JSON object.
	bodyObject
	// bodyArray sends the "items" argument as a JSON array.
	bodyArray
	// bodyMultipart sends the BodyFields and FileFields as multipart/form-data.
	bodyMultipart
)

const bodyArrayArgument = "items"

type paramDef struct {
	Name     string
	Kind     paramKind
	Required bool
}

type headerDef struct {
	// Arg is the name of the tool argument.
	Arg string
	// Header is the name of the HTTP header the argument is sent in.
	Header   string
	Required bool
}

// requestSpec describes how the arguments of a tool map onto an API request.
// Path parameters are handled by the tool itself as they are part of the client method signature.
type requestSpec struct {
	Query      []paramDef
	Headers    []headerDef
	Body       bodyKind
	BodyFields []paramDef
	// FileFields are arguments that hold the path of a local file to upload.
	FileFields []paramDef
}

// pathValue returns a path parameter as a string. Numbers are accepted for integer path parameters.
func pathValue(property string, vals map[string]any) (string, error) {
	v, ok := vals[property]
	if !ok {
		return "", fmt.Errorf("missing required parameter: %s", property)
	}

	switch t := v.(type) {
	case string:
		if t == "" {
			return "", fmt.Errorf("missing required parameter: %s", property)
		}
		return t, nil
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("%s is not of type string", property)
	}
}

// buildRequestParams turns tool arguments into request parameters as described by spec.
func buildRequestParams(args map[string]any, spec requestSpec) (v1client.RequestParams, error) {
	var p v1client.RequestParams

	if len(spec.Query) > 0 {
		p.Query = url.Values{}
		for _, def := range spec.Query {
			s, ok, err := stringArgument(args, def)
			if err != nil {
				return p, err
			}
			if ok {
				p.Query.Set(def.Name, s)
			}
		}
	}

	if len(spec.Headers) > 0 {
		p.Headers = map[string]string{}
		for _, h := range spec.Headers {
			s, ok, err := stringArgument(args, paramDef{Name: h.Arg, Kind: kindString, Required: h.Required})
			if err != nil {
				return p, err
			}
			if ok {
				p.Headers[h.Header] = s
			}
		}
	}

	switch spec.Body {
	case bodyObject:
		body := map[string]any{}
		for _, def := range spec.BodyFields {
			v, ok, err := typedArgument(args, def)
			if err != nil {
				return p, err
			}
			if ok {
				body[def.Name] = v
			}
		}
		if len(body) > 0 {
			p.Body = body
		}
	case bodyArray:
		v, ok, err := typedArgument(args, paramDef{Name: bodyArrayArgument, Kind: kindArray, Required: true})
		if err != nil {
			return p, err
		}
		if ok {
			p.Body = v
		}
	case bodyMultipart:
		form := &v1client.MultipartForm{Fields: map[string]string{}, Files: map[string]string{}}
		for _, def := range spec.BodyFields {
			v, ok, err := typedArgument(args, def)
			if err != nil {
				return p, err
			}
			if !ok {
				continue
			}
			s, err := formValue(v)
			if err != nil {
				return p, fmt.Errorf("%s: %w", def.Name, err)
			}
			form.Fields[def.Name] = s
		}
		for _, def := range spec.FileFields {
			s, ok, err := stringArgument(args, def)
			if err != nil {
				return p, err
			}
			if ok {
				form.Files[def.Name] = s
			}
		}
		p.Form = form
	}

	return p, nil
}

// typedArgument returns the argument if present, checking it has the type described by def.
func typedArgument(args map[string]any, def paramDef) (any, bool, error) {
	v, ok := args[def.Name]
	if !ok || v == nil {
		if def.Required {
			return nil, false, fmt.Errorf("missing required parameter: %s", def.Name)
		}
		return nil, false, nil
	}

	valid := false
	switch def.Kind {
	case kindString:
		s, isString := v.(string)
		valid = isString
		if isString && s == "" {
			if def.Required {
				return nil, false, fmt.Errorf("missing required parameter: %s", def.Name)
			}
			return nil, false, nil
		}
	case kindNumber:
		_, valid = v.(float64)
	case kindBoolean:
		_, valid = v.(bool)
	case kindArray:
		_, valid = v.([]any)
	case kindObject:
		_, valid = v.(map[string]any)
	}
	if !valid {
		return nil, false, fmt.Errorf("%s is not of type %s", def.Name, def.Kind)
	}

	return v, true, nil
}

// stringArgument returns a scalar argument formatted as a string.
func stringArgument(args map[string]any, def paramDef) (string, bool, error) {
	v, ok, err := typedArgument(args, def)
	if err != nil || !ok {
		return "", false, err
	}
	s, err := formValue(v)
	if err != nil {
		return "", false, fmt.Errorf("%s: %w", def.Name, err)
	}
	return s, true, nil
}

func formValue(v any) (string, error) {
	switch t := v.(type) {
	case string:
		return t, nil
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(t), nil
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
}

// responseHeaders are returned to the caller as they carry information needed by follow-up calls.
// ETag is required for If-Match and Operation-Location identifies the task of asynchronous operations.
var responseHeaders = []string{"ETag", "Operation-Location", "Location"}

// handleResponse converts the response of any 2xx status code into a tool result.
// Binary responses are returned as embedded resources.
func handleResponse(r *http.Response, err error, msg string) (*mcp.CallToolResult, error) {
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = r.Body.Close()
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if r.StatusCode < 200 || r.StatusCode > 299 {
		return mcp.NewToolResultError(fmt.Sprintf("%s (HTTP %d): %s", msg, r.StatusCode, string(body))), nil
	}

	content := make([]mcp.Content, 0, 2)
	switch {
	case len(body) == 0:
		content = append(content, mcp.NewTextContent(fmt.Sprintf("Request succeeded (HTTP %d).", r.StatusCode)))
	case isTextContent(r.Header.Get("Content-Type")):
		content = append(content, mcp.NewTextContent(string(body)))
	default:
		mimeType := r.Header.Get("Content-Type")
		content = append(content, mcp.NewEmbeddedResource(mcp.BlobResourceContents{
			URI:      "vision-one:response",
			MIMEType: mimeType,
			Blob:     base64.StdEncoding.EncodeToString(body),
		}))
	}

	var headers []string
	for _, name := range responseHeaders {
		if v := r.Header.Get(name); v != "" {
			headers = append(headers, fmt.Sprintf("%s: %s", name, v))
		}
	}
	if len(headers) > 0 {
		content = append(content, mcp.NewTextContent(strings.Join(headers, "\n")))
	}

	return &mcp.CallToolResult{Content: content}, nil
}

func isTextContent(contentType string) bool {
	if contentType == "" {
		return true
	}

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}

	switch {
	case strings.HasPrefix(mediaType, "text/"),
		strings.HasSuffix(mediaType, "json"),
		strings.HasSuffix(mediaType, "xml"),
		strings.HasSuffix(mediaType, "yaml"),
		strings.HasSuffix(mediaType, "csv"):
		return true
	}
	return false
}
