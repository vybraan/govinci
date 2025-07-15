package render

import (
	"encoding/json"
	"github.com/GraHms/govinci/core"
	"github.com/GraHms/govinci/reconcile"
)

type Manager struct {
	currentTree *core.Node
	context     *core.Context
	renderFunc  func(*core.Context) core.View
}

func New(ctx *core.Context, rootView func(*core.Context) core.View) *Manager {
	if ctx.Theme() == nil {
		ctx = ctx.WithTheme(core.DefaultTheme)
	}
	return &Manager{
		context:    ctx,
		renderFunc: rootView,
	}
}

func (r *Manager) RenderInitial() string {
	r.context.Reset()
	r.currentTree = r.renderFunc(r.context).Render(r.context)
	return renderJSON(r.currentTree)
}

// RenderAgain ReRender Used after an event (input/click/state change) to get diff
func (r *Manager) RenderAgain() string {
	r.context.Reset()
	newTree := r.renderFunc(r.context).Render(r.context)
	patches := reconcile.Diff(r.currentTree, newTree, "root")
	r.currentTree = newTree
	r.context.ClearDirty()
	core.PurgeUnusedCallbacks()
	return renderJSON(patches)
}

// JSON encoder
func renderJSON[T any](v T) string {
	data, err := json.Marshal(v)
	if err != nil {
		return `{"error":"failed to encode JSON"}`
	}
	return string(data)
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
