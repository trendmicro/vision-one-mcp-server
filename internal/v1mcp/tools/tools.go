package tools

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

var asc_desc = []string{"asc", "desc"}

// Accepts an array of keys used to sort and returns all the available combinations.
// withOrdering("hello") -> ["hello asc", "hello desc"]
func withOrdering(keywords []string, keys ...string) []string {
	orderingKeys := make([]string, 0, len(keys)*2)
	for _, key := range keys {
		for _, op := range keywords {
			orderingKeys = append(orderingKeys, fmt.Sprintf("%s %s", key, op))
		}
	}
	return orderingKeys
}

func requiredValue[T comparable](property string, vals map[string]any) (T, error) {
	var defaultValue T

	if _, ok := vals[property]; !ok {
		return defaultValue, fmt.Errorf("missing required parameter: %s", property)
	}

	if _, ok := vals[property].(T); !ok {
		return defaultValue, fmt.Errorf("%s is not of type %T", property, defaultValue)
	}

	if vals[property] == defaultValue {
		return defaultValue, fmt.Errorf("missing required parameter: %s", property)
	}

	return vals[property].(T), nil
}

func optionalIntValue(property string, vals map[string]any) (int, error) {
	val, err := optionalValue[float64](property, vals)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

// optionalStrInt retrieves an optional string and converts
// to an integer via strconv.Atoi.
func optionalStrInt(property string, vals map[string]any) (int, error) {
	val, err := optionalValue[string](property, vals)
	if err != nil {
		return 0, err
	}

	if val == "" {
		return 0, nil
	}

	n, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return n, err
}

func optionalTimeValue(property string, vals map[string]any) (time.Time, error) {
	val, err := optionalValue[string](property, vals)
	if err != nil {
		return time.Time{}, err
	}

	if val == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return time.Time{}, fmt.Errorf("%s is not an RFC3339 time string: %q", property, val)
	}
	return t, nil
}

func optionalValue[T comparable](property string, vals map[string]any) (T, error) {
	var defaultValue T

	if _, ok := vals[property]; !ok {
		return defaultValue, nil
	}

	if _, ok := vals[property].(T); !ok {
		return defaultValue, fmt.Errorf("%s is not of type %T", property, defaultValue)
	}

	return vals[property].(T), nil
}

func optionalPointerValue[T any](property string, vals map[string]any) (*T, error) {
	var defaultValue *T

	if _, ok := vals[property]; !ok {
		return defaultValue, nil
	}

	if _, ok := vals[property].(T); !ok {
		return defaultValue, fmt.Errorf("%s is not of type %T", property, defaultValue)
	}

	returnVal, _ := vals[property].(T)
	return &returnVal, nil
}

func handleStatusResponse(r *http.Response, err error, expectedStatusCode int, msg string) (*mcp.CallToolResult, error) {
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

	if r.StatusCode != expectedStatusCode {
		return mcp.NewToolResultError(fmt.Sprintf("%s: %s", msg, string(body))), nil
	}

	return mcp.NewToolResultText(string(body)), nil
}

func toPtr[T any](t T) *T {
	return &t
}
