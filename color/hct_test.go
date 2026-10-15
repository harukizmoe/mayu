package color

import (
	"math"
	"testing"

	"github.com/egoist/mygo/ui"
)

func TestHCT_SeedCalculations(t *testing.T) {
	seed := ui.Hex("#6750A4")
	hct := HCTFromColor(seed)

	// #6750A4 tone should be Tone 40 (~40.08)
	if math.Abs(hct.Tone-40.0) > 1.0 {
		t.Errorf("expected Tone ~40 for #6750A4, got %f", hct.Tone)
	}

	// Tone 0 must be pure black
	blackHCT := HCTOf(hct.Hue, hct.Chroma, 0)
	black := blackHCT.Color()
	if black.R != 0 || black.G != 0 || black.B != 0 {
		t.Errorf("expected Tone 0 to be black, got (%d, %d, %d)", black.R, black.G, black.B)
	}

	// Tone 100 must be pure white
	whiteHCT := HCTOf(hct.Hue, hct.Chroma, 100)
	white := whiteHCT.Color()
	if white.R != 255 || white.G != 255 || white.B != 255 {
		t.Errorf("expected Tone 100 to be white, got (%d, %d, %d)", white.R, white.G, white.B)
	}

	// Color() on hct should reconstruct #6750A4 within 1 unit per channel
	reconstructed := hct.Color()
	if math.Abs(float64(reconstructed.R)-float64(seed.R)) > 1 ||
		math.Abs(float64(reconstructed.G)-float64(seed.G)) > 1 ||
		math.Abs(float64(reconstructed.B)-float64(seed.B)) > 1 {
		t.Errorf("HCT roundtrip mismatch: got (%d, %d, %d), expected (%d, %d, %d)",
			reconstructed.R, reconstructed.G, reconstructed.B, seed.R, seed.G, seed.B)
	}
}

func TestTonalPalette(t *testing.T) {
	seed := ui.Hex("#6750A4")
	palette := TonalPaletteFromColor(seed)

	t10 := palette.Tone(10)
	t40 := palette.Tone(40)
	t80 := palette.Tone(80)
	t90 := palette.Tone(90)

	// Tone 10 should be very dark (low luminance)
	lum10 := 0.2126*float64(t10.R) + 0.7152*float64(t10.G) + 0.0722*float64(t10.B)
	lum40 := 0.2126*float64(t40.R) + 0.7152*float64(t40.G) + 0.0722*float64(t40.B)
	lum80 := 0.2126*float64(t80.R) + 0.7152*float64(t80.G) + 0.0722*float64(t80.B)
	lum90 := 0.2126*float64(t90.R) + 0.7152*float64(t90.G) + 0.0722*float64(t90.B)

	if !(lum10 < lum40 && lum40 < lum80 && lum80 < lum90) {
		t.Errorf("monotonic luminance violated: lum10=%f lum40=%f lum80=%f lum90=%f",
			lum10, lum40, lum80, lum90)
	}

	// T40 should be close to seed color
	if math.Abs(float64(t40.R)-float64(seed.R)) > 2 ||
		math.Abs(float64(t40.G)-float64(seed.G)) > 2 ||
		math.Abs(float64(t40.B)-float64(seed.B)) > 2 {
		t.Errorf("T40 mismatch with seed: got (%d, %d, %d), seed is (%d, %d, %d)",
			t40.R, t40.G, t40.B, seed.R, seed.G, seed.B)
	}
}

func TestDynamicScheme(t *testing.T) {
	seed := ui.Hex("#6750A4")
	scheme := NewDynamicScheme(seed)

	if scheme.Primary == nil || scheme.Secondary == nil || scheme.Tertiary == nil ||
		scheme.Neutral == nil || scheme.NeutralVariant == nil || scheme.Error == nil {
		t.Fatal("expected all 6 tonal palettes to be non-nil")
	}

	// Secondary palette should have lower chroma than primary
	if scheme.Secondary.Chroma >= scheme.Primary.Chroma {
		t.Errorf("expected secondary chroma (%f) < primary chroma (%f)",
			scheme.Secondary.Chroma, scheme.Primary.Chroma)
	}
}
