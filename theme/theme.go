package theme

// Theme bundles the color scheme and density configuration of Mayu.
type Theme struct {
	Scheme  ColorScheme
	Density Density
}

// NewTheme creates a Theme with the given ColorScheme and Density.
func NewTheme(scheme ColorScheme, density Density) *Theme {
	return &Theme{
		Scheme:  scheme,
		Density: density,
	}
}
