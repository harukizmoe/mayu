package components

import (
	"math"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// RangeSliderOptions configures a RangeSlider.
type RangeSliderOptions struct {
	// Step rounds the values the keys take to multiples of it, 0 for a
	// hundredth of the range.
	Step float64
	// Label names the control for assistive technology. A range is two values,
	// which mygo's slider role describes as one, so the name is where the
	// control says what it is: "Price range".
	Label string
}

// RangeSlider creates a Material 3 range slider: one track with a thumb at each
// end of a span, the span between them filled, and neither thumb passing the
// other. Drag either thumb, or press the track to bring the nearer one there.
//
// The lower end moves with Left and Right while the control has the keyboard
// focus, and the upper end with Shift and the same keys.
func RangeSlider(c *ui.Context, low, high *float64, lo, hi float64, opts RangeSliderOptions, t *theme.Theme) ui.Element {
	const (
		diameter  = 20 // a thumb's diameter
		thickness = 4  // the track's thickness
	)
	if hi <= lo {
		hi = lo + 1
	}
	clamp := func(v, a, b float64) float64 { return min(max(v, a), b) }
	*low, *high = clamp(*low, lo, hi), clamp(*high, lo, hi)
	if *low > *high {
		*low = *high
	}

	step := opts.Step
	if step <= 0 {
		step = (hi - lo) / 100
	}

	box := ui.Box(c).Height(44).FillWidth().Focusable()
	if opts.Label != "" {
		box.Label(opts.Label)
	}
	box.Role(ui.RoleSlider)

	// The drag and the width the pointer is measured against, which the paint
	// of the frame before leaves here: input arrives before the next build.
	st := ui.Local(box, rangeKey{}, func() rangeState { return rangeState{} })
	if st == nil {
		return box
	}

	// fraction is where v sits between lo and hi, from 0 to 1.
	fraction := func(v float64) float32 { return float32(clamp((v-lo)/(hi-lo), 0, 1)) }
	// at is the fraction the pointer is at, inset by half a thumb at each end
	// so a thumb at either extreme stays inside the control.
	at := func(x float32) float64 {
		span := st.width - diameter
		if span <= 0 {
			return 0
		}
		return clamp(float64((x-diameter/2)/span), 0, 1)
	}
	// take moves the held thumb to f, without letting it pass the other.
	take := func(f float64) {
		v := lo + f*(hi-lo)
		switch st.held {
		case 1:
			*low = clamp(v, lo, *high)
		case 2:
			*high = clamp(v, *low, hi)
		}
	}

	box.HandleInput(func(ev ui.InputEvent) bool {
		switch ev.Kind {
		case ui.InputPointerDown:
			if ev.Button != 0 {
				return false
			}
			f := at(ev.X)
			// The nearer thumb takes the press: within half a thumb of either
			// the pointer grabs it where it is, and elsewhere it moves there.
			st.held = 1
			if absf(float64(fraction(*high))-f) < absf(float64(fraction(*low))-f) {
				st.held = 2
			}
			take(f)
			return true

		case ui.InputPointerUp:
			if st.held == 0 {
				return false
			}
			st.held = 0
			return true

		case ui.InputPointerMove:
			if st.held == 0 || ev.Button != 0 {
				return false
			}
			take(at(ev.X))
			return true

		case ui.InputKeyDown:
			up := ev.Mods&ui.Shift != 0
			by := func(d float64) {
				if up {
					*high = clamp(round(*high+d, step, lo), *low, hi)
					return
				}
				*low = clamp(round(*low+d, step, lo), lo, *high)
			}
			switch ev.Key {
			case ui.KeyLeft:
				by(-step)
			case ui.KeyRight:
				by(step)
			case ui.KeyHome:
				*low = lo
			case ui.KeyEnd:
				*high = hi
			default:
				return false
			}
			return true
		}
		return false
	})

	box.Draw(func(p *ui.Painter, r ui.Rect) {
		st.width = r.W
		if r.W <= diameter {
			return
		}
		inset, span := r.X+diameter/2, r.W-diameter
		cy := r.Y + r.H/2
		trackY := cy - thickness/2

		p.Fill(ui.Rect{X: inset, Y: trackY, W: span, H: thickness},
			t.Scheme.SecondaryContainer, thickness/2)

		x0 := inset + span*fraction(*low)
		x1 := inset + span*fraction(*high)
		if x1 > x0 {
			p.Fill(ui.Rect{X: x0, Y: trackY, W: x1 - x0, H: thickness},
				t.Scheme.Primary, thickness/2)
		}

		for i, x := range []float32{x0, x1} {
			// The held thumb is drawn a little larger, so it is clear which one
			// the pointer has.
			d := float32(diameter)
			if st.held == i+1 {
				d += 4
			}
			p.Fill(ui.Rect{X: x - d/2, Y: cy - d/2, W: d, H: d}, t.Scheme.Primary, d/2)
		}
	})

	return box
}

// rangeState is what a range slider keeps from frame to frame.
type rangeState struct {
	// held is the thumb the pointer has: 0 none, 1 the lower, 2 the upper.
	held int
	// width is the control's width, which the pointer is measured against.
	width float32
}

type rangeKey struct{}

// round takes v to the nearest multiple of step from base, as a step slider
// does, so the keys land on the same values the pointer does.
func round(v, step, base float64) float64 {
	if step <= 0 {
		return v
	}
	return base + float64(int((v-base)/step+math.Copysign(0.5, v-base)))*step
}

func absf(v float64) float64 { return math.Abs(v) }
