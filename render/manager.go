package render

import (
	"encoding/json"
	"govinci/core"
	"govinci/reconcile"
)

type Manager struct {
	currentTree *core.Node
	context     *core.Context
	renderFunc  func(*core.Context) core.View
}

func New(ctx *core.Context, rootView func(*core.Context) core.View) *Manager {
	return &Manager{
		context:    ctx,
		renderFunc: rootView,
	}
}

func (r *Manager) RenderAndGetPatches() string {
	r.context.Cursor = 0
	newTree := r.renderFunc(r.context).Render(r.context)

	if r.currentTree == nil {
		r.currentTree = newTree
		return render(newTree)
	}

	patches := reconcile.Diff(r.currentTree, newTree, "root")
	r.currentTree = newTree
	return render(patches)
}

func render[T any](tree T) string {
	data, err := json.Marshal(tree)
	if err != nil {
		return `{"error":"failed to encode render tree"}`
	}
	return string(data)
}
