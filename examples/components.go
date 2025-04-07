package main

import (
	"fmt"
	"govinci/core"
	"govinci/htmlout"
)

func App(ctx *core.Context) core.View {
	name := core.NewState(ctx, "")
	password := core.NewState(ctx, "")
	age := core.NewState(ctx, 0)
	bio := core.NewState(ctx, "")
	termsAccepted := core.NewState(ctx, false)

	return core.Column(
		core.Text("Criar Conta"),

		core.Input(name.Get(), "Nome completo", func(val string) {
			name.Set(val)
		}),

		core.InputPassword(password.Get(), "Senha", func(val string) {
			password.Set(val)
		}),

		core.NumericInput(age.Get(), func(val int) {
			age.Set(val)
		}),

		core.TextArea(bio.Get(), func(val string) {
			bio.Set(val)
		}, 4),

		core.Checkbox(termsAccepted.Get(), func(val bool) {
			termsAccepted.Set(val)
		}),

		core.Spacer(16),

		core.Text(fmt.Sprintf("Olá, %s (%d anos)", name.Get(), age.Get())),
		core.Text(fmt.Sprintf("Biografia: %s", bio.Get())),
		core.Text(fmt.Sprintf("Termos aceites: %v", termsAccepted.Get())),
	)
}

func main() {
	ctx := core.NewContext()
	node := App(ctx).Render(ctx)
	output := htmlout.ExportHTML(node)

	fmt.Println(output)
}
