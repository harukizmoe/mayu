package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// SliderBuilder wraps a Material 3 Slider.
type SliderBuilder struct {
	ui.Element
	c       *ui.Context
	value   *float64
	lo      float64
	hi      float64
	scheme  theme.ColorScheme
	density theme.Density
}

// NewSlider constructs a Material 3 continuous Slider.
func NewSlider(c *ui.Context, value *float64, lo, hi float64, t *theme.Theme) *SliderBuilder {
	const knobSize float32 = 20
	const trackHeight float32 = 4

	s := ui.SliderBase(c, value, lo, hi).
		Height(44).
		MinWidth(120).
		PaddingX(knobSize / 2).
		FocusRing(false)

	builder := &SliderBuilder{
		Element: s,
		c:       c,
		value:   value,
		lo:      lo,
		hi:      hi,
		scheme:  t.Scheme,
		density: t.Density,
	}

	frac := float32(0)
	if hi > lo {
		frac = float32(max(0, min(1, (*value-lo)/(hi-lo))))
	}

	s.Draw(func(p *ui.Painter, r ui.Rect) {
		cy := r.Y + r.H/2
		trackY := cy - trackHeight/2

		// 1. Inactive track (SecondaryContainer)
		p.Fill(ui.Rect{X: r.X, Y: trackY, W: r.W, H: trackHeight}, t.Scheme.SecondaryContainer, trackHeight/2)

		// Knob x position
		kx := r.X + r.W*frac

		// 2. Active track (Primary)
		if kx > r.X {
			p.Fill(ui.Rect{X: r.X, Y: trackY, W: kx - r.X, H: trackHeight}, t.Scheme.Primary, trackHeight/2)
		}

		// 3. Knob thumb (20x20 round circle in Primary)
		knobRect := ui.Rect{X: kx - knobSize/2, Y: cy - knobSize/2, W: knobSize, H: knobSize}
		p.Fill(knobRect, t.Scheme.Primary, knobSize/2)
	})

	return builder
}
