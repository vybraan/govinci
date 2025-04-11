//go:build js && wasm

package main

import (
	"encoding/json"
	"govinci/core"
	. "govinci/examples/social"
	"govinci/hooks"
	"govinci/render"
	"syscall/js"
)

var (
	ctx = core.NewContext().WithTheme(core.DefaultTheme)
)

var manager *render.Manager

func renderInitial(this js.Value, args []js.Value) any {
	manager = render.New(ctx, App) // `App` é tua função de root view
	hooks.ClearIntervals()
	out := manager.RenderInitial()
	return js.ValueOf(out)
}
func RequestPermission(p Permission, onResult func(granted bool)) {
	js.Global().Call("GovinciRequestPermission", string(p), js.FuncOf(func(this js.Value, args []js.Value) any {
		granted := args[0].Bool()
		onResult(granted)
		return nil
	}))
}

func isDirty(this js.Value, args []js.Value) any {
	return js.ValueOf(ctx.IsDirty())
}

func renderAgain(this js.Value, args []js.Value) any {
	out := manager.RenderAgain()
	return js.ValueOf(out)
}

type Permission string

const (
	PermissionCamera      Permission = "camera"
	PermissionMicrophone  Permission = "microphone"
	PermissionGeolocation Permission = "geolocation"
)

func receiveEvent(this js.Value, args []js.Value) any {
	id := args[0].String()
	payloadStr := args[1].String()

	var payload map[string]any
	err := json.Unmarshal([]byte(payloadStr), &payload)
	if err != nil {
		println("Erro ao fazer parse do payload JSON:", err.Error())
		return nil
	}

	core.ReceiveEventPayload(map[string]any{
		"callback": id,
		"value":    payload["value"],
	})
	return nil
}

func registerCallbacks() {
	js.Global().Set("GovinciWASM", map[string]any{
		"RenderInitial": js.FuncOf(renderInitial),
		"RenderAgain":   js.FuncOf(renderAgain),
		"ReceiveEvent":  js.FuncOf(receiveEvent),
		"IsDirty":       js.FuncOf(isDirty),
	})
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	println("Govinci WASM ready.")
	<-c
}

func App(ctx *core.Context) core.View {
	return core.Navigator(func(ctx *core.Context) core.View {
		currentTab := core.NewState(ctx, "home")

		return core.Column(
			core.Match(currentTab.Get(),
				core.Case("home", HomePage(ctx)),
				core.Case("search", SearchPage(ctx)),
				core.Case("profile", ProfilePage(ctx)),
			),
			core.Row( // tab bar
				TabButton("🏠", "home", currentTab),
				TabButton("🔍", "search", currentTab),
				TabButton("👤", "profile", currentTab),
			),
		)
	})
}

func TabsComponent(ctx *core.Context, activeTab core.State[string]) core.View {
	tabButton := func(label, key string) core.View {
		active := activeTab.Get() == key
		return core.Button(label, func() {
			activeTab.Set(key)
		},
			core.Padding(10),
			core.Margin(4),
			core.BorderRadius(6),
			core.FontWeight(core.Bold),
			core.BackgroundColor(ifThen(active, "#007AFF", "#E0E0E0")),
			core.TextColor(ifThen(active, "#FFFFFF", "#000000")),
		)
	}

	return core.Column(
		core.Text("🗂️ Selecione uma aba:", core.FontSize(20), core.Margin(8)),

		core.Row(
			tabButton("Informações", "info"),
			tabButton("Configurações", "settings"),
			tabButton("Ajuda", "help"),
		),

		core.Spacer(16),

		core.Match(activeTab.Get(),
			core.Case("info", core.Text("📘 Esta é a aba de informações.")),
			core.Case("settings", core.Text("⚙️ Configurações do sistema.")),
			core.Case("help", core.Text("🆘 Ajuda e suporte técnico.")),
			core.Default[string](core.Text("❓ Aba desconhecida")),
		),
	)
}

func ifThen(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}

func HomeScreen(ctx *core.Context) core.View {
	return core.Column(
		core.Text("🏠 Tela Inicial", core.FontSize(22), core.FontWeight(core.Bold)),
		core.Spacer(12),
		core.Button("Ir para Detalhes", func() {
			core.Push(ctx, DetailsScreen)
		}),
	)
}

func DetailsScreen(ctx *core.Context) core.View {

	return core.Column(
		core.Text("📄 Tela de Detalhes", core.FontSize(20), core.FontWeight(core.Bold)),
		core.Spacer(8),
		core.Button("Incrementar", func() {

		}),
		core.Spacer(12),
		core.Button("⬅️ Voltar", func() {
			core.Pop(ctx)
		}),
	)
}
