package core

type Style struct {
	FontSize     int
	FontWeight   Weight
	TextColor    string
	Background   string
	Padding      EdgeInsets
	Margin       EdgeInsets
	BorderRadius int
	Shadow       int
	Align        Alignment
	Display      DisplayMode
}
type Weight int

const (
	Light  Weight = 200
	Normal Weight = 400
	Bold   Weight = 700
)

type EdgeInsets struct {
	Top, Right, Bottom, Left int
	Horizontal, Vertical     int
}
type StyleProp interface {
	Apply(*Style)
}
type styleFunc func(*Style)

func (f styleFunc) Apply(s *Style) {
	f(s)
}
func FontSize(size int) StyleProp {
	return styleFunc(func(s *Style) {
		s.FontSize = size
	})
}

func TextColor(hex string) StyleProp {
	return styleFunc(func(s *Style) {
		s.TextColor = hex
	})
}

func BackgroundColor(hex string) StyleProp {
	return styleFunc(func(s *Style) {
		s.Background = hex
	})
}
func Align(a Alignment) StyleProp {
	return styleFunc(func(s *Style) {
		s.Align = a
	})
}

func Display(mode DisplayMode) StyleProp {
	return styleFunc(func(s *Style) {
		s.Display = mode
	})
}

func Padding(all int) StyleProp {
	return styleFunc(func(s *Style) {
		s.Padding = EdgeInsets{
			Top: all, Right: all, Bottom: all, Left: all,
		}
	})
}
func BorderRadius(px int) StyleProp {
	return styleFunc(func(s *Style) {
		s.BorderRadius = px
	})
}

func Shadow(elevation int) StyleProp {
	return styleFunc(func(s *Style) {
		s.Shadow = elevation
	})
}

func FontWeight(weight Weight) StyleProp {
	return styleFunc(func(s *Style) {
		s.FontWeight = weight
	})
}

func Margin(all int) StyleProp {
	return styleFunc(func(s *Style) {
		s.Margin = EdgeInsets{
			Top: all, Right: all, Bottom: all, Left: all,
		}
	})
}

func PaddingVertical(px int) StyleProp {
	return styleFunc(func(s *Style) {
		s.Padding.Vertical = px
	})
}
func (s Style) With(other Style) Style {
	merged := s
	UseStyle(other).Apply(&merged)
	return merged
}

func UseStyle(s Style) StyleProp {
	return styleFunc(func(target *Style) {
		// Merge only non-zero fields
		if s.FontSize != 0 {
			target.FontSize = s.FontSize
		}
		if s.FontWeight != 0 {
			target.FontWeight = s.FontWeight
		}
		if s.TextColor != "" {
			target.TextColor = s.TextColor
		}
		if s.Background != "" {
			target.Background = s.Background
		}
		if s.BorderRadius != 0 {
			target.BorderRadius = s.BorderRadius
		}
		if s.Shadow != 0 {
			target.Shadow = s.Shadow
		}
		if s.Align != "" {
			target.Align = s.Align
		}
		if s.Display != "" {
			target.Display = s.Display
		}
		if s.Padding != (EdgeInsets{}) {
			target.Padding = s.Padding
		}
		if s.Margin != (EdgeInsets{}) {
			target.Margin = s.Margin
		}
	})
}
func PrimaryColor() string { return "#007AFF" }
func DangerColor() string  { return "#FF3B30" }
func RoundedShadowBox() StyleProp {
	return UseStyle(Style{
		BorderRadius: 12,
		Shadow:       2,
		Background:   "#FFFFFF",
	})
}

var TextInputStyle = UseStyle(Style{
	FontSize:     16,
	TextColor:    "#000000",
	Background:   "#FFFFFF",
	Padding:      EdgeInsets{Top: 10, Bottom: 10, Left: 12, Right: 12},
	BorderRadius: 8,
	Shadow:       1,
})

func PaddingTop(px int) StyleProp {
	return styleFunc(func(s *Style) {
		s.Padding.Top = px
	})
}

func PaddingHorizontal(px int) StyleProp {
	return styleFunc(func(s *Style) {
		s.Padding.Horizontal = px
	})
}

type ResponsiveStyle map[string]Style // "mobile", "tablet", "desktop"

type Alignment string

const (
	AlignStart    Alignment = "start"
	AlignCenter   Alignment = "center"
	AlignEnd      Alignment = "end"
	AlignStretch  Alignment = "stretch"
	AlignBaseline Alignment = "baseline"
	AlignJustify  Alignment = "justify"
)

type DisplayMode string

const (
	DisplayVisible DisplayMode = "visible"
	DisplayHidden  DisplayMode = "hidden"
	DisplayNone    DisplayMode = "none"
	DisplayInline  DisplayMode = "inline"
	DisplayBlock   DisplayMode = "block"
)
