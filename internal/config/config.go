package config

import (
	_ "embed"
	"encoding/json"
)

var (
	WinWidth  float64 = 800
	WinHeight float64 = 600
	TargetFPS float64 = 30
)

var TargetFrameTime float64

//go:embed config.json
var configJSON []byte

func LoadConfig() {
	rawConfig := make(map[string]any)
	json.Unmarshal(configJSON, &rawConfig)

	windowConfig := rawConfig["window"].(map[string]any)
	WinHeight = windowConfig["height"].(float64)
	WinWidth = windowConfig["width"].(float64)
	TargetFPS = rawConfig["targetFPS"].(float64)

	TargetFrameTime = 1000.0 / TargetFPS
}
