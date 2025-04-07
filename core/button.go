package core

func Button(label string, onClick func(), styleProps ...StyleProp) View {
	return ComponentFunc(func(ctx *Context) *Node {
		style := &Style{}
		for _, sp := range styleProps {
			sp.Apply(style)
		}

		return &Node{
			Type: "Button",
			Props: map[string]any{
				"label":   label,
				"onClick": registerCallback(onClick),
			},
			Style: style,
		}
	})
}
