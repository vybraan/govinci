package hooks

import (
	"github.com/GraHms/govinci/core"
	"reflect"
)

var previousDeps = map[int][]any{}
var effectIndex = 0

func UseEffect(ctx *core.Context, effect func(), deps ...any) {
	index := effectIndex
	effectIndex++

	prev, exists := previousDeps[index]
	shouldRun := !exists || !reflect.DeepEqual(prev, deps)

	if shouldRun {
		previousDeps[index] = deps
		go effect() // roda em background, pode adaptar conforme necessário
	}
}

func ResetEffects() {
	effectIndex = 0
}
