//go:build govinci

package main

import (
	"govinci/core"
	"govinci/render"
	"myapp/app"
)

var manager *render.Manager

// Exported to native (called once)
func InitApp() {
	ctx := core.NewContext().With(
		core.WithThemeOpt(app.AppTheme),
		core.WithConfigOpt(app.Config),
	)
	manager = render.New(ctx, app.App)
}

// Exported to native (to get first render)
func RenderInitial() string {
	return manager.RenderAndGetPatches()
}

// Exported to native (to simulate external event)
func TriggerCallback(id string) string {
	core.TriggerCallback(id)
	return manager.RenderAndGetPatches()
}

func TriggerTextCallback(id, val string) string {
	core.TriggerTextCallback(id, val)
	return manager.RenderAndGetPatches()
}
