package color

import "math"

// ViewingConditions contains parameters describing the visual environment under which a color is viewed.
type ViewingConditions struct {
	WhitePoint        [3]float64
	AdaptingLuminance float64
	BackgroundLstar   float64
	Surround          float64
	N                 float64
	Aw                float64
	Nbb               float64
	Ncb               float64
	C                 float64
	Nc                float64
	Z                 float64
	Drgb              [3]float64
	Fl                float64
}

// DefaultViewingConditions represents standard sRGB viewing conditions:
// White point D65, 200 cd/m^2 display luminance in standard office environment (LA = 11.7257).
var DefaultViewingConditions = func() ViewingConditions {
	whitePoint := [3]float64{95.047, 100.0, 108.883}
	adaptingLuminance := 11.72567793503785
	backgroundLstar := 50.0

	yb := lstarToY(backgroundLstar)
	n := yb / whitePoint[1]
	z := 1.48 + math.Sqrt(n)
	nbb := 0.725 * math.Pow(1.0/n, 0.2)
	ncb := nbb
	c := 0.69
	nc := 1.0
	f := 1.0

	// rgbW = M16 * whitePoint
	rgbW := [3]float64{
		0.401288*whitePoint[0] + 0.650173*whitePoint[1] - 0.051461*whitePoint[2],
		-0.250268*whitePoint[0] + 1.204414*whitePoint[1] + 0.045854*whitePoint[2],
		-0.002079*whitePoint[0] + 0.048952*whitePoint[1] + 0.953127*whitePoint[2],
	}

	k := 1.0 / (5.0*adaptingLuminance + 1.0)
	k4 := k * k * k * k
	fl := 0.2*k4*(5.0*adaptingLuminance) + 0.1*(1.0-k4)*(1.0-k4)*math.Cbrt(5.0*adaptingLuminance)

	d := f * (1.0 - (1.0/3.6)*math.Exp((-adaptingLuminance-42.0)/92.0))
	if d < 0 {
		d = 0
	} else if d > 1 {
		d = 1
	}

	drgb := [3]float64{
		d*(whitePoint[1]/rgbW[0]) + (1.0 - d),
		d*(whitePoint[1]/rgbW[1]) + (1.0 - d),
		d*(whitePoint[1]/rgbW[2]) + (1.0 - d),
	}

	rgbCW := [3]float64{rgbW[0] * drgb[0], rgbW[1] * drgb[1], rgbW[2] * drgb[2]}
	compress := func(v float64) float64 {
		val := math.Pow(fl*math.Abs(v)/100.0, 0.42)
		return math.Copysign(400.0*val/(val+27.13), v)
	}
	rgbAW := [3]float64{compress(rgbCW[0]), compress(rgbCW[1]), compress(rgbCW[2])}
	aw := (40.0*rgbAW[0] + 20.0*rgbAW[1] + rgbAW[2]) / 20.0 * nbb

	return ViewingConditions{
		WhitePoint:        whitePoint,
		AdaptingLuminance: adaptingLuminance,
		BackgroundLstar:   backgroundLstar,
		Surround:          1.0,
		N:                 n,
		Aw:                aw,
		Nbb:               nbb,
		Ncb:               ncb,
		C:                 c,
		Nc:                nc,
		Z:                 z,
		Drgb:              drgb,
		Fl:                fl,
	}
}()

// CAM16 represents a color in the CAM16 color appearance model.
type CAM16 struct {
	Hue    float64 // Hue angle in degrees [0, 360)
	Chroma float64 // Colorfulness relative to brightness
	J      float64 // Lightness [0, 100]
	Q      float64 // Brightness
	M      float64 // Colorfulness
	S      float64 // Saturation
}

// CAM16FromXYZ computes CAM16 attributes from CIE XYZ coordinates under default viewing conditions.
func CAM16FromXYZ(x, y, z float64) CAM16 {
	vc := &DefaultViewingConditions

	// Transform to cone response
	r := 0.401288*x + 0.650173*y - 0.051461*z
	g := -0.250268*x + 1.204414*y + 0.045854*z
	b := -0.002079*x + 0.048952*y + 0.953127*z

	// Chromatic adaptation
	rC := r * vc.Drgb[0]
	gC := g * vc.Drgb[1]
	bC := b * vc.Drgb[2]

	// Nonlinear compression
	compress := func(v float64) float64 {
		val := math.Pow(vc.Fl*math.Abs(v)/100.0, 0.42)
		return math.Copysign(400.0*val/(val+27.13), v)
	}
	rA := compress(rC)
	gA := compress(gC)
	bA := compress(bC)

	// Opponent color components
	a := rA + (-12.0*gA+bA)/11.0
	bComp := (rA + gA - 2.0*bA) / 9.0
	u := (20.0*rA + 20.0*gA + 21.0*bA) / 20.0
	p2 := (40.0*rA + 20.0*gA + bA) / 20.0

	// Hue angle
	atanDegrees := math.Atan2(bComp, a) * 180.0 / math.Pi
	if atanDegrees < 0 {
		atanDegrees += 360.0
	}
	hue := atanDegrees

	// Eccentricity & achromatic response
	hRad := hue * math.Pi / 180.0
	eT := 0.25 * (math.Cos(hRad+2.0) + 3.8)
	aResp := p2 * vc.Nbb

	j := 100.0 * math.Pow(aResp/vc.Aw, vc.C*vc.Z)
	alpha := (math.Hypot(a, bComp) * eT * (50000.0 / 13.0) * vc.Nc * vc.Ncb) / (u + 0.305)
	c := math.Pow(alpha, 0.9) * math.Sqrt(j/100.0) * math.Pow(1.64-math.Pow(0.29, vc.N), 0.73)

	return CAM16{
		Hue:    hue,
		Chroma: c,
		J:      j,
	}
}

// CAM16FromSRGB computes CAM16 from sRGB components.
func CAM16FromSRGB(r, g, b uint8) CAM16 {
	x, y, z := srgbToXYZ(r, g, b)
	return CAM16FromXYZ(x, y, z)
}

// ToXYZ converts CAM16 back to CIE XYZ coordinates under default viewing conditions.
func (c CAM16) ToXYZ() (float64, float64, float64) {
	vc := &DefaultViewingConditions

	if c.J <= 0.0 || c.Chroma <= 0.0 {
		if c.J <= 0.0 {
			return 0, 0, 0
		}
		// Achromatic case
		y := lstarToY(c.J)
		return (vc.WhitePoint[0] / 100.0) * y, y, (vc.WhitePoint[2] / 100.0) * y
	}

	hRad := c.Hue * math.Pi / 180.0
	cosH := math.Cos(hRad)
	sinH := math.Sin(hRad)

	t := math.Pow(c.Chroma/(math.Sqrt(c.J/100.0)*math.Pow(1.64-math.Pow(0.29, vc.N), 0.73)), 1.0/0.9)
	eT := 0.25 * (math.Cos(hRad+2.0) + 3.8)
	a := vc.Aw * math.Pow(c.J/100.0, 1.0/(vc.C*vc.Z))
	p2 := a / vc.Nbb
	gamma := (50000.0 / 13.0) * vc.Nc * vc.Ncb * eT

	aScale := t * cosH / gamma
	bScale := t * sinH / gamma
	s := (11.0/23.0)*aScale + (108.0/23.0)*bScale
	u := (p2 - 0.305*s) / (1.0 + s)

	aComp := (u + 0.305) * aScale
	bComp := (u + 0.305) * bScale

	// Inverse M_ab transform: solve Ra, Ga, Ba
	rA := 0.32786885*p2 + 0.32145403*aComp + 0.20527441*bComp
	gA := 0.32786885*p2 - 0.63506771*aComp - 0.18602994*bComp
	bA := 0.32786885*p2 - 0.15680684*aComp - 4.49037776*bComp

	uncompress := func(rVal float64) float64 {
		sign := 1.0
		if rVal < 0 {
			sign = -1.0
		}
		absR := math.Abs(rVal)
		if absR >= 399.9999 {
			absR = 399.9999
		}
		v := (27.13 * absR) / (400.0 - absR)
		return sign * (100.0 / vc.Fl) * math.Pow(v, 1.0/0.42)
	}

	rC := uncompress(rA)
	gC := uncompress(gA)
	bC := uncompress(bA)

	r := rC / vc.Drgb[0]
	g := gC / vc.Drgb[1]
	b := bC / vc.Drgb[2]

	// Inverse M16
	x := 1.86206786*r - 1.01125463*g + 0.14918678*b
	y := 0.38752654*r + 0.62144744*g - 0.00897399*b
	z := -0.01584150*r - 0.03412294*g + 1.04996444*b

	return x, y, z
}

// srgbToXYZ converts 8-bit sRGB to CIE 1931 XYZ.
func srgbToXYZ(r, g, b uint8) (float64, float64, float64) {
	pivot := func(c uint8) float64 {
		norm := float64(c) / 255.0
		if norm <= 0.04045 {
			return norm / 12.92 * 100.0
		}
		return math.Pow((norm+0.055)/1.055, 2.4) * 100.0
	}

	rLin := pivot(r)
	gLin := pivot(g)
	bLin := pivot(b)

	x := 0.41233895*rLin + 0.35762064*gLin + 0.18051042*bLin
	y := 0.21260000*rLin + 0.71520000*gLin + 0.07220000*bLin
	z := 0.01932141*rLin + 0.11916382*gLin + 0.95034478*bLin

	return x, y, z
}

// xyzToSRGB converts CIE 1931 XYZ to 8-bit sRGB.
func xyzToSRGB(x, y, z float64) (uint8, uint8, uint8) {
	xn, yn, zn := x/100.0, y/100.0, z/100.0

	rLin := 3.2404542*xn - 1.5371385*yn - 0.4985314*zn
	gLin := -0.9692660*xn + 1.8760108*yn + 0.0415560*zn
	bLin := 0.0556434*xn - 0.2040259*yn + 1.0572252*zn

	unpivot := func(c float64) uint8 {
		if c < 0 {
			c = 0
		} else if c > 1 {
			c = 1
		}
		var s float64
		if c <= 0.0031308 {
			s = 12.92 * c
		} else {
			s = 1.055*math.Pow(c, 1.0/2.4) - 0.055
		}
		val := math.Round(s * 255.0)
		if val < 0 {
			return 0
		}
		if val > 255 {
			return 255
		}
		return uint8(val)
	}

	return unpivot(rLin), unpivot(gLin), unpivot(bLin)
}

func yToLstar(y float64) float64 {
	yNorm := y / 100.0
	e := 216.0 / 24389.0
	k := 24389.0 / 27.0
	if yNorm > e {
		return 116.0*math.Cbrt(yNorm) - 16.0
	}
	return k * yNorm
}

func lstarToY(lstar float64) float64 {
	k := 24389.0 / 27.0
	if lstar > 8.0 {
		return 100.0 * math.Pow((lstar+16.0)/116.0, 3.0)
	}
	return 100.0 * (lstar / k)
}
