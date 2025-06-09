package core

type BehaviorProp interface {
	Apply(*Node)
}

type behaviorFunc func(*Node)

func (f behaviorFunc) Apply(n *Node) {
	f(n)
}
func On(event string, handler func()) BehaviorProp {
	return behaviorFunc(func(n *Node) {
		if n.Props == nil {
			n.Props = map[string]any{}
		}
		n.Props["on"+event] = registerCallback(handler)
	})
}

func OnClick(handler func()) BehaviorProp {
	return On("Click", handler)
}

func OnTouch(handler func()) BehaviorProp {
	return On("Touch", handler)
}
