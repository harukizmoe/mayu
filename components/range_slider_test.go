package components

import (
	"math"
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// The control spans x 16..304 in a 320 wide window under a Column's padding of
// 16, 44 tall from y 16, with its track inset by half a thumb at each end: 26
// to 294 for a thumb of 20. The pointer is measured against that.
func rangeTest(t *testing.T, low, high *float64) *ui.Tester {
	t.Helper()
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)
	return ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Padding(16).Children(func() {
			RangeSlider(c, low, high, 0, 100, RangeSliderOptions{Label: "Range"}, th)
		})
	}, 320, 80)
}

func TestRangeSliderDragsEitherThumb(t *testing.T) {
	low, high := 25.0, 75.0
	tt := rangeTest(t, &low, &high)

	// Pressing the track left of centre brings the lower thumb there.
	tt.Press(60, 38)
	tt.Release(60, 38)
	if low >= 25 {
		t.Errorf("the press did not move the lower thumb: low=%v", low)
	}
	if high != 75 {
		t.Errorf("the press moved the upper thumb too: high=%v", high)
	}

	// Dragging it does not stop at the press.
	tt.Press(60, 38)
	tt.Move(140, 38)
	tt.Release(140, 38)
	if low <= 20 || low >= high {
		t.Errorf("the drag left the lower thumb at %v, with the upper at %v", low, high)
	}

	// Pressing right of centre takes the upper thumb instead.
	wasLow := low
	tt.Press(260, 38)
	tt.Release(260, 38)
	if high <= 75 {
		t.Errorf("the press did not move the upper thumb: high=%v", high)
	}
	if low != wasLow {
		t.Errorf("the press moved the lower thumb: low=%v, want %v", low, wasLow)
	}
}

// Neither thumb passes the other: dragging the lower one past the upper stops
// it there rather than swapping them.
func TestRangeSliderThumbsDoNotCross(t *testing.T) {
	low, high := 20.0, 40.0
	tt := rangeTest(t, &low, &high)

	tt.Press(60, 38)
	tt.Move(300, 38)
	tt.Release(300, 38)

	if low > high {
		t.Errorf("the lower thumb passed the upper: low=%v, high=%v", low, high)
	}
	if math.Abs(low-high) > 0.001 {
		t.Errorf("the lower thumb did not stop at the upper: low=%v, high=%v", low, high)
	}
}

// The keys move the lower end, and Shift the upper.
func TestRangeSliderKeys(t *testing.T) {
	low, high := 20.0, 80.0
	tt := rangeTest(t, &low, &high)

	if err := tt.Click("Range"); err != nil {
		t.Fatalf("failed focusing the range slider: %v", err)
	}
	tt.Key(0, ui.KeyRight)
	if low <= 20 {
		t.Errorf("Right did not move the lower end: low=%v", low)
	}
	wasLow := low
	tt.Key(ui.Shift, ui.KeyLeft)
	if high >= 80 {
		t.Errorf("Shift+Left did not move the upper end: high=%v", high)
	}
	if low != wasLow {
		t.Errorf("Shift+Left moved the lower end too: low=%v, want %v", low, wasLow)
	}
}
