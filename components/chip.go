package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// ChipBuilder provides a fluent builder for Material 3 Chips.
type ChipBuilder struct {
	ui.Element
	c        *ui.Context
	label    string
	icon     *ui.SVG
	selected bool
	scheme   theme.ColorScheme
	density  theme.Density
}

// NewChip constructs an M3 Chip (32dp height, 8dp radius).
func NewChip(c *ui.Context, label string, t *theme.Theme) *ChipBuilder {
	b := ui.ButtonBase(c).Height(32).Radius(theme.CornerSmall).Padding(0, 12).Gap(8).Center()

	builder := &ChipBuilder{
		Element: b,
		c:       c,
		label:   label,
		scheme:  t.Scheme,
		density: t.Density,
	}

	builder.style()
	if label != "" {
		b.Label(label)
		b.Children(func() {
			ui.Text(c, label).FontSize(theme.LabelLarge.Size).SingleLine()
		})
	}

	return builder
}

// Icon adds a leading icon to the chip.
func (b *ChipBuilder) Icon(s *ui.SVG) *ChipBuilder {
	if s == nil {
		return b
	}
	b.icon = s
	b.Element.Reverse()
	b.Element.Padding(0, 12, 0, 8)
	b.Element.Children(func() {
		ui.Icon(b.c, s).Size(18, 18)
	})
	return b
}

// Selected sets the selected state for filter chips.
func (b *ChipBuilder) Selected(sel bool) *ChipBuilder {
	b.selected = sel
	b.style()
	if sel {
		b.Element.Reverse()
		b.Element.Children(func() {
			ui.Icon(b.c, symbols.Check).Size(18, 18)
		})
	}
	return b
}

// OnClick registers an action callback on the chip.
func (b *ChipBuilder) OnClick(fn func()) *ChipBuilder {
	if fn != nil && b.Element.Clicked() && !b.Element.IsDisabled() {
		fn()
	}
	return b
}

func (b *ChipBuilder) style() {
	hovered := b.Element.Hovered()
	focused := b.Element.Focused()
	pressed := b.Element.Pressed()

	var bg ui.Color
	var fg ui.Color
	var border ui.Color
	var borderWidth float32

	if b.selected {
		base := b.scheme.SecondaryContainer
		fg = b.scheme.OnSecondaryContainer
		bg = ApplyStateLayer(base, fg, hovered, focused, pressed)
	} else {
		fg = b.scheme.OnSurface
		bg = ApplyStateLayer(ui.Transparent, fg, hovered, focused, pressed)
		border = b.scheme.OutlineVariant
		borderWidth = 1
	}

	b.Element.Background(bg).TextColor(fg)
	if borderWidth > 0 && border.A > 0 {
		b.Element.Border(borderWidth, border)
	} else {
		b.Element.Border(0, ui.Transparent)
	}
}
