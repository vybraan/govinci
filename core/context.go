package core

import "log"

type Context struct {
	stateSlots []any
	cursor     int
	theme      *Theme
	config     *AppConfig
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
		stateSlots: make([]any, 0),
		cursor:     0,
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
	return &DefaultTheme // fallback
}

func (ctx *Context) Config() *AppConfig {
	if ctx.config == nil {
		return &AppConfig{}
	}
	return ctx.config
}

func (ctx *Context) WithConfig(cfg *AppConfig) *Context {
	return &Context{
		stateSlots: ctx.stateSlots,
		cursor:     ctx.cursor,
		theme:      ctx.theme,
		config:     cfg,
	}
}

func (ctx *Context) WithTheme(theme *Theme) *Context {
	return &Context{
		stateSlots: ctx.stateSlots,
		cursor:     ctx.cursor,
		theme:      theme,
	}
}

func NewState[T any](ctx *Context, initial T) State[T] {
	log.Printf("NewState at cursor: %d (len=%d)", ctx.cursor, len(ctx.stateSlots))
	if ctx.cursor >= len(ctx.stateSlots) {
		log.Printf("Allocating slot %d with value: %#v", ctx.cursor, initial)
		ctx.stateSlots = append(ctx.stateSlots, initial)
	}

	index := ctx.cursor
	ctx.cursor++

	return State[T]{
		get: func() T {
			return ctx.stateSlots[index].(T)
		},
		set: func(val T) {
			log.Printf("Updating slot %d with value: %#v", index, val)
			ctx.stateSlots[index] = val
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
