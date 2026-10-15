package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// TextFieldVariant specifies whether a field is Filled or Outlined.
type TextFieldVariant int

const (
	// TextFieldFilled uses a SurfaceContainerHighest background with a bottom
	// active indicator.
	TextFieldFilled TextFieldVariant = iota
	// TextFieldOutlined uses an outline border.
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

// NewTextField constructs a Material 3 Text Field: one line, 56 DIPs tall.
func NewTextField(c *ui.Context, value *string, label string, variant TextFieldVariant, t *theme.Theme) *TextFieldBuilder {
	return newField(c, value, label, variant, false, t)
}

// NewTextArea constructs a Material 3 Text Area: the same field, several lines
// tall, holding its label at the top like a filled field does.
func NewTextArea(c *ui.Context, value *string, label string, variant TextFieldVariant, t *theme.Theme) *TextFieldBuilder {
	return newField(c, value, label, variant, true, t)
}

// newField builds the shape the two share, which differs only in its height,
// the input it fills and where the label sits.
func newField(c *ui.Context, value *string, label string, variant TextFieldVariant, multiline bool, t *theme.Theme) *TextFieldBuilder {
	const (
		lineHeight = 56
		areaHeight = 120
	)
	height := float32(lineHeight)
	if multiline {
		height = areaHeight
	}

	builder := &TextFieldBuilder{
		Element: ui.Column(c).Gap(4).FillWidth(),
		c:       c,
		variant: variant,
		value:   value,
		scheme:  t.Scheme,
		density: t.Density,
	}

	bg, border := ui.Transparent, t.Scheme.Outline

	builder.Element.Children(func() {
		box := ui.Row(c).MinHeight(height).Padding(8, 16).Gap(12).FillWidth().
			Background(bg)
		if multiline {
			box.AlignItems(ui.Stretch)
		} else {
			box.AlignItems(ui.Center)
		}

		switch variant {
		case TextFieldFilled:
			box.Background(t.Scheme.SurfaceContainerHighest)
			// Square at the top, the active indicator's corners underneath.
			box.Radius(4, 4, 0, 0)
			box.BorderWidth(0, 0, 1, 0).BorderColor(border)
		default:
			box.Radius(theme.CornerSmall)
			box.Border(1, border)
		}

		box.Children(func() {
			inner := ui.Column(c).Grow(1).Gap(4).Children(func() {
				if label != "" {
					ui.Text(c, label).FontSize(11).TextColor(t.Scheme.Primary).SingleLine()
				}
				builder.input = fieldInput(c, value, multiline, height, t.Scheme.OnSurface)
			})
			if multiline {
				inner.Justify(ui.Start).Padding(6, 0)
			} else {
				inner.Justify(ui.Center)
			}
		})
	})

	return builder
}

// fieldInput is the editable part of a field: one line, or several that scroll.
func fieldInput(c *ui.Context, value *string, multiline bool, height float32, fg ui.Color) ui.Element {
	if multiline {
		// The label and the box's padding come out of the input's room.
		return ui.TextAreaBase(c, value).FillWidth().MinHeight(height - 48).TextColor(fg)
	}
	return ui.TextInputBase(c, value).FillWidth().TextColor(fg)
}

// Placeholder sets the input placeholder when empty.
func (b *TextFieldBuilder) Placeholder(p string) *TextFieldBuilder {
	b.input.Placeholder(p)
	return b
}

// Disabled marks the field as disabled.
func (b *TextFieldBuilder) Disabled(d bool) *TextFieldBuilder {
	b.input.Disabled(d)
	return b
}
