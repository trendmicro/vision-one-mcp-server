package v1client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// MultipartForm describes a multipart/form-data request body.
type MultipartForm struct {
	// Fields are plain form fields.
	Fields map[string]string
	// Files maps a form field name to the path of a local file to upload.
	Files map[string]string
}

// RequestParams holds everything a request needs other than its method and path.
type RequestParams struct {
	Query   url.Values
	Headers map[string]string
	// Body is marshalled as JSON when not nil. Ignored when Form is set.
	Body any
	Form *MultipartForm
}

// pathf formats a path, URL-escaping every argument so that identifiers can't alter the path.
func pathf(format string, args ...string) string {
	escaped := make([]any, len(args))
	for i, a := range args {
		escaped[i] = url.PathEscape(a)
	}
	return fmt.Sprintf(format, escaped...)
}

// do sends a request built from method, path and params.
// Path MUST NOT start with a "/".
func (c *V1ApiClient) do(method, path string, p RequestParams) (*http.Response, error) {
	var (
		body        io.Reader = http.NoBody
		contentType string
	)

	switch {
	case p.Form != nil:
		b, ct, err := encodeMultipart(p.Form)
		if err != nil {
			return nil, err
		}
		body, contentType = b, ct
	case p.Body != nil:
		b, err := json.Marshal(p.Body)
		if err != nil {
			return nil, err
		}
		body, contentType = bytes.NewReader(b), "application/json"
	}

	opts := make([]requestOptionFunc, 0, len(p.Headers)+3)
	if len(p.Query) > 0 {
		opts = append(opts, withUrlParameters(p.Query))
	}
	if contentType != "" {
		opts = append(opts, withHeader("Content-Type", contentType))
	}

	names := make([]string, 0, len(p.Headers))
	for name := range p.Headers {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		opts = append(opts, withHeader(name, p.Headers[name]))
	}

	r, err := c.newRequest(method, path, body, opts...)
	if err != nil {
		return nil, err
	}

	return c.client.Do(r)
}

func encodeMultipart(form *MultipartForm) (*bytes.Buffer, string, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	fields := make([]string, 0, len(form.Fields))
	for name := range form.Fields {
		fields = append(fields, name)
	}
	sort.Strings(fields)
	for _, name := range fields {
		if err := w.WriteField(name, form.Fields[name]); err != nil {
			return nil, "", err
		}
	}

	files := make([]string, 0, len(form.Files))
	for name := range form.Files {
		files = append(files, name)
	}
	sort.Strings(files)
	for _, name := range files {
		if err := addFilePart(w, name, form.Files[name]); err != nil {
			return nil, "", err
		}
	}

	if err := w.Close(); err != nil {
		return nil, "", err
	}
	return buf, w.FormDataContentType(), nil
}

func addFilePart(w *multipart.Writer, field, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open %s file: %w", field, err)
	}
	defer func() {
		_ = f.Close()
	}()

	part, err := w.CreateFormFile(field, filepath.Base(strings.TrimSpace(path)))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, f)
	return err
}
