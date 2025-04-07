package core

type View interface {
	Render(ctx *Context) *Node
}

type ComponentFunc func(ctx *Context) *Node

func (f ComponentFunc) Render(ctx *Context) *Node {
	return f(ctx)
}
