package main

import (
	"fmt"
	"govinci/core"
	"govinci/render"
)

func App(ctx *core.Context) core.View {
	name := core.NewState(ctx, "")

	return core.Column(
		core.Text("Bem-vindo ao Govinci"),
		core.Input(name.Get(), "Digite o seu nome", func(val string) {
			name.Set(val)
		}),
		core.Text("Olá, "+name.Get()),
	)
}
func main() {
	ctx := core.NewContext()
	manager := render.New(ctx, App)

	fmt.Println("🔁 Primeira renderização:")
	fullRender := manager.RenderAndGetPatches()
	fmt.Println(fullRender)

	// Simulando evento de input
	fmt.Println("✏️ Simulando input...")
	core.TriggerTextCallback("txt_cb_0", "Ismael")

	// Re-render com patches
	fmt.Println("🔁 Re-render com patches:")
	patches := manager.RenderAndGetPatches()
	fmt.Println(patches)
}
