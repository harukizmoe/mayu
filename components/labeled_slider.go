package components

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// LabeledSliderOptions configures a card-style labeled slider with a value badge.
type LabeledSliderOptions struct {
	// Label is the header title, as "亮度 (Brightness)".
	Label string
	// Icon shows before the label, nil for none.
	Icon *ui.SVG
	// FormatValue formats the badge text; nil shows "90 %".
	FormatValue func(v float64) string
	// IconColor overrides the header icon color, zero for the primary color.
	IconColor ui.Color
	// ActiveColor overrides the filled track and the thumb, zero for the
	// primary color.
	ActiveColor ui.Color
	// InactiveColor overrides the empty track, zero for a tint of the active
	// color over the card.
	InactiveColor ui.Color
}

// LabeledSlider creates an M3 card-style slider: an elevated container holding a
// header of an icon, a label and a value badge, above a full-width track with a
// round thumb ringed by the card's own color.
func LabeledSlider(c *ui.Context, value *float64, lo, hi float64, opts LabeledSliderOptions, t *theme.Theme) ui.Element {
	scheme := t.Scheme

	const (
		trackH = 8  // the track's height
		thumbR = 13 // the thumb's outer radius
		coreR  = 9  // the thumb's filled center
		badgeH = 28 // the value badge's height
		cardR  = 28 // the card's corner radius
		padR   = 14 // the value badge's horizontal padding
		vGap   = 12 // between the header and the track
	)

	cardBg := scheme.SurfaceContainerLow

	active := opts.ActiveColor
	if active == (ui.Color{}) {
		active = scheme.Primary
	}
	inactive := opts.InactiveColor
	if inactive == (ui.Color{}) {
		inactive = active.Alpha(0.16).Over(cardBg)
	}
	iconColor := opts.IconColor
	if iconColor == (ui.Color{}) {
		iconColor = scheme.Primary
	}
	format := opts.FormatValue
	if format == nil {
		format = func(v float64) string { return fmt.Sprintf("%.0f %%", v) }
	}

	card := ui.Column(c).FillWidth().Gap(vGap).Radius(cardR).Padding(16).
		Background(cardBg).
		Shadow(0, 2, 8, 2, scheme.Shadow.Alpha(0.12))

	card.Children(func() {
		// Header: icon, label, value badge.
		ui.Row(c).FillWidth().AlignItems(ui.Center).Gap(8).Children(func() {
			if opts.Icon != nil {
				ui.Icon(c, opts.Icon).Size(20, 20).TextColor(iconColor)
			}
			if opts.Label != "" {
				ui.Text(c, opts.Label).FontSize(theme.TitleSmall.Size).
					TextColor(scheme.OnSurface).Grow(1).SingleLine()
			} else {
				ui.Spacer(c)
			}

			badge := ui.Row(c).Height(badgeH).Radius(badgeH/2).Padding(0, padR).
				Center().Shrink(0).Border(1, scheme.OutlineVariant)
			badge.Children(func() {
				ui.Text(c, format(*value)).FontSize(theme.LabelLarge.Size).
					Bold().TextColor(scheme.OnSurface).SingleLine()
			})
		})

		// Track and thumb, dragged through SliderBase.
		s := ui.SliderBase(c, value, lo, hi).Height(thumbR * 2).FillWidth().
			PaddingX(thumbR).FocusRing(false)
		s.Draw(func(p *ui.Painter, r ui.Rect) {
			if r.W <= 0 || r.H <= 0 {
				return
			}
			frac := float32(0)
			if hi > lo {
				frac = float32(max(0, min(1, (*value-lo)/(hi-lo))))
			}

			cy := r.Y + r.H/2
			trackY := cy - trackH/2
			p.Fill(ui.Rect{X: r.X, Y: trackY, W: r.W, H: trackH}, inactive, trackH/2)

			x := r.X + r.W*frac
			if x > r.X {
				p.Fill(ui.Rect{X: r.X, Y: trackY, W: x - r.X, H: trackH}, active, trackH/2)
			}

			// The thumb: a soft shadow, a ring in the card's color, the core.
			outer := ui.Rect{X: x - thumbR, Y: cy - thumbR, W: thumbR * 2, H: thumbR * 2}
			p.Shadow(outer, thumbR, 0, 1, 3, 1, scheme.Shadow.Alpha(0.22))
			p.Fill(outer, cardBg, thumbR)
			p.Fill(ui.Rect{X: x - coreR, Y: cy - coreR, W: coreR * 2, H: coreR * 2}, active, coreR)
		})
	})

	return card
}
