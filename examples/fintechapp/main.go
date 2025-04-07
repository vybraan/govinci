package main

import (
	"fmt"
	"govinci/core"
	"govinci/htmlout"
)

func App(ctx *core.Context) core.View {
	return core.SafeArea(
		core.Scroll(
			core.Column(
				core.Spacer(24),

				core.Image("https://dummyimage.com/80x80/009688/ffffff&text=G", // logotipo
					core.Margin(0),
					core.Padding(8),
					core.Align(core.AlignCenter),
				),

				core.Text("GovinciFin",
					core.FontSize(24),
					core.FontWeight(core.Bold),
					core.TextColor("#009688"),
					core.Align(core.AlignCenter),
					core.Margin(12),
				),

				core.Card(
					core.Text("Saldo disponível",
						core.TextColor("#555"),
						core.FontSize(14),
					),
					core.Text("MZN 42,750.00",
						core.FontSize(22),
						core.FontWeight(core.Bold),
						core.TextColor("#009688"),
					),
					core.Spacer(12),
					core.Row(
						core.Button("Transferir", func() {}, core.BackgroundColor("#2196F3"), core.TextColor("#FFF")),
						core.Spacer(8),
						core.Button("Carregar", func() {}, core.BackgroundColor("#FFC107"), core.TextColor("#000")),
					),
				),

				core.Spacer(24),

				core.Text("Movimentos Recentes",
					core.FontSize(18),
					core.FontWeight(core.Bold),
					core.TextColor("#222"),
				),

				core.Column(
					core.Card(
						core.Text("Farmácia"),
						core.Text("-750 MZN", core.TextColor("#F44336")),
					),
					core.Card(
						core.Text("Transferência recebida"),
						core.Text("+10,000 MZN", core.TextColor("#4CAF50")),
					),
					core.Card(
						core.Text("Recarga de saldo"),
						core.Text("+3,500 MZN", core.TextColor("#4CAF50")),
					),
				),

				core.Spacer(32),
			),
		),
	)
}

func main() {
	baseCtx := core.NewContext()

	cfg := &core.AppConfig{
		Name:        "GovinciFin",
		Description: "App financeira feita com Govinci DSL",
		Version:     "1.0.0",
		Locale:      "pt-MZ",
		Author:      "Ismael GraHms",
	}

	ctx := baseCtx.WithConfig(cfg)
	node := App(ctx).Render(ctx)
	html := htmlout.ExportHTML(node)
	fmt.Println(html)
}
