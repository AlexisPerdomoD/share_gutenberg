package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	c "share-Gutenberg/config"
	s "share-Gutenberg/services"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	//load env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("env variables problem")
	}
	// Create an instance of the app structure
	app := NewApp()
	db, errDB := c.ConnectUsersDB()
	if errDB != nil {
		fmt.Println("there was an error in the db")
		os.Exit(1)
	}

	cm := s.CMT{DB: db}
	um := s.UMT{DB: db}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "share-Gutenberg",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			&app,
			&cm,
			&um,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
