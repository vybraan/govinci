//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"govinci/core"
	"govinci/hooks"
	"govinci/render"
	"math/rand"
	"syscall/js"
	"time"
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
func isDirty(this js.Value, args []js.Value) any {
	return js.ValueOf(ctx.IsDirty())
}

func renderAgain(this js.Value, args []js.Value) any {
	out := manager.RenderAgain()
	return js.ValueOf(out)
}

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
	name := core.NewState(ctx, "")
	email := core.NewState(ctx, "")
	message := core.NewState(ctx, "")
	output := core.NewState(ctx, "")
	count := core.NewState(ctx, 0)

	// Atualiza o contador a cada segundo
	hooks.UseInterval(ctx, func() {
		count.Set(count.Get() + 1)
	}, time.Second)

	formField := func(label string, value string, placeholder string, onChange func(string)) core.View {
		return core.Column(
			core.Text(label, core.FontWeight(core.Bold), core.FontSize(14), core.Margin(4)),
			core.Input(value, placeholder, onChange,
				core.FontSize(16),
				core.Padding(10),
				core.Margin(4),
				core.BorderRadius(6),
				core.BackgroundColor("#F5F5F5"),
			),
			core.Spacer(12),
		)
	}

	return core.Row(
		core.Column(
			core.Text("Formulário de Contacto",
				core.FontSize(26),
				core.FontWeight(core.Bold),
				core.Margin(16),
				core.Align(core.AlignCenter),
			),

			core.Column(
				core.Image("https://example.com/avatar.jpg", core.UseStyle(core.Style{BorderRadius: 40})),
				formField("Nome Completo", name.Get(), "Digite o seu nome", name.Set),
				formField("Email", email.Get(), "Digite o seu email", email.Set),
				formField("Mensagem", message.Get(), "Digite a sua mensagem", message.Set),

				core.Button("Enviar", func() {
					summary := fmt.Sprintf("📬 Submetido:\nNome: %s\nEmail: %s\nMensagem: %s", name.Get(), email.Get(), message.Get())
					output.Set(summary)
				},
					core.Padding(12),
					core.FontSize(16),
					core.FontWeight(core.Bold),
					core.BackgroundColor("#007AFF"),
					core.TextColor("#FFFFFF"),
					core.BorderRadius(8),
					core.Shadow(2),
					core.Align(core.AlignCenter),
				),

				core.Spacer(20),

				core.Text(output.Get(),
					core.FontSize(15),
					core.TextColor("#2C3E50"),
					core.BackgroundColor("#ECF0F1"),
					core.Padding(12),
					core.Margin(8),
					core.BorderRadius(6),
					core.Shadow(1),
				),
				core.Padding(20),
				core.BackgroundColor("#FFFFFF"),
				core.BorderRadius(12),
				core.Shadow(2),
			),
			core.Align(core.AlignCenter),
		),
		core.Column(
			core.Text("⏱️ Temporizador Automático:"),
			core.Text(fmt.Sprintf("Contagem: %d", count.Get())),
		),
	)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	println("the generated name is", string(b))
	return string(b)
}
