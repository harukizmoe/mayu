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
