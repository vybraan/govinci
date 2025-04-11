package social

import (
	"govinci/core"
)

func HomePage(ctx *core.Context) core.View {
	return core.Column(
		core.Text("🏠 Página Inicial", core.FontSize(24), core.FontWeight(core.Bold)),
		core.Spacer(12),
		core.Button("Abrir Detalhes", func() {
			core.Push(ctx, DetailsPage)
		}),
	)
}

func DetailsPage(ctx *core.Context) core.View {
	//counter := core.NewState(ctx, 0)

	return core.Column(
		core.Text("📄 Detalhes", core.FontSize(22), core.FontWeight(core.Bold)),
		core.Spacer(10),
		core.Spacer(8),
		core.Button("⬅️ Voltar", func() {
			core.Pop(ctx)
		}),
	)
}

func SearchPage(ctx *core.Context) core.View {
	return core.Column(
		core.Text("🔍 Pesquisa", core.FontSize(24), core.FontWeight(core.Bold)),
		core.Input("", "Digite algo...", func(val string) {}),
	)
}

func ProfilePage(ctx *core.Context) core.View {
	return core.Column(
		core.Text("👤 Perfil", core.FontSize(24), core.FontWeight(core.Bold)),
		core.Text("Nome: Ismael GraHms", core.FontSize(16)),
		core.Text("Profissão: Engenheiro de Software"),
	)
}
