package main

import (
	"embed"
	"tugas-gui/interfaces/api"
	"tugas-gui/internal/config"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	cfg, err := config.NewConfig("config/config.yaml")
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	apiInstance := api.NewApi(cfg)

	err = wails.Run(&options.App{
		Title:  "Sensor Monitor",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: apiInstance.Startup,
		Bind: []interface{}{
			apiInstance,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
