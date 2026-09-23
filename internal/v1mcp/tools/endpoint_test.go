package tools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestScheduleVEventSchemaConstraints guards against the bug where the endpoint
// security schedules create/update tools accepted vEvents as an untyped object,
// giving the model no guidance about which iCalendar fields the API actually
// accepts. dtStart, duration and rRule must each carry a pattern that mirrors
// the API's own validation (only DAILY/WEEKLY/MONTHLY frequencies for rRule,
// only days/hours for duration), so invalid values are rejected before a
// request is ever sent.
func TestScheduleVEventSchemaConstraints(t *testing.T) {
	for _, tc := range []struct {
		name    string
		factory func() map[string]any
	}{
		{
			name: "endpoint_security_schedules_create",
			factory: func() map[string]any {
				return toolEndpointSecuritySchedulesCreate(nil).Tool.InputSchema.Properties
			},
		},
		{
			name: "endpoint_security_schedules_update",
			factory: func() map[string]any {
				return toolEndpointSecuritySchedulesUpdate(nil).Tool.InputSchema.Properties
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			itemsProp, ok := tc.factory()["items"].(map[string]any)
			require.True(t, ok, "expected an \"items\" property")

			scheduleSchema, ok := itemsProp["items"].(map[string]any)
			require.True(t, ok, "expected a schema for each schedule object")

			scheduleProps, ok := scheduleSchema["properties"].(map[string]any)
			require.True(t, ok)

			vEvents, ok := scheduleProps["vEvents"].(map[string]any)
			require.True(t, ok, "expected a vEvents property")
			require.Equal(t, 1, vEvents["minItems"])
			require.Equal(t, 1, vEvents["maxItems"])

			vEventSchema, ok := vEvents["items"].(map[string]any)
			require.True(t, ok, "expected vEvents items to have a defined schema, not an untyped object")

			vEventProps, ok := vEventSchema["properties"].(map[string]any)
			require.True(t, ok)

			dtStart, ok := vEventProps["dtStart"].(map[string]any)
			require.True(t, ok)
			require.Equal(t, "^[0-9]{8}T[0-9]{6}$", dtStart["pattern"])

			duration, ok := vEventProps["duration"].(map[string]any)
			require.True(t, ok)
			require.Equal(t, `^P(\d+D)?(T(\d+H)?)?$`, duration["pattern"])

			rRule, ok := vEventProps["rRule"].(map[string]any)
			require.True(t, ok)
			require.Equal(t, `^FREQ=(DAILY|WEEKLY|MONTHLY)(;[A-Z]+=[^;]+)*$`, rRule["pattern"])

			required, ok := vEventSchema["required"].([]string)
			require.True(t, ok)
			require.ElementsMatch(t, []string{"dtStart", "duration", "rRule"}, required)
		})
	}
}

// TestValidateScheduleItems guards against the schedule create/update tools silently
// forwarding an invalid vEvent to the Trend Vision One API, which responds with an opaque
// error code (e.g. "Error_001001") and no explanation. validateScheduleItems must catch the
// same violations client-side and return a specific, human-readable error instead.
func TestValidateScheduleItems(t *testing.T) {
	validItem := map[string]any{
		"name": "test",
		"vEvents": []any{
			map[string]any{
				"dtStart":  "20261001T020000",
				"duration": "PT2H",
				"rRule":    "FREQ=DAILY",
			},
		},
	}

	t.Run("valid item passes", func(t *testing.T) {
		err := validateScheduleItems(map[string]any{"items": []any{validItem}})
		require.NoError(t, err)
	})

	t.Run("no items is fine", func(t *testing.T) {
		require.NoError(t, validateScheduleItems(map[string]any{}))
	})

	for _, tc := range []struct {
		name    string
		field   string
		value   string
		wantErr string
	}{
		{name: "bad dtStart", field: "dtStart", value: "2026-10-01T02:00:00Z", wantErr: "invalid dtStart"},
		{name: "bad duration unit", field: "duration", value: "P1W", wantErr: "invalid duration"},
		{name: "bad duration minutes", field: "duration", value: "PT30M", wantErr: "invalid duration"},
		{name: "bad rRule frequency", field: "rRule", value: "FREQ=YEARLY", wantErr: "invalid rRule"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			vEvent := map[string]any{
				"dtStart":  "20261001T020000",
				"duration": "PT2H",
				"rRule":    "FREQ=DAILY",
			}
			vEvent[tc.field] = tc.value

			item := map[string]any{
				"name":    "test",
				"vEvents": []any{vEvent},
			}

			err := validateScheduleItems(map[string]any{"items": []any{item}})
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.wantErr)
			require.Contains(t, err.Error(), tc.value)
		})
	}
}
