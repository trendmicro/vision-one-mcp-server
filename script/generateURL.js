const config = {
	name: "trend-vision-one-mcp",
	inputs: [
		{
			"type": "promptString",
			"id": "trend-vision-one-api-key",
			"description": "Trend Vision One API Key",
			"password": true
		},
		{
			"type": "promptString",
			"id": "trend-vision-one-region",
			"description": "Trend Vision One Region"
		}
	],
	command: "docker",
	args: [
		"run",
		"-i",
		"--rm",
		"-e",
		"TREND_VISION_ONE_API_KEY",
		"ghcr.io/trendmicro/vision-one-mcp-server",
		"-region",
		"${input:trend-vision-one-region}",
		"-readonly=true"
	],
	env: {
		"TREND_VISION_ONE_API_KEY": "${input:trend-vision-one-api-key}"
	}
}

const link = `vscode:mcp/install?${encodeURIComponent(JSON.stringify(config))}`;

console.log(link)
