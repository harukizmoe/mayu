package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// CheckboxBuilder wraps an M3 Checkbox element.
type CheckboxBuilder struct {
	ui.Element
	c       *ui.Context
	checked *bool
	label   string
	scheme  theme.ColorScheme
	density theme.Density
}

// NewCheckbox creates a Material 3 Checkbox control.
func NewCheckbox(c *ui.Context, checked *bool, label string, t *theme.Theme) *CheckboxBuilder {
	row := ui.CheckboxBase(c, checked).AlignItems(ui.Center).Gap(t.Density.SpacingUnit() * 2)

	builder := &CheckboxBuilder{
		Element: row,
		c:       c,
		checked: checked,
		label:   label,
		scheme:  t.Scheme,
		density: t.Density,
	}

	on := *checked
	disabled := row.IsDisabled()

	row.Children(func() {
		// Box container (18x18, radius 2)
		box := ui.Box(c).Size(18, 18).Radius(2).Center().Shrink(0)

		if disabled {
			if on {
				box.Background(builder.scheme.OnSurface.Alpha(DisabledContainerAlpha))
			} else {
				box.Border(2, builder.scheme.OnSurface.Alpha(DisabledOutlineAlpha))
			}
		} else if on {
			bg := builder.scheme.Primary
			if row.Pressed() {
				bg = builder.scheme.OnPrimary.Alpha(StatePressedAlpha).Over(bg)
			} else if row.Hovered() {
				bg = builder.scheme.OnPrimary.Alpha(StateHoverAlpha).Over(bg)
			}
			box.Background(bg)
		} else {
			border := builder.scheme.Outline
			if row.Focused() {
				border = builder.scheme.Primary
			}
			box.Border(2, border)
			if row.Hovered() {
				box.Background(builder.scheme.OnSurface.Alpha(StateHoverAlpha))
			}
		}

		// Checkmark inside box
		if on {
			fg := builder.scheme.OnPrimary
			if disabled {
				fg = builder.scheme.OnSurface.Alpha(DisabledContentAlpha)
			}
			box.Children(func() {
				ui.Icon(c, symbols.Check).Size(14, 14).TextColor(fg)
			})
		}

		// Accompanying label
		if label != "" {
			labelColor := builder.scheme.OnSurface
			if disabled {
				labelColor = builder.scheme.OnSurface.Alpha(DisabledContentAlpha)
			}
			ui.Text(c, label).TextColor(labelColor).SingleLine()
		}
	})

	return builder
}

// SwitchBuilder wraps an M3 Switch element.
type SwitchBuilder struct {
	ui.Element
	c       *ui.Context
	on      *bool
	label   string
	scheme  theme.ColorScheme
	density theme.Density
}

// NewSwitch creates a Material 3 Switch control.
func NewSwitch(c *ui.Context, on *bool, label string, t *theme.Theme) *SwitchBuilder {
	row := ui.Row(c).AlignItems(ui.Center).Gap(t.Density.SpacingUnit() * 2).Focusable()
	if row.Clicked() && !row.IsDisabled() {
		*on = !*on
	}
	builder := &SwitchBuilder{
		Element: row,
		c:       c,
		on:      on,
		label:   label,
		scheme:  t.Scheme,
		density: t.Density,
	}

	isOn := *on
	disabled := row.IsDisabled()

	row.Children(func() {
		// Switch track (52x32, radius 16)
		sw := ui.SwitchBase(c, on).Size(52, 32).Radius(16).Shrink(0).AlignItems(ui.Center)

		var trackBg ui.Color
		var trackBorder ui.Color
		var trackBorderWidth float32
		var thumbBg ui.Color

		if disabled {
			if isOn {
				trackBg = builder.scheme.OnSurface.Alpha(DisabledContainerAlpha)
				thumbBg = builder.scheme.Surface
			} else {
				trackBg = builder.scheme.SurfaceContainerHighest
				trackBorder = builder.scheme.OnSurface.Alpha(DisabledOutlineAlpha)
				trackBorderWidth = 2
				thumbBg = builder.scheme.OnSurface.Alpha(DisabledContentAlpha)
			}
		} else if isOn {
			trackBg = builder.scheme.Primary
			if sw.Pressed() {
				trackBg = builder.scheme.OnPrimary.Alpha(StatePressedAlpha).Over(trackBg)
			} else if sw.Hovered() {
				trackBg = builder.scheme.OnPrimary.Alpha(StateHoverAlpha).Over(trackBg)
			}
			thumbBg = builder.scheme.OnPrimary
		} else {
			trackBg = builder.scheme.SurfaceContainerHighest
			trackBorder = builder.scheme.Outline
			trackBorderWidth = 2
			thumbBg = builder.scheme.Outline
			if sw.Hovered() {
				thumbBg = builder.scheme.OnSurfaceVariant
			}
		}

		sw.Background(trackBg)
		if trackBorderWidth > 0 {
			sw.Border(trackBorderWidth, trackBorder)
		}

		// Thumb placement (24x24, radius 12)
		sw.Children(func() {
			if isOn {
				ui.Row(c).FillWidth().Justify(ui.End).Padding(0, 4).Children(func() {
					ui.Box(c).Size(24, 24).Radius(12).Background(thumbBg).Center()
				})
			} else {
				ui.Row(c).FillWidth().Justify(ui.Start).Padding(0, 4).Children(func() {
					ui.Box(c).Size(16, 16).Radius(8).Background(thumbBg).Center()
				})
			}
		})

		if label != "" {
			labelColor := builder.scheme.OnSurface
			if disabled {
				labelColor = builder.scheme.OnSurface.Alpha(DisabledContentAlpha)
			}
			ui.Text(c, label).TextColor(labelColor).SingleLine()
		}
	})

	return builder
}

// RadioBuilder wraps an M3 Radio element.
type RadioBuilder[T comparable] struct {
	ui.Element
	c        *ui.Context
	selected *T
	value    T
	label    string
	scheme   theme.ColorScheme
	density  theme.Density
}

// NewRadio creates a Material 3 Radio control.
func NewRadio[T comparable](c *ui.Context, selected *T, value T, label string, t *theme.Theme) *RadioBuilder[T] {
	row := ui.RadioBase(c, selected, value).AlignItems(ui.Center).Gap(t.Density.SpacingUnit() * 2)

	builder := &RadioBuilder[T]{
		Element:  row,
		c:        c,
		selected: selected,
		value:    value,
		label:    label,
		scheme:   t.Scheme,
		density:  t.Density,
	}

	on := *selected == value
	disabled := row.IsDisabled()

	row.Children(func() {
		// Outer circle: 20x20, radius 10
		outer := ui.Box(c).Size(20, 20).Radius(10).Center().Shrink(0)

		if disabled {
			outer.Border(2, builder.scheme.OnSurface.Alpha(DisabledOutlineAlpha))
			if on {
				outer.Children(func() {
					ui.Box(c).Size(10, 10).Radius(5).Background(builder.scheme.OnSurface.Alpha(DisabledContentAlpha))
				})
			}
		} else if on {
			ringColor := builder.scheme.Primary
			if row.Pressed() {
				ringColor = builder.scheme.Primary.Alpha(0.8)
			}
			outer.Border(2, ringColor)
			outer.Children(func() {
				ui.Box(c).Size(10, 10).Radius(5).Background(ringColor)
			})
		} else {
			borderColor := builder.scheme.Outline
			if row.Hovered() {
				borderColor = builder.scheme.OnSurface
			}
			outer.Border(2, borderColor)
		}

		if label != "" {
			labelColor := builder.scheme.OnSurface
			if disabled {
				labelColor = builder.scheme.OnSurface.Alpha(DisabledContentAlpha)
			}
			ui.Text(c, label).TextColor(labelColor).SingleLine()
		}
	})

	return builder
}
