package color

import (
	"math"
	"testing"

	"github.com/egoist/mygo/ui"
)

func TestCAM16_SeedRoundTrip(t *testing.T) {
	// Canonical M3 Seed: #6750A4
	c := ui.Hex("#6750A4")
	cam := CAM16FromSRGB(c.R, c.G, c.B)

	// In M3, #6750A4 has:
	// Hue: ~298.98° (295° - 302°)
	// Chroma: ~47.86 (45 - 50)
	// J (Lightness): ~31.40 (30 - 33)
	if cam.Hue < 295.0 || cam.Hue > 303.0 {
		t.Errorf("expected Hue in [295, 303], got %f", cam.Hue)
	}
	if cam.Chroma < 45.0 || cam.Chroma > 50.0 {
		t.Errorf("expected Chroma in [45, 50], got %f", cam.Chroma)
	}
	if cam.J < 30.0 || cam.J > 33.0 {
		t.Errorf("expected J in [30, 33], got %f", cam.J)
	}

	// Test ToXYZ roundtrip
	x, y, z := cam.ToXYZ()
	r, g, b := xyzToSRGB(x, y, z)
	if math.Abs(float64(r)-float64(c.R)) > 1 ||
		math.Abs(float64(g)-float64(c.G)) > 1 ||
		math.Abs(float64(b)-float64(c.B)) > 1 {
		t.Errorf("CAM16 roundtrip mismatch: got (%d, %d, %d), expected (%d, %d, %d)", r, g, b, c.R, c.G, c.B)
	}
}
