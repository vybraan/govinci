package main

import (
	"fmt"
	"govinci/core"
	"govinci/htmlout"
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
	runtime := core.NewRuntime(App(ctx), ctx)

	// First render
	tree := runtime.Render()
	html := htmlout.ExportHTML(tree)
	fmt.Println("Primeiro Render:")
	fmt.Println(html)

	// Simula evento de input vindo do nativo
	runtime.SendEvent(map[string]any{
		"callback": "txt_cb_0",
		"value":    "Ismael",
	})

	// Re-render após evento
	tree = runtime.Render()
	html = htmlout.ExportHTML(tree)
	fmt.Println("\nApós evento de input:")
	fmt.Println(html)
}
