package core

func Image(src string, styleProps ...StyleProp) View {
	return ComponentFunc(func(ctx *Context) *Node {
		style := &Style{}
		for _, sp := range styleProps {
			sp.Apply(style)
		}

		return &Node{
			Type: "Image",
			Props: map[string]any{
				"src": src,
			},
			Style: style,
		}
	})
}
