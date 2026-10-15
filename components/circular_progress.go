package components

import (
	"fmt"
	"math"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// CircularProgressOptions configures a CircularProgress.
type CircularProgressOptions struct {
	// Size is the ring's diameter in DIPs, 48 when zero.
	Size float32
	// Width is the stroke's width in DIPs, 4 when zero.
	Width float32
	// ShowValue draws the percentage in the middle of a determinate ring.
	ShowValue bool
	// TrackColor and IndicatorColor override the ring's colors.
	TrackColor     ui.Color
	IndicatorColor ui.Color
}

// CircularProgress draws a Material 3 circular progress ring. A progress
// between 0 and 1 is determinate; a negative progress is indeterminate and
// spins until the view reports one.
func CircularProgress(c *ui.Context, progress float32, opts CircularProgressOptions, t *theme.Theme) ui.Element {
	size := opts.Size
	if size <= 0 {
		size = 48
	}
	width := opts.Width
	if width <= 0 {
		width = 4
	}
	track := opts.TrackColor
	if track == (ui.Color{}) {
		track = t.Scheme.SecondaryContainer
	}
	indicator := opts.IndicatorColor
	if indicator == (ui.Color{}) {
		indicator = t.Scheme.Primary
	}

	spins := progress < 0
	fraction := min(max(progress, 0), 1)

	box := ui.Box(c).Size(size, size).Shrink(0)
	box.Draw(func(p *ui.Painter, r ui.Rect) {
		cx, cy := r.X+r.W/2, r.Y+r.H/2

		// Stroke paints inside the rectangle it is given, so the track's
		// rectangle is the box itself: its border runs from the edge inwards
		// one width. The arc is stroked on the centerline, half a width in.
		p.Stroke(r, track, min(r.W, r.H)/2, width)
		radius := min(r.W, r.H)/2 - width/2
		if radius <= 0 {
			return
		}

		// A determinate ring sweeps from twelve o'clock; an indeterminate one
		// carries an arc of it around, asking for the next frame as it goes.
		from, to := float32(0), fraction
		if spins {
			const (
				period = 1.4 // seconds for one turn
				sweep  = 0.28
			)
			from = float32(math.Mod(float64(p.Now().UnixNano())/1e9, period) / period)
			to = from + sweep
			p.AnimationFrame()
		}
		if to <= from {
			return
		}
		p.StrokePath(arcPath(cx, cy, radius, from, to), width, indicator)
	})

	if opts.ShowValue && !spins {
		box.Center().Children(func() {
			ui.Text(c, fmt.Sprintf("%.0f%%", fraction*100)).
				FontSize(size / 4).Bold().TextColor(indicator)
		})
	}

	return box
}

// arcPath returns the arc of the circle around (cx, cy) from the fraction from
// of it to the fraction to, clockwise from twelve o'clock, in cubic curves of
// at most a quarter turn each.
func arcPath(cx, cy, r, from, to float32) *ui.Path {
	path := &ui.Path{}
	if r <= 0 || to <= from {
		return path
	}
	if to-from >= 1 {
		return path.Circle(cx, cy, r)
	}

	a0 := 2 * math.Pi * float64(from)
	a1 := 2 * math.Pi * float64(to)
	n := int(math.Ceil((a1 - a0) / (math.Pi / 2)))
	step := (a1 - a0) / float64(n)
	k := float32(4.0 / 3 * math.Tan(step/4))

	x, y := polar(cx, cy, r, a0)
	path.MoveTo(x, y)
	for i := range n {
		b0, b1 := a0+float64(i)*step, a0+float64(i+1)*step
		x0, y0 := polar(cx, cy, r, b0)
		x1, y1 := polar(cx, cy, r, b1)
		// The tangent at angle a, going clockwise, is (cos a, sin a).
		t0x, t0y := float32(math.Cos(b0)), float32(math.Sin(b0))
		t1x, t1y := float32(math.Cos(b1)), float32(math.Sin(b1))
		path.CubeTo(x0+k*r*t0x, y0+k*r*t0y, x1-k*r*t1x, y1-k*r*t1y, x1, y1)
	}
	return path
}

// polar returns the point at angle a radians on the circle, with angle 0 at
// twelve o'clock and angles growing clockwise, as the screen's y axis does.
func polar(cx, cy, r float32, a float64) (float32, float32) {
	return cx + r*float32(math.Sin(a)), cy - r*float32(math.Cos(a))
}
