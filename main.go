package main

import (
	"embed"

	"vgtracker/backend"
	"vgtracker/backend/api"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := backend.NewApp()

	// Instantiate binding structs
	// https://wails.io/docs/howdoesitwork/#method-binding

	// Create application with options
	//https://wails.io/docs/reference/options
	err := wails.Run(&options.App{
		Title:  "vgtracker",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
			&api.ProfileBackend{},
		},
		Debug: options.Debug{
			// enables debugging ui
			OpenInspectorOnStartup: false,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
