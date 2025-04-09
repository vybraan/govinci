package core

import (
	"fmt"
	"sync"
)

type RenderManager struct {
	mu      sync.Mutex
	subs    map[string]func()
	counter int
}

func NewRenderManager() *RenderManager {
	return &RenderManager{
		subs: make(map[string]func()),
	}
}

func (r *RenderManager) RegisterRender(fn func()) string {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := fmt.Sprintf("render_%d", r.counter)
	r.counter++
	r.subs[id] = fn
	return id
}

func (r *RenderManager) TriggerRender(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if fn, ok := r.subs[id]; ok {
		go fn() // async
	}
}

func (ctx *Context) SubscribeRender(fn func()) {
	ctx.renderManager.RegisterRender(fn)
}
