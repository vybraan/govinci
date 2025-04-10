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
	Button   Style
	Card     Style
	Input    Style
	Column   Style
	Row      Style
	Camera   Style
	CheckBox Style
	TextArea Style
	Text     Style
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

var DefaultTheme = &Theme{
	Colors: ColorPalette{
		Primary:       "#007AFF",   // iOS system blue
		Secondary:     "#34C759",   // iOS system green
		Background:    "#FFFFFF",   // white
		Surface:       "#F2F2F7",   // light gray
		TextPrimary:   "#000000",   // black
		TextSecondary: "#3C3C4399", // secondary label
		Error:         "#FF3B30",   // iOS system red
	},
	Typography: Typography{
		Title: Style{
			FontSize:   28,
			FontWeight: Bold,
			TextColor:  "#000000",
			Display:    DisplayBlock,
		},
		Subtitle: Style{
			FontSize:   22,
			FontWeight: Normal,
			TextColor:  "#3C3C4399",
			Display:    DisplayBlock,
		},
		Body: Style{
			FontSize:   17,
			FontWeight: Normal,
			TextColor:  "#000000",
			Display:    DisplayBlock,
		},
		Caption: Style{
			FontSize:   13,
			FontWeight: Normal,
			TextColor:  "#3C3C4399",
			Display:    DisplayBlock,
		},
	},
	Spacing: SpacingScale{
		XS: 4,
		SM: 8,
		MD: 16,
		LG: 24,
		XL: 32,
	},
	Components: ComponentDefaults{
		Button: Style{
			FontSize:     17,
			FontWeight:   Normal,
			TextColor:    "#FFFFFF",
			Background:   "#007AFF",
			Padding:      EdgeInsets{Top: 10, Bottom: 10, Left: 16, Right: 16},
			BorderRadius: 8,
			Shadow:       1,
			Align:        AlignCenter,
			Display:      DisplayInline,
		},
		Card: Style{
			Background:   "#FFFFFF",
			Padding:      EdgeInsets{Top: 16, Bottom: 16, Left: 16, Right: 16},
			Margin:       EdgeInsets{Top: 8, Bottom: 8, Left: 8, Right: 8},
			BorderRadius: 12,
			Shadow:       2,
			Display:      DisplayBlock,
		},
		Input: Style{
			FontSize:     17,
			FontWeight:   Normal,
			TextColor:    "#000000",
			Background:   "#FFFFFF",
			Padding:      EdgeInsets{Top: 8, Bottom: 8, Left: 12, Right: 12},
			BorderRadius: 6,
			Shadow:       0,
			Display:      DisplayBlock,
		},
		CheckBox: Style{
			Background:   "#FFFFFF",
			BorderRadius: 6,
			Shadow:       0,
			Display:      DisplayInline,
		},
		TextArea: Style{
			FontSize:     17,
			FontWeight:   Normal,
			TextColor:    "#000000",
			Background:   "#FFFFFF",
			Padding:      EdgeInsets{Top: 12, Bottom: 12, Left: 12, Right: 12},
			BorderRadius: 6,
			Display:      DisplayBlock,
		},

		Column: Style{
			Padding: EdgeInsets{Top: 12, Bottom: 12, Left: 16, Right: 16},
		},
		Row: Style{
			Padding: EdgeInsets{Top: 8, Bottom: 8, Left: 16, Right: 16},
		},
		Camera: Style{
			Background: "#000000",
			Display:    DisplayBlock,
		},
		Text: Style{
			FontSize:     17,
			FontWeight:   Normal,
			TextColor:    "#000000",
			Background:   "#FFFFFF",
			Padding:      EdgeInsets{Top: 12, Bottom: 12, Left: 12, Right: 12},
			BorderRadius: 6,
			Display:      DisplayBlock,
		},
	},
}

var MaterialTheme = &Theme{
	Colors: ColorPalette{
		Primary:       "#6200EE",
		Secondary:     "#03DAC6",
		Background:    "#FFFFFF",
		Surface:       "#F5F5F5",
		TextPrimary:   "#212121",
		TextSecondary: "#757575",
		Error:         "#B00020",
	},
	Typography: Typography{
		Title:    Style{FontSize: 22, FontWeight: Bold, TextColor: "#212121"},
		Subtitle: Style{FontSize: 18, FontWeight: Normal, TextColor: "#424242"},
		Body:     Style{FontSize: 14, FontWeight: Normal, TextColor: "#333333"},
		Caption:  Style{FontSize: 12, FontWeight: Light, TextColor: "#888888"},
	},
	Spacing: SpacingScale{
		XS: 4,
		SM: 8,
		MD: 16,
		LG: 24,
		XL: 32,
	},
	Components: ComponentDefaults{
		Button: Style{
			Background:   "#6200EE",
			TextColor:    "#FFFFFF",
			Padding:      EdgeInsets{Top: 10, Bottom: 10, Left: 20, Right: 20},
			BorderRadius: 4,
		},
		Card: Style{
			Background:   "#FFFFFF",
			BorderRadius: 8,
			Shadow:       1,
			Padding:      EdgeInsets{Top: 16, Bottom: 16, Left: 16, Right: 16},
		},
		Input: Style{
			Background: "#FAFAFA",
			Padding:    EdgeInsets{Top: 10, Bottom: 10, Left: 12, Right: 12},
		},
		Column: Style{
			Padding: EdgeInsets{Top: 12, Bottom: 12, Left: 16, Right: 16},
		},
		Row: Style{
			Padding: EdgeInsets{Top: 8, Bottom: 8, Left: 16, Right: 16},
		},
		Camera: Style{
			Background: "#000000",
			Display:    DisplayBlock,
		},

		CheckBox: Style{
			Display:      DisplayInline,
			Margin:       EdgeInsets{Right: 8},
			TextColor:    "#212121",
			BorderRadius: 2,
		},

		TextArea: Style{
			Background:   "#FAFAFA",
			TextColor:    "#212121",
			Padding:      EdgeInsets{Top: 8, Bottom: 8, Left: 12, Right: 12},
			BorderRadius: 4,
		},
	},
}
