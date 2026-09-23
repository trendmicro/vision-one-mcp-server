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
