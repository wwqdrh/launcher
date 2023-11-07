package main

import (
	"context"
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

func (a *App) AppList() []string {
	// _, res := cli.DefaultCommandWorkspace.AppList(false)
	return []string{}
}

func (a *App) StartApp(name string) string {
	// go func() {
	// 	err := cli.DefaultCommandWorkspace.StartApp(name)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }()
	return "http://localhost:8000"
}
