package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

func TestLabeledSlider(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	brightness := 90.0

	// The card fills a 400 DIP window with no outer padding, so its parts sit
	// at known places: 16 DIP of padding, a 28 DIP header, a 12 DIP gap, and a
	// 26 DIP track. The track's content box is inset by the thumb's radius.
	tt := ui.NewTester(func(c *ui.Context) {
		LabeledSlider(c, &brightness, 0, 100, LabeledSliderOptions{
			Label: "亮度 (Brightness)",
			Icon:  symbols.Lightbulb,
		}, th)
	}, 400, 200)

	if !tt.HasText("亮度 (Brightness)") {
		t.Fatal("labeled slider header not found")
	}
	if !tt.HasText("90 %") {
		t.Fatalf("value badge not found, texts: %q", tt.Texts())
	}

	// Pressing the middle of the track sets the middle value.
	tt.ClickAt(200, 69)
	if brightness != 50 {
		t.Errorf("pressing the middle of the track: got %v, want 50", brightness)
	}

	tt.Frame()
	if !tt.HasText("50 %") {
		t.Errorf("the badge does not follow the value: %q", tt.Texts())
	}
}
