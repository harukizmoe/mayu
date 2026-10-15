package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// FABBuilder provides a fluent builder for Floating Action Buttons.
type FABBuilder struct {
	ui.Element
	c       *ui.Context
	icon    *ui.SVG
	label   string
	scheme  theme.ColorScheme
	density theme.Density
}

// NewFAB constructs a standard 56x56 Material 3 Floating Action Button.
func NewFAB(c *ui.Context, icon *ui.SVG, t *theme.Theme) *FABBuilder {
	b := ui.ButtonBase(c).Size(56, 56).Radius(theme.CornerLarge).Center()

	builder := &FABBuilder{
		Element: b,
		c:       c,
		icon:    icon,
		scheme:  t.Scheme,
		density: t.Density,
	}

	builder.style()
	if icon != nil {
		b.Children(func() {
			ui.Icon(c, icon).Size(24, 24)
		})
	}

	return builder
}

// NewExtendedFAB constructs an Extended FAB showing an icon and a label.
func NewExtendedFAB(c *ui.Context, icon *ui.SVG, label string, t *theme.Theme) *FABBuilder {
	b := ui.ButtonBase(c).Height(56).Radius(theme.CornerLarge).Padding(0, 20).Gap(12).Center()

	builder := &FABBuilder{
		Element: b,
		c:       c,
		icon:    icon,
		label:   label,
		scheme:  t.Scheme,
		density: t.Density,
	}

	builder.style()
	b.Children(func() {
		if icon != nil {
			ui.Icon(c, icon).Size(24, 24)
		}
		if label != "" {
			ui.Text(c, label).FontSize(theme.LabelLarge.Size).Bold().SingleLine()
		}
	})

	return builder
}

// OnClick registers an action callback on the FAB.
func (b *FABBuilder) OnClick(fn func()) *FABBuilder {
	if fn != nil && b.Element.Clicked() && !b.Element.IsDisabled() {
		fn()
	}
	return b
}

func (b *FABBuilder) style() {
	hovered := b.Element.Hovered()
	focused := b.Element.Focused()
	pressed := b.Element.Pressed()

	base := b.scheme.PrimaryContainer
	fg := b.scheme.OnPrimaryContainer
	bg := ApplyStateLayer(base, fg, hovered, focused, pressed)

	shadow := float32(3)
	if pressed {
		shadow = 3
	} else if hovered {
		shadow = 4
	}

	b.Element.Background(bg).TextColor(fg)
	b.Element.Shadow(0, shadow, shadow*3, shadow, b.scheme.Shadow.Alpha(0.20))
}
