package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// ButtonVariant specifies the M3 button emphasis style.
type ButtonVariant int

const (
	// VariantFilled is the high-emphasis primary button.
	VariantFilled ButtonVariant = iota
	// VariantElevated is the high-emphasis elevated surface button.
	VariantElevated
	// VariantFilledTonal is the medium-emphasis secondary container button.
	VariantFilledTonal
	// VariantOutlined is the medium-low emphasis outlined button.
	VariantOutlined
	// VariantText is the lowest-emphasis borderless button.
	VariantText
)

// ButtonBuilder provides a fluent builder interface for Material 3 buttons.
type ButtonBuilder struct {
	ui.Element
	c       *ui.Context
	variant ButtonVariant
	label   string
	density theme.Density
	scheme  theme.ColorScheme
	theme   *theme.Theme
}

// NewButton creates a Material 3 button with the given variant and theme.
func NewButton(c *ui.Context, label string, variant ButtonVariant, t *theme.Theme) *ButtonBuilder {
	b := ui.ButtonBase(c)

	builder := &ButtonBuilder{
		Element: b,
		c:       c,
		variant: variant,
		label:   label,
		density: t.Density,
		scheme:  t.Scheme,
		theme:   t,
	}

	height := t.Density.ControlHeight()
	padX := t.Density.ButtonPaddingX()
	gap := t.Density.SpacingUnit() * 2

	// Setup basic metrics: pill shape (radius = height / 2)
	b.Height(height).MinHeight(height).
		Radius(height/2).
		Gap(gap).
		Padding(0, padX).
		Center()

	if label != "" {
		b.Label(label)
		b.Children(func() {
			ui.Text(c, label).SingleLine()
		})
	}

	builder.style()
	return builder
}

// Variant restyles the button as another of the five Material 3 variants. It
// lets one button change emphasis — a choice that becomes filled once it is
// the chosen one, say — without building a second element beside it.
func (b *ButtonBuilder) Variant(v ButtonVariant) *ButtonBuilder {
	b.variant = v
	b.style()
	return b
}

// Icon adds a leading Material Symbol / SVG icon to the button.
func (b *ButtonBuilder) Icon(s *ui.SVG) *ButtonBuilder {
	if s == nil {
		return b
	}
	return b.Leading(func() { ui.Icon(b.c, s).Size(18, 18) })
}

// Leading puts build's content before the label, for a leading slot an icon
// alone cannot draw, as a color swatch:
//
//	mayu.FilledButton(c, "Coral Red").Leading(func() {
//		ui.Box(c).Size(14, 14).Radius(7).Background(coral)
//	})
//
// The row is reversed so the content built here comes first, and the
// horizontal padding becomes the narrower leading one, as Material 3 asks of
// a button with a leading slot.
func (b *ButtonBuilder) Leading(build func()) *ButtonBuilder {
	if build == nil {
		return b
	}
	b.Element.Reverse()
	b.Element.Padding(0, b.density.ButtonPaddingX(), 0, b.density.ButtonPaddingXLeadingIcon())
	b.Element.Children(build)
	return b
}

// TrailingIcon adds a trailing Material Symbol / SVG icon to the button.
func (b *ButtonBuilder) TrailingIcon(s *ui.SVG) *ButtonBuilder {
	if s == nil {
		return b
	}
	return b.Trailing(func() { ui.Icon(b.c, s).Size(18, 18) })
}

// Trailing puts build's content after the label, in a trailing slot of its own.
func (b *ButtonBuilder) Trailing(build func()) *ButtonBuilder {
	if build == nil {
		return b
	}
	b.Element.Padding(0, b.density.ButtonPaddingXLeadingIcon(), 0, b.density.ButtonPaddingX())
	b.Element.Children(build)
	return b
}

// OnClick registers a callback invoked when the button is clicked and not disabled.
func (b *ButtonBuilder) OnClick(fn func()) *ButtonBuilder {
	if fn != nil && b.Element.Clicked() && !b.Element.IsDisabled() {
		fn()
	}
	return b
}

// Disabled marks the button as disabled or enabled.
func (b *ButtonBuilder) Disabled(d bool) *ButtonBuilder {
	b.Element.Disabled(d)
	b.style()
	return b
}

func (b *ButtonBuilder) style() {
	disabled := b.Element.IsDisabled()
	hovered := b.Element.Hovered()
	focused := b.Element.Focused()
	pressed := b.Element.Pressed()

	colors := resolveButtonColors(b.scheme, b.variant, hovered, focused, pressed, disabled)

	b.Element.Background(colors.Container).TextColor(colors.Content)

	if colors.BorderWidth > 0 && colors.Border.A > 0 {
		b.Element.Border(colors.BorderWidth, colors.Border)
	} else {
		b.Element.Border(0, ui.Transparent)
	}

	if colors.Shadow > 0 {
		b.Element.Shadow(0, colors.Shadow, colors.Shadow*3, colors.Shadow, b.scheme.Shadow.Alpha(0.18))
	}
}

type buttonResolvedColors struct {
	Container   ui.Color
	Content     ui.Color
	Border      ui.Color
	BorderWidth float32
	Shadow      float32
}

func resolveButtonColors(scheme theme.ColorScheme, variant ButtonVariant, hovered, focused, pressed, disabled bool) buttonResolvedColors {
	if disabled {
		switch variant {
		case VariantOutlined:
			return buttonResolvedColors{
				Container:   ui.Transparent,
				Content:     scheme.OnSurface.Alpha(DisabledContentAlpha),
				Border:      scheme.OnSurface.Alpha(DisabledOutlineAlpha),
				BorderWidth: 1,
			}
		case VariantText:
			return buttonResolvedColors{
				Container: ui.Transparent,
				Content:   scheme.OnSurface.Alpha(DisabledContentAlpha),
			}
		default: // Filled, Elevated, FilledTonal
			return buttonResolvedColors{
				Container: scheme.OnSurface.Alpha(DisabledContainerAlpha),
				Content:   scheme.OnSurface.Alpha(DisabledContentAlpha),
			}
		}
	}

	switch variant {
	case VariantElevated:
		base := scheme.SurfaceContainerLow
		fg := scheme.Primary
		bg := ApplyStateLayer(base, scheme.Primary, hovered, focused, pressed)
		shadow := float32(1)
		if pressed {
			shadow = 1
		} else if hovered {
			shadow = 2
		}
		return buttonResolvedColors{
			Container: bg,
			Content:   fg,
			Shadow:    shadow,
		}

	case VariantFilledTonal:
		base := scheme.SecondaryContainer
		fg := scheme.OnSecondaryContainer
		bg := ApplyStateLayer(base, scheme.OnSecondaryContainer, hovered, focused, pressed)
		return buttonResolvedColors{
			Container: bg,
			Content:   fg,
		}

	case VariantOutlined:
		fg := scheme.Primary
		bg := ApplyStateLayer(ui.Transparent, scheme.Primary, hovered, focused, pressed)
		borderColor := scheme.Outline
		if focused {
			borderColor = scheme.Primary
		}
		return buttonResolvedColors{
			Container:   bg,
			Content:     fg,
			Border:      borderColor,
			BorderWidth: 1,
		}

	case VariantText:
		fg := scheme.Primary
		bg := ApplyStateLayer(ui.Transparent, scheme.Primary, hovered, focused, pressed)
		return buttonResolvedColors{
			Container: bg,
			Content:   fg,
		}

	default: // VariantFilled
		base := scheme.Primary
		fg := scheme.OnPrimary
		bg := ApplyStateLayer(base, scheme.OnPrimary, hovered, focused, pressed)
		return buttonResolvedColors{
			Container: bg,
			Content:   fg,
		}
	}
}
