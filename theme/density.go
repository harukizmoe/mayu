package theme

// Density specifies the UI layout density scale for keyboard/mouse environments.
type Density int

const (
	// DensityStandard is the default Material 3 desktop density.
	DensityStandard Density = iota
	// DensityCompact is designed for data-heavy or tool-dense desktop windows.
	DensityCompact
	// DensityComfortable provides generous touch-friendly spacing and larger controls.
	DensityComfortable
)

// ControlHeight returns the standard height for buttons and inputs.
func (d Density) ControlHeight() float32 {
	switch d {
	case DensityCompact:
		return 32
	case DensityComfortable:
		return 48
	default:
		return 40
	}
}

// SpacingUnit returns the fundamental layout spacing unit in DIPs.
func (d Density) SpacingUnit() float32 {
	switch d {
	case DensityCompact:
		return 3
	case DensityComfortable:
		return 5
	default:
		return 4
	}
}

// BaseFontSize returns the default font size in DIPs.
func (d Density) BaseFontSize() float32 {
	switch d {
	case DensityCompact:
		return 13
	case DensityComfortable:
		return 15
	default:
		return 14
	}
}

// ButtonPaddingX returns the horizontal padding for buttons without icons.
func (d Density) ButtonPaddingX() float32 {
	switch d {
	case DensityCompact:
		return 16
	case DensityComfortable:
		return 28
	default:
		return 24
	}
}

// ButtonPaddingXLeadingIcon returns the leading horizontal padding when an icon is present.
func (d Density) ButtonPaddingXLeadingIcon() float32 {
	switch d {
	case DensityCompact:
		return 12
	case DensityComfortable:
		return 20
	default:
		return 16
	}
}
