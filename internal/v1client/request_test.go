package v1client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) *V1ApiClient {
	t.Helper()

	srv := httptest.NewTLSServer(handler)
	t.Cleanup(srv.Close)

	u, err := url.Parse(srv.URL + "/")
	require.NoError(t, err)

	return &V1ApiClient{client: srv.Client(), apiKey: "test-key", baseUrl: u, UserAgent: "test"}
}

func TestPathf(t *testing.T) {
	require.Equal(t, "v3.0/a/one/b/two", pathf("v3.0/a/%s/b/%s", "one", "two"))
	require.Equal(t, "v3.0/a/x%2Fy%3Fz", pathf("v3.0/a/%s", "x/y?z"))
}

func TestDo(t *testing.T) {
	t.Run("sends query, headers and json body", func(t *testing.T) {
		c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodPatch, r.Method)
			require.Equal(t, "/v3.0/things/abc", r.URL.Path)
			require.Equal(t, "5", r.URL.Query().Get("top"))
			require.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
			require.Equal(t, "application/json", r.Header.Get("Content-Type"))
			require.Equal(t, "etag-1", r.Header.Get("If-Match"))
			require.Empty(t, r.Header.Get("TMV1-Filter"), "empty headers must not be sent")

			b, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			var got map[string]any
			require.NoError(t, json.Unmarshal(b, &got))
			require.Equal(t, map[string]any{"name": "n"}, got)
			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := c.do(http.MethodPatch, pathf("v3.0/things/%s", "abc"), RequestParams{
			Query:   url.Values{"top": {"5"}},
			Headers: map[string]string{"If-Match": "etag-1", "TMV1-Filter": ""},
			Body:    map[string]any{"name": "n"},
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	t.Run("sends json arrays", func(t *testing.T) {
		c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			b, _ := io.ReadAll(r.Body)
			require.JSONEq(t, `[{"a":1},{"a":2}]`, string(b))
			w.WriteHeader(http.StatusMultiStatus)
		})

		resp, err := c.do(http.MethodPost, "v3.0/things/delete", RequestParams{Body: []any{map[string]any{"a": 1}, map[string]any{"a": 2}}})
		require.NoError(t, err)
		require.Equal(t, http.StatusMultiStatus, resp.StatusCode)
	})

	t.Run("sends multipart forms", func(t *testing.T) {
		dir := t.TempDir()
		file := filepath.Join(dir, "sample.txt")
		require.NoError(t, os.WriteFile(file, []byte("hello"), 0o600))

		c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			require.NoError(t, r.ParseMultipartForm(1<<20))
			require.Equal(t, "value", r.FormValue("field"))

			f, header, err := r.FormFile("file")
			require.NoError(t, err)
			defer func() { _ = f.Close() }()
			require.Equal(t, "sample.txt", header.Filename)
			b, _ := io.ReadAll(f)
			require.Equal(t, "hello", string(b))
			w.WriteHeader(http.StatusAccepted)
		})

		resp, err := c.do(http.MethodPost, "v3.0/files/analyze", RequestParams{
			Form: &MultipartForm{Fields: map[string]string{"field": "value"}, Files: map[string]string{"file": file}},
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusAccepted, resp.StatusCode)
	})

	t.Run("fails when the file can't be opened", func(t *testing.T) {
		c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			t.Error("request must not be sent")
		})

		_, err := c.do(http.MethodPost, "v3.0/files/analyze", RequestParams{
			Form: &MultipartForm{Files: map[string]string{"file": "/does/not/exist"}},
		})
		require.ErrorContains(t, err, "could not open file file")
	})
}

func TestRegions(t *testing.T) {
	for region, host := range map[string]string{
		"us":  "api.xdr.trendmicro.com",
		"jp":  "api.xdr.trendmicro.co.jp",
		"uk":  "api.uk.xdr.trendmicro.com",
		"ca":  "api.ca.xdr.trendmicro.com",
		"za":  "api.za.xdr.trendmicro.com",
		"id":  "api.id.xdr.trendmicro.com",
		"mea": "api.mea.xdr.trendmicro.com",
	} {
		u, err := getRegionURL(region)
		require.NoError(t, err)
		require.Equal(t, host, u.Hostname(), region)
	}
}
