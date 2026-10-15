package theme

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/color"
)

func TestColorScheme_ContrastInvariants(t *testing.T) {
	seed := ui.Hex("#6750A4")
	dyn := color.NewDynamicScheme(seed)

	light := LightScheme(dyn)
	dark := DarkScheme(dyn)

	// Light scheme checks
	lightPrimaryTone := color.HCTFromColor(light.Primary).Tone
	lightOnPrimaryTone := color.HCTFromColor(light.OnPrimary).Tone
	if lightOnPrimaryTone-lightPrimaryTone < 40 {
		t.Errorf("insufficient contrast between Primary and OnPrimary in light: %f vs %f",
			lightPrimaryTone, lightOnPrimaryTone)
	}

	// Dark scheme checks
	darkPrimaryTone := color.HCTFromColor(dark.Primary).Tone
	darkOnPrimaryTone := color.HCTFromColor(dark.OnPrimary).Tone
	if darkPrimaryTone-darkOnPrimaryTone < 40 {
		t.Errorf("insufficient contrast between Primary and OnPrimary in dark: %f vs %f",
			darkPrimaryTone, darkOnPrimaryTone)
	}

	// Dark mode surface containers tonal elevation graduation:
	// Lowest <= Low <= Default <= High <= Highest
	tLowest := color.HCTFromColor(dark.SurfaceContainerLowest).Tone
	tLow := color.HCTFromColor(dark.SurfaceContainerLow).Tone
	tDefault := color.HCTFromColor(dark.SurfaceContainer).Tone
	tHigh := color.HCTFromColor(dark.SurfaceContainerHigh).Tone
	tHighest := color.HCTFromColor(dark.SurfaceContainerHighest).Tone

	if !(tLowest <= tLow && tLow <= tDefault && tDefault <= tHigh && tHigh <= tHighest) {
		t.Errorf("dark surface container graduation violated: %f <= %f <= %f <= %f <= %f",
			tLowest, tLow, tDefault, tHigh, tHighest)
	}
}

func TestDensity(t *testing.T) {
	if DensityCompact.ControlHeight() >= DensityStandard.ControlHeight() {
		t.Errorf("expected compact height < standard height")
	}
	if DensityStandard.ControlHeight() >= DensityComfortable.ControlHeight() {
		t.Errorf("expected standard height < comfortable height")
	}
}
