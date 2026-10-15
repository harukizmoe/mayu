package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

func TestCapsuleSlider_RenderAndDrag(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	brightnessVal := 75.0
	lightSwitchVal := 30.0

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Gap(16).Children(func() {
			VerticalCapsuleSlider(c, &brightnessVal, 0, 100, symbols.Lightbulb, th)
			HorizontalCapsuleSlider(c, &lightSwitchVal, 0, 100, symbols.LightMode, "开灯", th)
		})
	}, 400, 400)

	if !tt.HasText("开灯") {
		t.Fatal("horizontal capsule slider label '开灯' not found")
	}

	if brightnessVal != 75.0 {
		t.Errorf("expected initial brightness == 75, got %f", brightnessVal)
	}
	if lightSwitchVal != 30.0 {
		t.Errorf("expected initial lightSwitchVal == 30, got %f", lightSwitchVal)
	}
}

// The capsule takes the scheme's own roles, so it follows the seed and the
// mode: its empty track is a raised container and its filled part the primary.
// Sampled at the two ends of the track, away from the icon at the foot, in both
// modes.
func TestCapsuleSliderUsesSchemeRoles(t *testing.T) {
	const (
		width  = 72
		height = 160
		value  = 60.0
	)
	for _, dark := range []bool{false, true} {
		scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), dark)
		th := theme.NewTheme(scheme, theme.DensityStandard)

		v := value
		tt := ui.NewTester(func(c *ui.Context) {
			ui.Column(c).Padding(16).Children(func() {
				VerticalCapsuleSlider(c, &v, 0, 100, symbols.Lightbulb, th)
			})
		}, width+32, height+32)

		img := tt.Image()
		at := func(x, y int) ui.Color {
			i := img.PixOffset(x, y)
			return ui.RGB(img.Pix[i], img.Pix[i+1], img.Pix[i+2])
		}

		// The capsule spans 16..16+height vertically, filled from the foot up
		// 60% of the way, so 40 DIPs down from the top is track and 20 DIPs up
		// from the foot is fill. Both sampled left of the icon.
		x := 16 + 12
		if got := at(x, 16+40); got != scheme.SurfaceContainerHighest {
			t.Errorf("dark=%v: the empty track is #%02X%02X%02X, want SurfaceContainerHighest #%02X%02X%02X",
				dark, got.R, got.G, got.B,
				scheme.SurfaceContainerHighest.R, scheme.SurfaceContainerHighest.G, scheme.SurfaceContainerHighest.B)
		}
		if got := at(x, 16+height-20); got != scheme.Primary {
			t.Errorf("dark=%v: the filled part is #%02X%02X%02X, want the primary #%02X%02X%02X",
				dark, got.R, got.G, got.B, scheme.Primary.R, scheme.Primary.G, scheme.Primary.B)
		}
	}
}
