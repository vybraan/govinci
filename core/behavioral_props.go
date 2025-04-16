package core

type BehaviorProp interface {
	Apply(*Node)
}

type behaviorFunc func(*Node)

func (f behaviorFunc) Apply(n *Node) {
	f(n)
}
func OnClick(handler func()) BehaviorProp {
	return behaviorFunc(func(n *Node) {
		n.Props["onClick"] = registerCallback(handler)
	})
}
