package core

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

var (
	callbacks     = map[string]func(){}
	textCallbacks = map[string]func(string){}
	boolCallbacks = map[string]func(bool){}
	callbackMux   sync.Mutex
	counter       int
	textCounter   int
	boolCounter   int

	usedCallbacks = map[string]bool{}
)

func registerCallback(fn func()) string {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	id := fmt.Sprintf("cb_%d", counter)
	counter++
	callbacks[id] = fn
	usedCallbacks[id] = true
	return id
}

func TriggerCallback(id string) {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	if fn, ok := callbacks[id]; ok {
		fn()
		usedCallbacks[id] = true
	}
}

func registerTextCallback(fn func(string)) string {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	id := fmt.Sprintf("txt_cb_%d", textCounter)
	textCounter++
	textCallbacks[id] = fn
	usedCallbacks[id] = true
	return id
}

func TriggerTextCallback(id string, val string) {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	if fn, ok := textCallbacks[id]; ok {
		fn(val)
		usedCallbacks[id] = true
	}
}

func registerBoolCallback(fn func(bool)) string {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	id := fmt.Sprintf("bool_cb_%d", boolCounter)
	boolCounter++
	boolCallbacks[id] = fn
	usedCallbacks[id] = true
	return id
}

func TriggerBoolCallback(id string, val bool) {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	if fn, ok := boolCallbacks[id]; ok {
		fn(val)
		usedCallbacks[id] = true
	}
}

func ReceiveEventPayload(payload map[string]any) {
	id, ok := payload["callback"].(string)
	if !ok {
		log.Println("callback ID inválido")
		return
	}
	println("updating callback", id)

	switch val := payload["value"].(type) {
	case string:
		// Tenta deserializar
		var parsed map[string]any
		if err := json.Unmarshal([]byte(val), &parsed); err == nil {
			if v, ok := parsed["value"].(string); ok {
				TriggerTextCallback(id, v)
				return
			}
			if b, ok := parsed["value"].(bool); ok {
				TriggerBoolCallback(id, b)
				return
			}
		}

		// Fallback: trata como string normal
		TriggerTextCallback(id, val)

	case bool:
		TriggerBoolCallback(id, val)
	case nil:
		TriggerCallback(id)
	default:
		TriggerCallback(id)
	}
}

func PurgeUnusedCallbacks() {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	newCallbacks := make(map[string]func())
	newTextCallbacks := make(map[string]func(string))
	newBoolCallbacks := make(map[string]func(bool))

	for id, fn := range callbacks {
		if usedCallbacks[id] {
			newCallbacks[id] = fn
		}
	}
	for id, fn := range textCallbacks {
		if usedCallbacks[id] {
			newTextCallbacks[id] = fn
		}
	}
	for id, fn := range boolCallbacks {
		if usedCallbacks[id] {
			newBoolCallbacks[id] = fn
		}
	}

	callbacks = newCallbacks
	textCallbacks = newTextCallbacks
	boolCallbacks = newBoolCallbacks
	usedCallbacks = make(map[string]bool) // Clean up
}
