package core

import (
	"log"
	"sync"
)

type Context struct {
	slots           []any
	Cursor          int
	theme           *Theme
	config          *AppConfig
	idGen           int
	lock            sync.Mutex
	renderManager   *RenderManager
	callbackMap     map[string]any // Stable ID to callback
	callbackCounter map[string]int // Per-type counter
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
		slots:           make([]any, 0),
		Cursor:          0,
		renderManager:   NewRenderManager(),
		callbackMap:     make(map[string]any),
		callbackCounter: make(map[string]int),
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
		slots:           ctx.slots,
		Cursor:          ctx.Cursor,
		theme:           ctx.theme,
		config:          cfg,
		renderManager:   ctx.renderManager,
		callbackMap:     ctx.callbackMap,
		callbackCounter: ctx.callbackCounter,
	}
}

func (ctx *Context) WithTheme(theme *Theme) *Context {
	return &Context{
		slots:           ctx.slots,
		Cursor:          ctx.Cursor,
		theme:           theme,
		config:          ctx.config,
		renderManager:   ctx.renderManager,
		callbackMap:     ctx.callbackMap,
		callbackCounter: ctx.callbackCounter,
	}
}

func NewState[T any](ctx *Context, initial T) State[T] {
	log.Printf("NewState at Cursor: %d (len=%d)", ctx.Cursor, len(ctx.slots))
	if ctx.Cursor >= len(ctx.slots) {
		log.Printf("Allocating slot %d with value: %#v", ctx.Cursor, initial)
		ctx.slots = append(ctx.slots, initial)
	}

	index := ctx.Cursor
	ctx.Cursor++

	return State[T]{
		get: func() T {
			return ctx.slots[index].(T)
		},
		set: func(val T) {
			log.Printf("Updating slot %d with value: %#v", index, val)
			ctx.slots[index] = val
			ctx.renderManager.TriggerRender("default")
		},
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

func (ctx *Context) RegisterCallback(key string, fn any) string {
	typeName := callbackType(fn)
	ctx.lock.Lock()
	defer ctx.lock.Unlock()

	if ctx.callbackCounter == nil {
		ctx.callbackCounter = make(map[string]int)
	}

	id := typeName + ":" + key
	ctx.callbackMap[id] = fn
	ctx.callbackCounter[typeName]++
	return id
}

func callbackType(fn any) string {
	switch fn.(type) {
	case func():
		return "cb"
	case func(string):
		return "txt_cb"
	case func(bool):
		return "bool_cb"
	default:
		return "unknown"
	}
}
