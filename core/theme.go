package core

type Theme struct {
	Colors     ColorPalette
	Typography Typography
	Spacing    SpacingScale
	Components ComponentDefaults
}

type ColorPalette struct {
	Primary       string
	Secondary     string
	Background    string
	Surface       string
	TextPrimary   string
	TextSecondary string
	Error         string
}

type Typography struct {
	Title    Style
	Subtitle Style
	Body     Style
	Caption  Style
}

type SpacingScale struct {
	XS, SM, MD, LG, XL int
}

type ComponentDefaults struct {
	Button Style
	Card   Style
	Input  Style
}

func WithTheme(theme *Theme, children ...View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		newCtx := ctx.WithTheme(theme)
		var rendered []*Node
		for _, child := range children {
			rendered = append(rendered, child.Render(newCtx))
		}
		return &Node{
			Type:     "Theme",
			Props:    map[string]any{},
			Children: rendered,
		}
	})
}

var DefaultTheme = Theme{
	Colors: ColorPalette{
		Primary:       "#2196F3",
		Secondary:     "#FF4081",
		Background:    "#FFFFFF",
		Surface:       "#F5F5F5",
		TextPrimary:   "#212121",
		TextSecondary: "#757575",
		Error:         "#F44336",
	},
	Typography: Typography{
		Title:    Style{FontSize: 24, FontWeight: Bold, TextColor: "#212121"},
		Subtitle: Style{FontSize: 18, FontWeight: Normal, TextColor: "#424242"},
		Body:     Style{FontSize: 14, FontWeight: Normal, TextColor: "#616161"},
		Caption:  Style{FontSize: 12, FontWeight: Light, TextColor: "#9E9E9E"},
	},
	Spacing: SpacingScale{XS: 4, SM: 8, MD: 16, LG: 24, XL: 32},
	Components: ComponentDefaults{
		Button: Style{
			Background:   "#2196F3",
			TextColor:    "#FFFFFF",
			Padding:      EdgeInsets{Top: 8, Bottom: 8, Left: 16, Right: 16},
			BorderRadius: 6,
		},
	},
}
