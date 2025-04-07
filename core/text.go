package core

func Text(content string, styleProps ...StyleProp) View {
	return ComponentFunc(func(ctx *Context) *Node {
		style := &Style{}
		for _, sp := range styleProps {
			sp.Apply(style)
		}

		return &Node{
			Type:  "Text",
			Props: map[string]any{"content": content},
			Style: style,
		}
	})
}
