package core

func Row(children ...View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		var nodes []*Node
		for _, child := range children {
			nodes = append(nodes, child.Render(ctx))
		}
		return &Node{
			Type:     "Row",
			Props:    map[string]any{},
			Children: nodes,
		}
	})
}

func Card(children ...View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		var nodes []*Node
		for _, child := range children {
			nodes = append(nodes, child.Render(ctx))
		}
		return &Node{
			Type:     "Card",
			Props:    map[string]any{},
			Style:    &ctx.Theme().Components.Card,
			Children: nodes,
		}
	})
}

func Spacer(size int) View {
	return ComponentFunc(func(ctx *Context) *Node {
		return &Node{
			Type: "Spacer",
			Props: map[string]any{
				"size": size,
			},
		}
	})
}

func Scroll(children ...View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		var nodes []*Node
		for _, child := range children {
			nodes = append(nodes, child.Render(ctx))
		}
		return &Node{
			Type:     "Scroll",
			Props:    map[string]any{},
			Children: nodes,
		}
	})
}

func SafeArea(child View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		return &Node{
			Type:     "SafeArea",
			Props:    map[string]any{},
			Children: []*Node{child.Render(ctx)},
		}
	})
}

func Fragment(children ...View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		if len(children) == 1 {
			return children[0].Render(ctx)
		}
		return &Node{
			Type:     "Fragment",
			Children: renderAll(ctx, children),
		}
	})
}
func Column(children ...View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		var nodes []*Node
		for _, child := range children {
			nodes = append(nodes, child.Render(ctx))
		}

		return &Node{
			Type:     "Column",
			Props:    map[string]any{}, // No specific props yet
			Children: nodes,
		}
	})
}

func renderAll(ctx *Context, views []View) []*Node {
	var out []*Node
	for _, v := range views {
		out = append(out, v.Render(ctx))
	}
	return out
}
