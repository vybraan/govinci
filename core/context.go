package core

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Context struct {
	slots         []any
	Cursor        int
	theme         *Theme
	config        *AppConfig
	idGen         int
	lock          sync.Mutex
	renderManager *RenderManager
	contextID     string
}

type AppConfig struct {
	Name        string
	Description string
	Version     string
	Locale      string
	Author      string
	Meta        map[string]string
}

func NewContext() *Context {
	return &Context{
		slots:         make([]any, 0),
		Cursor:        0,
		renderManager: NewRenderManager(),
		contextID:     fmt.Sprintf("ctx_%d", time.Now().UnixNano()),
	}
}

func NewState[T any](ctx *Context, initial T) State[T] {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()

	slotID := ctx.Cursor
	if slotID >= len(ctx.slots) {
		ctx.slots = append(ctx.slots, initial)
	}
	ctx.Cursor++

	return State[T]{
		get: func() T {
			return ctx.slots[slotID].(T)
		},
		set: func(val T) {
			ctx.slots[slotID] = val
			log.Printf("[State] Updating slot %d with value: %#v", slotID, val)
			// Trigger re-render
			ctx.renderManager.TriggerRender(ctx.contextID)
		},
	}
}

type State[T any] struct {
	get func() T
	set func(T)
}

func (s *State[T]) Get() T {
	return s.get()
}

func (s *State[T]) Set(val T) {
	s.set(val)
}

func (ctx *Context) Theme() *Theme {
	if ctx.theme != nil {
		return ctx.theme
	}
	return DefaultTheme // fallback
}

func (ctx *Context) Config() *AppConfig {
	if ctx.config == nil {
		return &AppConfig{}
	}
	return ctx.config
}

func (ctx *Context) WithConfig(cfg *AppConfig) *Context {
	return &Context{
		slots:  ctx.slots,
		Cursor: ctx.Cursor,
		theme:  ctx.theme,
		config: cfg,
	}
}

func (ctx *Context) WithTheme(theme *Theme) *Context {
	return &Context{
		slots:  ctx.slots,
		Cursor: ctx.Cursor,
		theme:  theme,
	}
}

func (ctx *Context) With(opts ...func(*Context)) *Context {
	for _, fn := range opts {
		fn(ctx)
	}
	return ctx
}

func WithThemeOpt(t *Theme) func(*Context) {
	return func(ctx *Context) {
		ctx.theme = t
	}
}

func WithConfigOpt(c *AppConfig) func(*Context) {
	return func(ctx *Context) {
		ctx.config = c
	}
}
