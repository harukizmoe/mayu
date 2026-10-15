package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// IconButtonVariant defines the emphasis style of an IconButton.
type IconButtonVariant int

const (
	IconButtonStandard IconButtonVariant = iota
	IconButtonFilled
	IconButtonFilledTonal
	IconButtonOutlined
)

// IconButtonBuilder provides a fluent builder for IconButtons.
type IconButtonBuilder struct {
	ui.Element
	c       *ui.Context
	icon    *ui.SVG
	variant IconButtonVariant
	scheme  theme.ColorScheme
	density theme.Density
}

// NewIconButton constructs a Material 3 circular Icon Button (40x40 standard).
func NewIconButton(c *ui.Context, icon *ui.SVG, variant IconButtonVariant, t *theme.Theme) *IconButtonBuilder {
	b := ui.ButtonBase(c).Size(40, 40).Radius(20).Center()

	builder := &IconButtonBuilder{
		Element: b,
		c:       c,
		icon:    icon,
		variant: variant,
		scheme:  t.Scheme,
		density: t.Density,
	}

	builder.style()
	if icon != nil {
		b.Children(func() {
			ui.Icon(c, icon).Size(20, 20)
		})
	}

	return builder
}

// OnClick registers an action callback on the icon button.
func (b *IconButtonBuilder) OnClick(fn func()) *IconButtonBuilder {
	if fn != nil && b.Element.Clicked() && !b.Element.IsDisabled() {
		fn()
	}
	return b
}

func (b *IconButtonBuilder) style() {
	hovered := b.Element.Hovered()
	focused := b.Element.Focused()
	pressed := b.Element.Pressed()
	disabled := b.Element.IsDisabled()

	if disabled {
		b.Element.Background(ui.Transparent).TextColor(b.scheme.OnSurface.Alpha(DisabledContentAlpha))
		return
	}

	var bg ui.Color
	var fg ui.Color
	var border ui.Color
	var borderWidth float32

	switch b.variant {
	case IconButtonFilled:
		base := b.scheme.Primary
		fg = b.scheme.OnPrimary
		bg = ApplyStateLayer(base, fg, hovered, focused, pressed)
	case IconButtonFilledTonal:
		base := b.scheme.SecondaryContainer
		fg = b.scheme.OnSecondaryContainer
		bg = ApplyStateLayer(base, fg, hovered, focused, pressed)
	case IconButtonOutlined:
		fg = b.scheme.OnSurfaceVariant
		bg = ApplyStateLayer(ui.Transparent, fg, hovered, focused, pressed)
		border = b.scheme.Outline
		borderWidth = 1
	default: // Standard
		fg = b.scheme.OnSurfaceVariant
		bg = ApplyStateLayer(ui.Transparent, fg, hovered, focused, pressed)
	}

	b.Element.Background(bg).TextColor(fg)
	if borderWidth > 0 && border.A > 0 {
		b.Element.Border(borderWidth, border)
	} else {
		b.Element.Border(0, ui.Transparent)
	}
}
