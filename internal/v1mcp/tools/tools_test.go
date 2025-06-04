package tools

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWithOrdering(t *testing.T) {
	expected := []string{
		"key asc",
		"key desc",
		"key2 asc",
		"key2 desc",
	}
	actual := withOrdering(asc_desc, "key", "key2")
	require.Equal(t, expected, actual)
}

func TestRequiredValue(t *testing.T) {
	vals := map[string]any{
		"top": 1,
	}

	t.Run("should error on missing param", func(t *testing.T) {
		_, err := requiredValue[string]("m", vals)
		require.EqualError(t, err, "missing required parameter: m")
	})
}

func TestOptionalStrInt(t *testing.T) {
	vals := map[string]any{
		"top":        "1",
		"notStrInt":  "notStrInt",
		"anotherTop": 1,
	}

	t.Run("should return integer if found", func(t *testing.T) {
		n, err := optionalStrInt("top", vals)
		require.NoError(t, err)
		require.Equal(t, 1, n)
	})

	t.Run("should return 0 if not found", func(t *testing.T) {
		n, err := optionalStrInt("notHere", vals)
		require.NoError(t, err)
		require.Equal(t, 0, n)
	})

	t.Run("should return an error if cannot be parsed to int", func(t *testing.T) {
		n, err := optionalStrInt("notStrInt", vals)
		require.Error(t, err)
		require.Equal(t, 0, n)
	})

	t.Run("should return an error if not string", func(t *testing.T) {
		_, err := optionalStrInt("anotherTop", vals)
		require.Error(t, err)
	})
}

func TestOptionalValue(t *testing.T) {
	vals := map[string]any{
		"top": 1,
	}

	t.Run("should return default value if not found", func(t *testing.T) {
		n, err := optionalValue[int]("none", vals)
		require.NoError(t, err)
		require.Zero(t, n)
	})

	t.Run("should return an error if value is wrong type", func(t *testing.T) {
		_, err := optionalValue[string]("top", vals)
		require.Error(t, err)
	})

	t.Run("should return value if found", func(t *testing.T) {
		n, err := optionalValue[int]("top", vals)
		require.NoError(t, err)
		require.Equal(t, 1, n)
	})
}

func TestOptionalPointerValue(t *testing.T) {
	vals := map[string]any{
		"top": 1,
	}

	t.Run("should return a nil pointer if not found", func(t *testing.T) {
		n, err := optionalPointerValue[int]("none", vals)
		require.NoError(t, err)
		require.Nil(t, n)
	})

	t.Run("should return value if found", func(t *testing.T) {
		n, err := optionalPointerValue[int]("top", vals)
		require.NoError(t, err)
		require.Equal(t, 1, *n)
	})
}
