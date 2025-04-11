package core

var navigatorStack = make([]func(*Context) View, 0)

func Navigator(initial func(*Context) View) View {
	if len(navigatorStack) == 0 {
		navigatorStack = append(navigatorStack, initial)
	}
	return ComponentFunc(func(ctx *Context) *Node {
		current := navigatorStack[len(navigatorStack)-1]
		return Render(ctx, current(ctx))
	})
}

func Push(ctx *Context, route func(*Context) View) {
	navigatorStack = append(navigatorStack, route)
	ctx.MarkDirty()
}

func Pop(ctx *Context) {
	if len(navigatorStack) > 1 {
		navigatorStack = navigatorStack[:len(navigatorStack)-1]
		ctx.MarkDirty()
	}
}

func Replace(ctx *Context, route func(*Context) View) {
	if len(navigatorStack) > 0 {
		navigatorStack[len(navigatorStack)-1] = route
		ctx.MarkDirty()
	}
}

func Reset(ctx *Context, route func(*Context) View) {
	navigatorStack = []func(*Context) View{route}
	ctx.MarkDirty()
}

func Render(ctx *Context, view View) *Node {
	ctx.Reset()
	return view.Render(ctx)
}
