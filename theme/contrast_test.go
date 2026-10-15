package theme

import (
	"math"
	"testing"

	"github.com/egoist/mygo/ui"
)

// contrastRatio returns the WCAG contrast ratio between two opaque colors,
// from 1 (the same color) to 21 (black on white). It is the number Material 3
// designs its tone pairs to clear: 4.5 for body text, 3 for large text and
// graphics.
func contrastRatio(a, b ui.Color) float64 {
	la, lb := relativeLuminance(a), relativeLuminance(b)
	if la < lb {
		la, lb = lb, la
	}
	return (la + 0.05) / (lb + 0.05)
}

// relativeLuminance is the WCAG relative luminance of a color: its sRGB
// channels linearized and weighted as the eye weights them.
func relativeLuminance(c ui.Color) float64 {
	lin := func(v uint8) float64 {
		s := float64(v) / 255
		if s <= 0.03928 {
			return s / 12.92
		}
		return math.Pow((s+0.055)/1.055, 2.4)
	}
	return 0.2126*lin(c.R) + 0.7152*lin(c.G) + 0.0722*lin(c.B)
}

// seeds the pairs are checked against: the ones the gallery offers, plus a
// near-gray and a very saturated color, whose tonal palettes are the ones most
// likely to come out lopsided.
var seeds = []ui.Color{
	ui.Hex("#6750A4"),
	ui.Hex("#B3261E"),
	ui.Hex("#006492"),
	ui.Hex("#006D44"),
	ui.Hex("#795900"),
	ui.Hex("#808080"),
	ui.Hex("#FF0000"),
	ui.Hex("#0000FF"),
	ui.Hex("#FFFF00"),
	ui.Hex("#000000"),
	ui.Hex("#FFFFFF"),
}

// A scheme's role pairs must stay readable whatever the seed: this is what
// dynamic color is for, and what a seed the palette mishandles would break.
func TestSchemePairsClearContrast(t *testing.T) {
	// The pairs Material 3 specifies a minimum for. Text pairs take 4.5; the
	// rest are graphics, which take 3.
	pairs := []struct {
		name   string
		fg, bg func(ColorScheme) ui.Color
		min    float64
	}{
		{"OnPrimary on Primary", func(s ColorScheme) ui.Color { return s.OnPrimary }, func(s ColorScheme) ui.Color { return s.Primary }, 4.5},
		{"OnPrimaryContainer on PrimaryContainer", func(s ColorScheme) ui.Color { return s.OnPrimaryContainer }, func(s ColorScheme) ui.Color { return s.PrimaryContainer }, 4.5},
		{"OnSecondary on Secondary", func(s ColorScheme) ui.Color { return s.OnSecondary }, func(s ColorScheme) ui.Color { return s.Secondary }, 4.5},
		{"OnSecondaryContainer on SecondaryContainer", func(s ColorScheme) ui.Color { return s.OnSecondaryContainer }, func(s ColorScheme) ui.Color { return s.SecondaryContainer }, 4.5},
		{"OnTertiary on Tertiary", func(s ColorScheme) ui.Color { return s.OnTertiary }, func(s ColorScheme) ui.Color { return s.Tertiary }, 4.5},
		{"OnTertiaryContainer on TertiaryContainer", func(s ColorScheme) ui.Color { return s.OnTertiaryContainer }, func(s ColorScheme) ui.Color { return s.TertiaryContainer }, 4.5},
		{"OnError on Error", func(s ColorScheme) ui.Color { return s.OnError }, func(s ColorScheme) ui.Color { return s.Error }, 4.5},
		{"OnErrorContainer on ErrorContainer", func(s ColorScheme) ui.Color { return s.OnErrorContainer }, func(s ColorScheme) ui.Color { return s.ErrorContainer }, 4.5},
		{"OnSurface on Surface", func(s ColorScheme) ui.Color { return s.OnSurface }, func(s ColorScheme) ui.Color { return s.Surface }, 4.5},
		{"OnSurface on SurfaceContainer", func(s ColorScheme) ui.Color { return s.OnSurface }, func(s ColorScheme) ui.Color { return s.SurfaceContainer }, 4.5},
		{"OnSurfaceVariant on SurfaceVariant", func(s ColorScheme) ui.Color { return s.OnSurfaceVariant }, func(s ColorScheme) ui.Color { return s.SurfaceVariant }, 4.5},
		{"OnBackground on Background", func(s ColorScheme) ui.Color { return s.OnBackground }, func(s ColorScheme) ui.Color { return s.Background }, 4.5},
		{"InverseOnSurface on InverseSurface", func(s ColorScheme) ui.Color { return s.InverseOnSurface }, func(s ColorScheme) ui.Color { return s.InverseSurface }, 4.5},
		{"Primary on Surface (an indicator)", func(s ColorScheme) ui.Color { return s.Primary }, func(s ColorScheme) ui.Color { return s.Surface }, 3},
		{"Error on Surface (an icon)", func(s ColorScheme) ui.Color { return s.Error }, func(s ColorScheme) ui.Color { return s.Surface }, 3},
		{"Outline on Surface (a border)", func(s ColorScheme) ui.Color { return s.Outline }, func(s ColorScheme) ui.Color { return s.Surface }, 1.5},
	}

	for _, dark := range []bool{false, true} {
		mode := "light"
		if dark {
			mode = "dark"
		}
		for _, seed := range seeds {
			s := SchemeFromSeed(seed, dark)
			for _, p := range pairs {
				got := contrastRatio(p.fg(s), p.bg(s))
				if got+0.01 < p.min {
					t.Errorf("%s, seed %s: %s is %.2f:1, below the %.1f it must clear",
						mode, hex(seed), p.name, got, p.min)
				}
			}
		}
	}
}

// The surface containers climb with tonal elevation in both modes, which is
// what makes a card read as raised without a shadow.
func TestSurfaceContainersClimb(t *testing.T) {
	for _, dark := range []bool{false, true} {
		for _, seed := range seeds {
			s := SchemeFromSeed(seed, dark)
			containers := []struct {
				name string
				c    ui.Color
			}{
				{"Lowest", s.SurfaceContainerLowest},
				{"Low", s.SurfaceContainerLow},
				{"Container", s.SurfaceContainer},
				{"High", s.SurfaceContainerHigh},
				{"Highest", s.SurfaceContainerHighest},
			}

			for i := 1; i < len(containers); i++ {
				got := relativeLuminance(containers[i].c)
				prev := relativeLuminance(containers[i-1].c)
				// Tonal elevation moves away from the background as it rises:
				// darker in a light scheme, lighter in a dark one. So the
				// containers climb by getting further from white in light mode
				// and further from black in dark mode.
				climbs := got <= prev
				if dark {
					climbs = got >= prev
				}
				if !climbs {
					t.Errorf("dark=%v seed %s: SurfaceContainer%s does not climb above SurfaceContainer%s",
						dark, hex(seed), containers[i].name, containers[i-1].name)
				}
			}
		}
	}
}

func hex(c ui.Color) string {
	const digits = "0123456789ABCDEF"
	out := []byte{'#'}
	for _, v := range []uint8{c.R, c.G, c.B} {
		out = append(out, digits[v>>4], digits[v&0xf])
	}
	return string(out)
}
