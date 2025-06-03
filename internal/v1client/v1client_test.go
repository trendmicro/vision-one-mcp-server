package v1client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewV1ApiClient(t *testing.T) {
	t.Run("invalid configuration", func(t *testing.T) {
		require.Panics(t, func() { _, _ = NewV1ApiClient(ClientOptions{}) }, "the function did not panic")
	})

	t.Run("valid configuration", func(t *testing.T) {
		c, err := NewV1ApiClient(ClientOptions{
			Region: "au",
		})
		require.NotNil(t, c, "expected client instead found nil pointer")
		require.Nil(t, err, "expected error to be nil")
		require.Equal(t, "api.au.xdr.trendmicro.com", c.baseUrl.Hostname())
	})

	t.Run("expect host to be used instead of region", func(t *testing.T) {
		d, err := NewV1ApiClient(ClientOptions{
			Region: "au",
			Host:   "some.trendmicro.com",
		})
		require.Nil(t, err)
		require.Equal(
			t,
			"some.trendmicro.com",
			d.baseUrl.Hostname(),
			"wanted some.trendmicro.com, got %s",
			d.baseUrl.Hostname(),
		)
	})
}
