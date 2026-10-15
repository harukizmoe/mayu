package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// CardVariant specifies the M3 Card visual style.
type CardVariant int

const (
	// CardElevated has a low surface container fill with subtle tonal elevation shadow.
	CardElevated CardVariant = iota
	// CardFilled has a higher contrast surface container fill without elevation.
	CardFilled
	// CardOutlined has a surface fill with a clean 1dp outline border.
	CardOutlined
)

// CardBuilder provides a fluent builder for Material 3 cards.
type CardBuilder struct {
	ui.Element
	c       *ui.Context
	variant CardVariant
	scheme  theme.ColorScheme
	density theme.Density
	onClick func()
}

// NewCard constructs a Material 3 Card container.
func NewCard(c *ui.Context, variant CardVariant, t *theme.Theme) *CardBuilder {
	box := ui.Box(c).Radius(theme.CornerMedium).Padding(16).Clip()

	builder := &CardBuilder{
		Element: box,
		c:       c,
		variant: variant,
		scheme:  t.Scheme,
		density: t.Density,
	}

	builder.style()
	return builder
}

// OnClick makes the card clickable and sets an action handler.
func (b *CardBuilder) OnClick(fn func()) *CardBuilder {
	b.onClick = fn
	b.Element.Focusable()
	if fn != nil && b.Element.Clicked() && !b.Element.IsDisabled() {
		fn()
	}
	b.style()
	return b
}

// Children populates the content of the card.
func (b *CardBuilder) Children(fn func()) *CardBuilder {
	if fn != nil {
		b.Element.Children(fn)
	}
	return b
}

// Padding sets custom padding on the card.
func (b *CardBuilder) Padding(v ...float32) *CardBuilder {
	b.Element.Padding(v...)
	return b
}

// Width sets an explicit width on the card.
func (b *CardBuilder) Width(w float32) *CardBuilder {
	b.Element.Width(w)
	return b
}

// Height sets an explicit height on the card.
func (b *CardBuilder) Height(h float32) *CardBuilder {
	b.Element.Height(h)
	return b
}

func (b *CardBuilder) style() {
	hovered := b.Element.Hovered()
	focused := b.Element.Focused()
	pressed := b.Element.Pressed()
	clickable := b.onClick != nil

	var bg ui.Color
	var border ui.Color
	var borderWidth float32
	var shadow float32

	switch b.variant {
	case CardElevated:
		base := b.scheme.SurfaceContainerLow
		if clickable {
			bg = ApplyStateLayer(base, b.scheme.OnSurface, hovered, focused, pressed)
			if pressed {
				shadow = 1
			} else if hovered {
				shadow = 2
			} else {
				shadow = 1
			}
		} else {
			bg = base
			shadow = 1
		}

	case CardFilled:
		base := b.scheme.SurfaceContainerHighest
		if clickable {
			bg = ApplyStateLayer(base, b.scheme.OnSurface, hovered, focused, pressed)
			if hovered && !pressed {
				shadow = 1
			}
		} else {
			bg = base
		}

	case CardOutlined:
		base := b.scheme.Surface
		if clickable {
			bg = ApplyStateLayer(base, b.scheme.OnSurface, hovered, focused, pressed)
		} else {
			bg = base
		}
		border = b.scheme.OutlineVariant
		borderWidth = 1
	}

	b.Element.Background(bg).TextColor(b.scheme.OnSurface)

	if borderWidth > 0 && border.A > 0 {
		b.Element.Border(borderWidth, border)
	} else {
		b.Element.Border(0, ui.Transparent)
	}

	if shadow > 0 {
		b.Element.Shadow(0, shadow, shadow*3, shadow, b.scheme.Shadow.Alpha(0.15))
	}
}
