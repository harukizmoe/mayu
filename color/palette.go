package color

import (
	"math"
	"sync"

	"github.com/egoist/mygo/ui"
)

// TonalPalette provides a range of tones (luminance levels 0 to 100) for a fixed Hue and Chroma.
type TonalPalette struct {
	Hue    float64
	Chroma float64

	mu    sync.RWMutex
	cache map[int]ui.Color
}

// TonalPaletteOf constructs a TonalPalette with the given Hue and Chroma.
func TonalPaletteOf(hue, chroma float64) *TonalPalette {
	return &TonalPalette{
		Hue:    hue,
		Chroma: chroma,
		cache:  make(map[int]ui.Color),
	}
}

// TonalPaletteFromColor constructs a TonalPalette matching the Hue and Chroma of c.
func TonalPaletteFromColor(c ui.Color) *TonalPalette {
	hct := HCTFromColor(c)
	return TonalPaletteOf(hct.Hue, hct.Chroma)
}

// Tone returns the ui.Color for the given tone in [0, 100].
// Common M3 tone steps are 0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 95, 98, 99, 100.
func (p *TonalPalette) Tone(tone float64) ui.Color {
	intTone := int(math.Round(tone))

	p.mu.RLock()
	if c, ok := p.cache[intTone]; ok {
		p.mu.RUnlock()
		return c
	}
	p.mu.RUnlock()

	c := HCTOf(p.Hue, p.Chroma, tone).Color()

	p.mu.Lock()
	p.cache[intTone] = c
	p.mu.Unlock()

	return c
}

// HCT returns the HCT color representation for the given tone.
func (p *TonalPalette) HCT(tone float64) HCT {
	return HCTOf(p.Hue, p.Chroma, tone)
}
