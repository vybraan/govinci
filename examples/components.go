package main

import (
	"fmt"
	"govinci/core"
	"govinci/htmlout"
)

func App(ctx *core.Context) core.View {
	return core.WithTheme(&core.DefaultTheme,
		core.Column(
			core.Text("👋 Govinci DSL",
				core.UseStyle(ctx.Theme().Typography.Title),
				core.Align(core.AlignCenter),
			),
			core.Text("Elegant declarative UI in Go.",
				core.UseStyle(ctx.Theme().Typography.Body),
				core.Padding(12),
				core.TextColor(ctx.Theme().Colors.TextSecondary),
			),
			core.Text("Styled with love.",
				core.FontSize(14),
				core.FontWeight(core.Light),
				core.TextColor("#999"),
				core.Margin(8),
			),
		),
	)
}

func main() {
	ctx := core.NewContext()
	node := App(ctx).Render(ctx)
	output := htmlout.ExportHTML(node)

	fmt.Println(output)
}
