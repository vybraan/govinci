package core

type Node struct {
	Type     string
	Props    map[string]any
	Style    *Style
	Children []*Node
}
