//go:build js && wasm

package main

import (
	"encoding/json"
	"govinci/core"
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
	out := manager.RenderInitial()
	return js.ValueOf(out)
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

	return core.Column(
		core.Text("Bem-vindo ao Govinci"),
		core.Input(name.Get(), "Digite o seu nome", func(val string) {
			name.Set(val)
		}),
		core.Text("Olá, "+name.Get()),
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
