package core

import (
	"fmt"
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
)

func registerCallback(fn func()) string {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	id := fmt.Sprintf("cb_%d", counter)
	counter++
	callbacks[id] = fn
	return id
}

func TriggerCallback(id string) {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	if fn, ok := callbacks[id]; ok {
		fn()
	}
}

func registerTextCallback(fn func(string)) string {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	id := fmt.Sprintf("txt_cb_%d", textCounter)
	textCounter++
	textCallbacks[id] = fn
	return id
}

func TriggerTextCallback(id string, val string) {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	if fn, ok := textCallbacks[id]; ok {
		fn(val)
	}
}

func registerBoolCallback(fn func(bool)) string {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	id := fmt.Sprintf("bool_cb_%d", boolCounter)
	boolCounter++
	boolCallbacks[id] = fn
	return id
}

func ReceiveEventPayload(payload map[string]any) {
	if id, ok := payload["callback"].(string); ok {
		switch val := payload["value"].(type) {
		case string:
			TriggerTextCallback(id, val)
		case bool:
			TriggerBoolCallback(id, val)
		case nil:
			TriggerCallback(id)
		default:
			TriggerCallback(id)
		}
	}
}

func TriggerBoolCallback(id string, val bool) {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	if fn, ok := boolCallbacks[id]; ok {
		fn(val)
	}
}
