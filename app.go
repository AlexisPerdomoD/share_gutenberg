package main

import (
	"context"
	"net/url"
	m "share-Gutenberg/models"
	s "share-Gutenberg/services"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// HERE ITS A GOOD PLACE TO SET CONTROLLERS

func (a *App) GetBooks(params map[string]string) (*m.Gutendex, error) {
	paramsDone := url.Values{}
	for key, value := range params {
		if value != "" {
			paramsDone.Set(key, value)
		}
	}

	return s.BooksFetcher(paramsDone)
}
