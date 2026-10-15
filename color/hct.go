package color

import (
	"fmt"
	"math"

	"github.com/egoist/mygo/ui"
)

// HCT represents a color in the Google Material Design 3 HCT color space:
// - Hue: Perceptual hue from CAM16 [0, 360)
// - Chroma: Perceptual colorfulness from CAM16
// - Tone: Perceptual lightness/luminance L* from CIE Lab [0, 100]
type HCT struct {
	Hue    float64
	Chroma float64
	Tone   float64
}

// HCTOf constructs an HCT color from Hue, Chroma, and Tone.
func HCTOf(hue, chroma, tone float64) HCT {
	// Normalize hue to [0, 360)
	hue = math.Mod(hue, 360.0)
	if hue < 0 {
		hue += 360.0
	}
	if chroma < 0 {
		chroma = 0
	}
	tone = min(max(tone, 0), 100)
	return HCT{Hue: hue, Chroma: chroma, Tone: tone}
}

// HCTFromRGB constructs an HCT color from 8-bit sRGB components.
func HCTFromRGB(r, g, b uint8) HCT {
	cam := CAM16FromSRGB(r, g, b)
	_, y, _ := srgbToXYZ(r, g, b)
	tone := yToLstar(y)
	return HCT{
		Hue:    cam.Hue,
		Chroma: cam.Chroma,
		Tone:   tone,
	}
}

// HCTFromColor constructs an HCT color from a mygo/ui Color.
func HCTFromColor(c ui.Color) HCT {
	return HCTFromRGB(c.R, c.G, c.B)
}

// WithHue returns a copy of h with the given Hue.
func (h HCT) WithHue(hue float64) HCT {
	return HCTOf(hue, h.Chroma, h.Tone)
}

// WithChroma returns a copy of h with the given Chroma.
func (h HCT) WithChroma(chroma float64) HCT {
	return HCTOf(h.Hue, chroma, h.Tone)
}

// WithTone returns a copy of h with the given Tone.
func (h HCT) WithTone(tone float64) HCT {
	return HCTOf(h.Hue, h.Chroma, tone)
}

// Color converts HCT to a mygo/ui Color.
func (h HCT) Color() ui.Color {
	r, g, b := hctToSRGB(h.Hue, h.Chroma, h.Tone)
	return ui.RGB(r, g, b)
}

// Hex returns the #RRGGBB hexadecimal string for the HCT color.
func (h HCT) Hex() string {
	r, g, b := hctToSRGB(h.Hue, h.Chroma, h.Tone)
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// hctToSRGB converts (hue, chroma, tone) to 8-bit sRGB with gamut mapping.
func hctToSRGB(h, c, tone float64) (uint8, uint8, uint8) {
	if tone <= 0.0001 {
		return 0, 0, 0
	}
	if tone >= 99.9999 {
		return 255, 255, 255
	}

	targetY := lstarToY(tone)

	// helper: solve for RGB and test if inside sRGB gamut
	testGamut := func(testC float64) (bool, uint8, uint8, uint8) {
		if testC <= 0.0001 {
			// Achromatic D65
			vc := &DefaultViewingConditions
			x := (vc.WhitePoint[0] / 100.0) * targetY
			z := (vc.WhitePoint[2] / 100.0) * targetY
			r, g, b := xyzToSRGB(x, targetY, z)
			return true, r, g, b
		}

		vc := &DefaultViewingConditions
		hRad := h * math.Pi / 180.0
		cosH := math.Cos(hRad)
		sinH := math.Sin(hRad)
		eT := 0.25 * (math.Cos(hRad+2.0) + 3.8)
		gamma := (50000.0 / 13.0) * vc.Nc * vc.Ncb * eT

		lowJ, highJ := 0.0, 100.0
		var bestX, bestZ float64

		for range 30 {
			midJ := (lowJ + highJ) / 2.0
			if midJ <= 1e-5 {
				midJ = 1e-5
			}

			tVal := math.Pow(testC/(math.Sqrt(midJ/100.0)*math.Pow(1.64-math.Pow(0.29, vc.N), 0.73)), 1.0/0.9)
			aResp := vc.Aw * math.Pow(midJ/100.0, 1.0/(vc.C*vc.Z))
			p2 := aResp / vc.Nbb
			aScale := tVal * cosH / gamma
			bScale := tVal * sinH / gamma
			s := (11.0/23.0)*aScale + (108.0/23.0)*bScale
			u := (p2 - 0.305*s) / (1.0 + s)

			aComp := (u + 0.305) * aScale
			bComp := (u + 0.305) * bScale

			rA := 0.32786885*p2 + 0.32145403*aComp + 0.20527441*bComp
			gA := 0.32786885*p2 - 0.63506771*aComp - 0.18602994*bComp
			bA := 0.32786885*p2 - 0.15680684*aComp - 4.49037776*bComp

			if math.Abs(rA) >= 399.9 || math.Abs(gA) >= 399.9 || math.Abs(bA) >= 399.9 {
				lowJ = midJ
				continue
			}

			uncomp := func(rVal float64) float64 {
				sign := 1.0
				if rVal < 0 {
					sign = -1.0
				}
				absR := math.Abs(rVal)
				v := (27.13 * absR) / (400.0 - absR)
				return sign * (100.0 / vc.Fl) * math.Pow(v, 1.0/0.42)
			}

			rC := uncomp(rA)
			gC := uncomp(gA)
			bC := uncomp(bA)

			r := rC / vc.Drgb[0]
			g := gC / vc.Drgb[1]
			b := bC / vc.Drgb[2]

			// Inverse M16
			calcY := 0.38752654*r + 0.62144744*g - 0.00897399*b
			bestX = 1.86206786*r - 1.01125463*g + 0.14918678*b
			bestZ = -0.01584150*r - 0.03412294*g + 1.04996444*b

			if calcY < targetY {
				lowJ = midJ
			} else {
				highJ = midJ
			}
		}

		// Check linear sRGB bounds
		xn, yn, zn := bestX/100.0, targetY/100.0, bestZ/100.0
		rLin := 3.2404542*xn - 1.5371385*yn - 0.4985314*zn
		gLin := -0.9692660*xn + 1.8760108*yn + 0.0415560*zn
		bLin := 0.0556434*xn - 0.2040259*yn + 1.0572252*zn

		const eps = 0.005
		inGamut := rLin >= -eps && rLin <= 1.0+eps &&
			gLin >= -eps && gLin <= 1.0+eps &&
			bLin >= -eps && bLin <= 1.0+eps

		r, g, b := xyzToSRGB(bestX, targetY, bestZ)
		return inGamut, r, g, b
	}

	// 1. Try target chroma
	if ok, r, g, b := testGamut(c); ok {
		return r, g, b
	}

	// 2. Binary search chroma to gamut boundary
	lowC, highC := 0.0, c
	var bestR, bestG, bestB uint8
	// Fallback is achromatic
	_, bestR, bestG, bestB = testGamut(0.0)

	for range 20 {
		midC := (lowC + highC) / 2.0
		if ok, r, g, b := testGamut(midC); ok {
			bestR, bestG, bestB = r, g, b
			lowC = midC
		} else {
			highC = midC
		}
	}

	return bestR, bestG, bestB
}
