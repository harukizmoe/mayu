package color

import (
	"math"

	"github.com/egoist/mygo/ui"
)

// DynamicScheme holds the six standard M3 tonal palettes derived from a seed color.
type DynamicScheme struct {
	Seed           HCT
	Primary        *TonalPalette
	Secondary      *TonalPalette
	Tertiary       *TonalPalette
	Neutral        *TonalPalette
	NeutralVariant *TonalPalette
	Error          *TonalPalette
}

// NewDynamicScheme generates standard Material Design 3 tonal palettes from a seed color.
func NewDynamicScheme(seed ui.Color) *DynamicScheme {
	return NewDynamicSchemeFromHCT(HCTFromColor(seed))
}

// NewDynamicSchemeFromHCT generates standard Material Design 3 tonal palettes from a seed HCT.
func NewDynamicSchemeFromHCT(seed HCT) *DynamicScheme {
	// M3 Dynamic Color rules:
	// Primary: Seed Hue, minimum Chroma 48
	primaryChroma := max(seed.Chroma, 48.0)
	primary := TonalPaletteOf(seed.Hue, primaryChroma)

	// Secondary: Seed Hue, Chroma 16
	secondary := TonalPaletteOf(seed.Hue, 16.0)

	// Tertiary: Seed Hue shifted by +60 degrees, Chroma 24
	tertiaryHue := math.Mod(seed.Hue+60.0, 360.0)
	tertiary := TonalPaletteOf(tertiaryHue, 24.0)

	// Neutral: Seed Hue, Chroma 4
	neutral := TonalPaletteOf(seed.Hue, 4.0)

	// Neutral Variant: Seed Hue, Chroma 8
	neutralVariant := TonalPaletteOf(seed.Hue, 8.0)

	// Error: Hue 25.0 (red), Chroma 84.0
	errorPalette := TonalPaletteOf(25.0, 84.0)

	return &DynamicScheme{
		Seed:           seed,
		Primary:        primary,
		Secondary:      secondary,
		Tertiary:       tertiary,
		Neutral:        neutral,
		NeutralVariant: neutralVariant,
		Error:          errorPalette,
	}
}
