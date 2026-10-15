package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// TextFieldVariant specifies whether the text field is Filled or Outlined.
type TextFieldVariant int

const (
	// TextFieldFilled uses a SurfaceContainerHighest background with a bottom active indicator.
	TextFieldFilled TextFieldVariant = iota
	// TextFieldOutlined uses an outline border and floating label header.
	TextFieldOutlined
)

// TextFieldBuilder provides a fluent builder for Material 3 text inputs.
type TextFieldBuilder struct {
	ui.Element
	c       *ui.Context
	variant TextFieldVariant
	value   *string
	input   ui.Element
	scheme  theme.ColorScheme
	density theme.Density
}

// NewTextField constructs a Material 3 Text Field.
func NewTextField(c *ui.Context, value *string, label string, variant TextFieldVariant, t *theme.Theme) *TextFieldBuilder {
	col := ui.Column(c).Gap(4).FillWidth()

	builder := &TextFieldBuilder{
		Element: col,
		c:       c,
		variant: variant,
		value:   value,
		scheme:  t.Scheme,
		density: t.Density,
	}

	var bg ui.Color
	var border ui.Color
	var borderWidth float32

	if variant == TextFieldFilled {
		bg = t.Scheme.SurfaceContainerHighest
		border = t.Scheme.Outline
		borderWidth = 1
	} else {
		bg = ui.Transparent
		border = t.Scheme.Outline
		borderWidth = 1
	}

	col.Children(func() {
		inputBox := ui.Row(c).Height(56).AlignItems(ui.Center).Padding(8, 16).Gap(12).FillWidth().Background(bg)

		if variant == TextFieldFilled {
			inputBox.Radius(4, 4, 0, 0)
			inputBox.BorderWidth(0, 0, borderWidth, 0).BorderColor(border)
		} else {
			inputBox.Radius(theme.CornerSmall)
			inputBox.Border(borderWidth, border)
		}

		inputBox.Children(func() {
			ui.Column(c).Grow(1).Justify(ui.Center).Children(func() {
				if label != "" {
					ui.Text(c, label).FontSize(11).TextColor(t.Scheme.Primary).SingleLine()
				}
				builder.input = ui.TextInputBase(c, value).
					FillWidth().
					TextColor(t.Scheme.OnSurface)
			})
		})
	})

	return builder
}

// Placeholder sets the input placeholder when empty.
func (b *TextFieldBuilder) Placeholder(p string) *TextFieldBuilder {
	b.input.Placeholder(p)
	return b
}

// Disabled marks the text field as disabled.
func (b *TextFieldBuilder) Disabled(d bool) *TextFieldBuilder {
	b.input.Disabled(d)
	return b
}
