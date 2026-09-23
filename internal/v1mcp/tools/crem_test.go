package tools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCREMFirstSeenDateRangeSchema guards against the bug where
// firstSeenStartDateTime was registered twice (the second registration
// overwriting the first with the wrong description) and
// firstSeenEndDateTime was never exposed in the tool's input schema,
// even though the handler and client both accept it.
func TestCREMFirstSeenDateRangeSchema(t *testing.T) {
	for _, tc := range []struct {
		name    string
		factory func() map[string]any
	}{
		{
			name: "crem_attack_surface_devices_list",
			factory: func() map[string]any {
				return toolCREMAttackSurfaceDevicesList(nil).Tool.InputSchema.Properties
			},
		},
		{
			name: "crem_attack_surface_cloud_assets_list",
			factory: func() map[string]any {
				return toolCREMAttackSurfaceCloudAssetsList(nil).Tool.InputSchema.Properties
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			props := tc.factory()

			require.Contains(t, props, "firstSeenStartDateTime")
			require.Contains(t, props, "firstSeenEndDateTime")

			start, ok := props["firstSeenStartDateTime"].(map[string]any)
			require.True(t, ok)
			require.Contains(t, start["description"], "start time")

			end, ok := props["firstSeenEndDateTime"].(map[string]any)
			require.True(t, ok)
			require.Contains(t, end["description"], "end time")
		})
	}
}
