package tools

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/require"
)

func TestPathValue(t *testing.T) {
	vals := map[string]any{"id": "abc", "num": float64(42), "empty": "", "bad": true}

	v, err := pathValue("id", vals)
	require.NoError(t, err)
	require.Equal(t, "abc", v)

	v, err = pathValue("num", vals)
	require.NoError(t, err)
	require.Equal(t, "42", v)

	_, err = pathValue("missing", vals)
	require.EqualError(t, err, "missing required parameter: missing")

	_, err = pathValue("empty", vals)
	require.EqualError(t, err, "missing required parameter: empty")

	_, err = pathValue("bad", vals)
	require.EqualError(t, err, "bad is not of type string")
}

func TestBuildRequestParams(t *testing.T) {
	t.Run("query and headers", func(t *testing.T) {
		spec := requestSpec{
			Query: []paramDef{
				{Name: "top", Kind: kindString},
				{Name: "count", Kind: kindNumber},
				{Name: "flag", Kind: kindBoolean},
				{Name: "unset", Kind: kindString},
			},
			Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}, {Arg: "ifMatch", Header: "If-Match"}},
		}

		p, err := buildRequestParams(map[string]any{
			"top": "50", "count": float64(3), "flag": true, "filter": "a eq 'b'",
		}, spec)
		require.NoError(t, err)
		require.Equal(t, "50", p.Query.Get("top"))
		require.Equal(t, "3", p.Query.Get("count"))
		require.Equal(t, "true", p.Query.Get("flag"))
		require.False(t, p.Query.Has("unset"))
		require.Equal(t, map[string]string{"TMV1-Filter": "a eq 'b'"}, p.Headers)
		require.Nil(t, p.Body)
	})

	t.Run("required values", func(t *testing.T) {
		_, err := buildRequestParams(map[string]any{}, requestSpec{Query: []paramDef{{Name: "q", Kind: kindString, Required: true}}})
		require.EqualError(t, err, "missing required parameter: q")

		_, err = buildRequestParams(map[string]any{}, requestSpec{Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}}})
		require.EqualError(t, err, "missing required parameter: query")
	})

	t.Run("wrong types", func(t *testing.T) {
		_, err := buildRequestParams(map[string]any{"count": "3"}, requestSpec{Query: []paramDef{{Name: "count", Kind: kindNumber}}})
		require.EqualError(t, err, "count is not of type number")
	})

	t.Run("object body", func(t *testing.T) {
		spec := requestSpec{
			Body: bodyObject,
			BodyFields: []paramDef{
				{Name: "name", Kind: kindString, Required: true},
				{Name: "enabled", Kind: kindBoolean},
				{Name: "tags", Kind: kindArray},
				{Name: "settings", Kind: kindObject},
				{Name: "skipped", Kind: kindString},
			},
		}

		p, err := buildRequestParams(map[string]any{
			"name": "n", "enabled": false, "tags": []any{"a"}, "settings": map[string]any{"k": "v"},
		}, spec)
		require.NoError(t, err)
		require.Equal(t, map[string]any{
			"name": "n", "enabled": false, "tags": []any{"a"}, "settings": map[string]any{"k": "v"},
		}, p.Body)

		_, err = buildRequestParams(map[string]any{}, spec)
		require.EqualError(t, err, "missing required parameter: name")

		_, err = buildRequestParams(map[string]any{"name": "n", "tags": "a"}, spec)
		require.EqualError(t, err, "tags is not of type array")
	})

	t.Run("array body", func(t *testing.T) {
		p, err := buildRequestParams(map[string]any{"items": []any{map[string]any{"a": "b"}}}, requestSpec{Body: bodyArray})
		require.NoError(t, err)
		require.Equal(t, []any{map[string]any{"a": "b"}}, p.Body)

		_, err = buildRequestParams(map[string]any{}, requestSpec{Body: bodyArray})
		require.EqualError(t, err, "missing required parameter: items")
	})

	t.Run("multipart body", func(t *testing.T) {
		spec := requestSpec{
			Body:       bodyMultipart,
			BodyFields: []paramDef{{Name: "name", Kind: kindString}, {Name: "enabled", Kind: kindBoolean}, {Name: "options", Kind: kindObject}},
			FileFields: []paramDef{{Name: "file", Kind: kindString, Required: true}},
		}

		p, err := buildRequestParams(map[string]any{
			"name": "n", "enabled": true, "options": map[string]any{"a": float64(1)}, "file": "/tmp/f.txt",
		}, spec)
		require.NoError(t, err)
		require.Equal(t, map[string]string{"name": "n", "enabled": "true", "options": `{"a":1}`}, p.Form.Fields)
		require.Equal(t, map[string]string{"file": "/tmp/f.txt"}, p.Form.Files)

		_, err = buildRequestParams(map[string]any{"name": "n"}, spec)
		require.EqualError(t, err, "missing required parameter: file")
	})
}

func doResponse(t *testing.T, status int, header http.Header, body string) *http.Response {
	t.Helper()
	rec := httptest.NewRecorder()
	for k, v := range header {
		rec.Header()[k] = v
	}
	rec.WriteHeader(status)
	_, err := io.WriteString(rec, body)
	require.NoError(t, err)
	return rec.Result()
}

func textOf(t *testing.T, c mcp.Content) string {
	t.Helper()
	tc, ok := c.(mcp.TextContent)
	require.True(t, ok, "expected text content, got %T", c)
	return tc.Text
}

func TestHandleResponse(t *testing.T) {
	t.Run("returns json bodies", func(t *testing.T) {
		r := doResponse(t, http.StatusOK, http.Header{"Content-Type": {"application/json"}, "Etag": {`"v1"`}}, `{"a":1}`)
		res, err := handleResponse(r, nil, "failed")
		require.NoError(t, err)
		require.False(t, res.IsError)
		require.Len(t, res.Content, 2)
		require.Equal(t, `{"a":1}`, textOf(t, res.Content[0]))
		require.Equal(t, `ETag: "v1"`, textOf(t, res.Content[1]))
	})

	t.Run("accepts every 2xx status", func(t *testing.T) {
		for _, status := range []int{http.StatusCreated, http.StatusAccepted, http.StatusMultiStatus} {
			r := doResponse(t, status, http.Header{"Content-Type": {"application/json"}}, `[]`)
			res, err := handleResponse(r, nil, "failed")
			require.NoError(t, err)
			require.False(t, res.IsError, status)
		}
	})

	t.Run("reports empty responses and location headers", func(t *testing.T) {
		r := doResponse(t, http.StatusAccepted, http.Header{"Operation-Location": {"https://example.com/tasks/1"}}, "")
		res, err := handleResponse(r, nil, "failed")
		require.NoError(t, err)
		require.Equal(t, "Request succeeded (HTTP 202).", textOf(t, res.Content[0]))
		require.Equal(t, "Operation-Location: https://example.com/tasks/1", textOf(t, res.Content[1]))
	})

	t.Run("returns binary bodies as resources", func(t *testing.T) {
		r := doResponse(t, http.StatusOK, http.Header{"Content-Type": {"application/zip"}}, "PK\x03\x04")
		res, err := handleResponse(r, nil, "failed")
		require.NoError(t, err)
		require.Len(t, res.Content, 1)
		embedded, ok := res.Content[0].(mcp.EmbeddedResource)
		require.True(t, ok)
		blob, ok := embedded.Resource.(mcp.BlobResourceContents)
		require.True(t, ok)
		require.Equal(t, "application/zip", blob.MIMEType)
		require.Equal(t, "UEsDBA==", blob.Blob)
	})

	t.Run("returns csv as text", func(t *testing.T) {
		r := doResponse(t, http.StatusOK, http.Header{"Content-Type": {"text/csv; charset=utf-8"}}, "a,b")
		res, err := handleResponse(r, nil, "failed")
		require.NoError(t, err)
		require.Equal(t, "a,b", textOf(t, res.Content[0]))
	})

	t.Run("returns errors for non 2xx responses", func(t *testing.T) {
		r := doResponse(t, http.StatusForbidden, nil, `{"error":"denied"}`)
		res, err := handleResponse(r, nil, "failed to do thing")
		require.NoError(t, err)
		require.True(t, res.IsError)
		require.Equal(t, `failed to do thing (HTTP 403): {"error":"denied"}`, textOf(t, res.Content[0]))
	})

	t.Run("propagates transport errors", func(t *testing.T) {
		_, err := handleResponse(nil, io.ErrUnexpectedEOF, "failed")
		require.ErrorIs(t, err, io.ErrUnexpectedEOF)
	})
}
