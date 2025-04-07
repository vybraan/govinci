package htmlout

import (
	"fmt"
	"govinci/core"
	"strings"
)

func ExportHTML(node *core.Node) string {
	var builder strings.Builder
	renderNode(&builder, node, 0)
	return builder.String()
}

func renderNode(b *strings.Builder, node *core.Node, indent int) {
	pad := strings.Repeat("  ", indent)

	tag := tagForType(node.Type)
	if tag == "" {
		return
	}

	attrs := styleAttr(node.Style)

	// Props: data-onclick
	if id, ok := node.Props["onClick"].(string); ok {
		attrs += fmt.Sprintf(" data-onclick=\"%s\"", id)
	}

	// Open tag
	b.WriteString(fmt.Sprintf("%s<%s%s>", pad, tag, attrs))

	// Content
	switch node.Type {
	case "Text":
		if content, ok := node.Props["content"].(string); ok {
			b.WriteString(content)
		}
	case "Image":
		if src, ok := node.Props["src"].(string); ok {
			b.WriteString(fmt.Sprintf("<img src=\"%s\"/>", src))
		}
	case "Button":
		if label, ok := node.Props["label"].(string); ok {
			b.WriteString(label) // add label inside button
		}
	case "Input":
		b.WriteString("<input")
		if val, ok := node.Props["value"].(string); ok {
			b.WriteString(fmt.Sprintf(" value=\"%s\"", val))
		}
		if ph, ok := node.Props["placeholder"].(string); ok {
			b.WriteString(fmt.Sprintf(" placeholder=\"%s\"", ph))
		}
		if id, ok := node.Props["onChange"].(string); ok {
			b.WriteString(fmt.Sprintf(" data-onchange=\"%s\"", id))
		}
		b.WriteString(" />")
	case "Checkbox":
		b.WriteString("<input type=\"checkbox\"")
		if val, ok := node.Props["checked"].(bool); ok && val {
			b.WriteString(" checked")
		}
		if id, ok := node.Props["onToggle"].(string); ok {
			b.WriteString(fmt.Sprintf(" data-ontoggle=\"%s\"", id))
		}
		b.WriteString(" />")

	}
	if node.Type == "Spacer" {
		if size, ok := node.Props["size"].(int); ok {
			b.WriteString(fmt.Sprintf("<div style=\"height:%dpx\"></div>", size))
			return
		}
	}

	// Children
	for _, child := range node.Children {
		b.WriteString("\n")
		renderNode(b, child, indent+1)
	}

	// Close tag
	b.WriteString(fmt.Sprintf("</%s>\n", tag))
}

func tagForType(t string) string {
	switch t {
	case "Text":
		return "span"
	case "Column":
		return "div"
	case "Row":
		return "div"
	case "Button":
		return "button"
	case "Image":
		return "div"
	case "Card":
		return "div"
	case "Spacer":
		return "div"
	default:
		return "div"
	}
}

func styleAttr(s *core.Style) string {
	if s == nil {
		return ""
	}
	styles := []string{}
	if s.TextColor != "" {
		styles = append(styles, fmt.Sprintf("color:%s", s.TextColor))
	}
	if s.Background != "" {
		styles = append(styles, fmt.Sprintf("background:%s", s.Background))
	}
	if s.FontSize != 0 {
		styles = append(styles, fmt.Sprintf("font-size:%dpx", s.FontSize))
	}
	if s.Align != "" {
		switch s.Align {
		case core.AlignCenter:
			styles = append(styles, "text-align:center")
		case core.AlignStart:
			styles = append(styles, "text-align:left")
		case core.AlignEnd:
			styles = append(styles, "text-align:right")
		}
	}
	if s.Display != "" {
		styles = append(styles, fmt.Sprintf("display:%s", s.Display))
	}
	if s.Padding != (core.EdgeInsets{}) {
		styles = append(styles, fmt.Sprintf("padding:%dpx %dpx %dpx %dpx", s.Padding.Top, s.Padding.Right, s.Padding.Bottom, s.Padding.Left))
	}
	if s.Margin != (core.EdgeInsets{}) {
		styles = append(styles, fmt.Sprintf("margin:%dpx %dpx %dpx %dpx", s.Margin.Top, s.Margin.Right, s.Margin.Bottom, s.Margin.Left))
	}
	if s.BorderRadius != 0 {
		styles = append(styles, fmt.Sprintf("border-radius:%dpx", s.BorderRadius))
	}
	if len(styles) == 0 {
		return ""
	}
	return fmt.Sprintf(" style=\"%s\"", strings.Join(styles, "; "))
}
