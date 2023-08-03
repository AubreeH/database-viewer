package main

import (
	"context"
	"embed"

	"github.com/AubreeH/database-viewer/src/bindings/connectionsBinding"
	"github.com/AubreeH/database-viewer/src/bindings/queryBinding"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	_ "modernc.org/sqlite"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	conn, connContextCallback := connectionsBinding.NewConnectionsBinding()
	query, queryContextCallback := queryBinding.NewQueryBinding()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Aubree's Database Viewer",
		Width:            1024,
		Height:           768,
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 48, G: 48, B: 56, A: 0},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			connContextCallback(ctx)
			queryContextCallback(ctx)
		},
		Bind: []interface{}{
			conn,
			query,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
