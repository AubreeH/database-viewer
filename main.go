package main

import (
	"context"
	"embed"

	"github.com/AubreeH/database-viewer/src/bindings/connectionsBinding"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	conn, connContextCallback := connectionsBinding.NewConnectionsBinding()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "myproject",
		Width:            1024,
		Height:           768,
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			connContextCallback(ctx)
		},
		Bind: []interface{}{
			app,
			conn,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
