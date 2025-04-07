package core

func Input(value string, placeholder string, onChange func(string), styleProps ...StyleProp) View {
	return ComponentFunc(func(ctx *Context) *Node {
		style := &Style{}
		for _, sp := range styleProps {
			sp.Apply(style)
		}

		id := registerTextCallback(onChange)

		return &Node{
			Type: "Input",
			Props: map[string]any{
				"value":       value,
				"placeholder": placeholder,
				"onChange":    id,
			},
			Style: style,
		}
	})
}

func Checkbox(checked bool, onToggle func(bool), styleProps ...StyleProp) View {
	return ComponentFunc(func(ctx *Context) *Node {
		style := &Style{}
		for _, sp := range styleProps {
			sp.Apply(style)
		}

		id := registerBoolCallback(onToggle)

		return &Node{
			Type: "Checkbox",
			Props: map[string]any{
				"checked":  checked,
				"onToggle": id,
			},
			Style: style,
		}
	})
}
