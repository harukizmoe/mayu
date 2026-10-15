package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// CapsuleSliderOptions configures a wide capsule slider.
type CapsuleSliderOptions struct {
	Vertical   bool
	Width      float32
	Height     float32
	Icon       *ui.SVG
	Label      string
	FillColor  ui.Color
	TrackColor ui.Color
}

// CapsuleSliderBuilder provides a fluent builder for wide capsule sliders.
type CapsuleSliderBuilder struct {
	ui.Element
	c       *ui.Context
	value   *float64
	lo      float64
	hi      float64
	opts    CapsuleSliderOptions
	scheme  theme.ColorScheme
	density theme.Density
}

// VerticalCapsuleSlider creates an M3/Control Center vertical wide capsule slider (e.g. brightness or volume).
func VerticalCapsuleSlider(c *ui.Context, value *float64, lo, hi float64, icon *ui.SVG, t *theme.Theme) *CapsuleSliderBuilder {
	return NewCapsuleSlider(c, value, lo, hi, CapsuleSliderOptions{
		Vertical: true,
		Width:    72,
		Height:   160,
		Icon:     icon,
	}, t)
}

// HorizontalCapsuleSlider creates a horizontal wide capsule slider (e.g. wide pill tile with icon and label).
func HorizontalCapsuleSlider(c *ui.Context, value *float64, lo, hi float64, icon *ui.SVG, label string, t *theme.Theme) *CapsuleSliderBuilder {
	return NewCapsuleSlider(c, value, lo, hi, CapsuleSliderOptions{
		Vertical: false,
		Width:    220,
		Height:   60,
		Icon:     icon,
		Label:    label,
	}, t)
}

// NewCapsuleSlider creates a wide capsule slider with the specified options.
func NewCapsuleSlider(c *ui.Context, value *float64, lo, hi float64, opts CapsuleSliderOptions, t *theme.Theme) *CapsuleSliderBuilder {
	w, h := opts.Width, opts.Height
	if opts.Vertical {
		if w <= 0 {
			w = 72
		}
		if h <= 0 {
			h = 160
		}
	} else {
		if w <= 0 {
			w = 220
		}
		if h <= 0 {
			h = 60
		}
	}

	radius := w / 2
	if !opts.Vertical {
		radius = h / 2
	}

	s := ui.SliderBase(c, value, lo, hi).FocusRing(false).Clip().Size(w, h).MinWidth(w).MinHeight(h).Radius(radius)
	if opts.Vertical {
		s.Vertical()
	}

	builder := &CapsuleSliderBuilder{
		Element: s,
		c:       c,
		value:   value,
		lo:      lo,
		hi:      hi,
		opts:    opts,
		scheme:  t.Scheme,
		density: t.Density,
	}

	// The capsule takes the scheme's own roles, so it follows the seed and the
	// mode rather than a color of its own: a raised container for the empty
	// track, the primary for what is filled.
	trackBg := opts.TrackColor
	if trackBg == (ui.Color{}) {
		trackBg = t.Scheme.SurfaceContainerHighest
	}
	activeFill := opts.FillColor
	if activeFill == (ui.Color{}) {
		activeFill = t.Scheme.Primary
	}

	// Draw custom filled capsule
	s.Draw(func(p *ui.Painter, r ui.Rect) {
		rad := r.W / 2
		if !opts.Vertical {
			rad = r.H / 2
		}

		val := *value
		frac := float32(0)
		if hi > lo {
			frac = float32(max(0, min(1, (val-lo)/(hi-lo))))
		}

		p.Clip(r, rad, func() {
			// 1. Inactive track fill
			p.Fill(r, trackBg, rad)

			// 2. Dynamic active fill rect
			var fillRect ui.Rect
			if opts.Vertical {
				fillH := r.H * frac
				fillRect = ui.Rect{X: r.X, Y: r.Y + r.H - fillH, W: r.W, H: fillH}
			} else {
				fillW := r.W * frac
				fillRect = ui.Rect{X: r.X, Y: r.Y, W: fillW, H: r.H}
			}

			if fillRect.W > 0 && fillRect.H > 0 {
				p.Fill(fillRect, activeFill, 0)
			}
		})
	})

	// Add overlay Children on top of custom slider draw
	s.Children(func() {
		val := *value
		frac := float32(0)
		if hi > lo {
			frac = float32(max(0, min(1, (val-lo)/(hi-lo))))
		}

		// The icon sits on whichever of the two colors it overlaps, and takes
		// the role that reads on it.
		onFillFg := readableOn(activeFill, t.Scheme)
		onTrackFg := readableOn(trackBg, t.Scheme)

		if opts.Vertical {
			// Vertical layout: icon centered at bottom
			iconFg := onTrackFg
			if frac >= 0.22 {
				iconFg = onFillFg
			}

			ui.Column(c).Fill().Justify(ui.End).AlignItems(ui.Center).Padding(0, 0, 18, 0).Children(func() {
				if opts.Icon != nil {
					ui.Icon(c, opts.Icon).Size(28, 28).TextColor(iconFg)
				}
			})
		} else {
			// Horizontal layout: row with icon on left and label beside it
			contentFg := onTrackFg
			if frac >= 0.45 {
				contentFg = onFillFg
			}

			ui.Row(c).Fill().AlignItems(ui.Center).Padding(0, 20).Gap(14).Children(func() {
				if opts.Icon != nil {
					ui.Icon(c, opts.Icon).Size(24, 24).TextColor(contentFg)
				}
				if opts.Label != "" {
					ui.Text(c, opts.Label).FontSize(16).Bold().TextColor(contentFg).SingleLine()
				}
			})
		}
	})

	return builder
}

// readableOn returns the scheme role that reads on c: the on-color of the
// nearest role, so an icon on a capsule is themed rather than black or white by
// hand. The two roles a capsule is built from are checked first.
func readableOn(c ui.Color, s theme.ColorScheme) ui.Color {
	lum := func(v ui.Color) float64 {
		return 0.2126*float64(v.R) + 0.7152*float64(v.G) + 0.0722*float64(v.B)
	}
	switch c {
	case s.Primary:
		return s.OnPrimary
	case s.SurfaceContainerHighest:
		return s.OnSurface
	}
	if lum(c) > 128 {
		return s.OnSurface
	}
	return s.InverseOnSurface
}
