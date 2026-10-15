package mayu

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/components"
	"github.com/harukizmoe/mayu/theme"
)

// Density re-exports theme.Density for convenient access.
type Density = theme.Density

const (
	DensityStandard    Density = theme.DensityStandard
	DensityCompact     Density = theme.DensityCompact
	DensityComfortable Density = theme.DensityComfortable
)

// ThemeOptions configures the Material 3 design system for the window.
type ThemeOptions struct {
	SeedColor ui.Color
	Dark      bool
	Density   theme.Density
}

type themeKey struct{}

// Use applies Material 3 ThemeOptions to the window Context.
// Call it first in every frame of the root view.
func Use(c *ui.Context, opts ThemeOptions) *theme.Theme {
	seed := opts.SeedColor
	if seed == (ui.Color{}) {
		seed = ui.Hex("#6750A4") // Canonical M3 seed
	}

	scheme := theme.SchemeFromSeed(seed, opts.Dark)
	t := theme.NewTheme(scheme, opts.Density)

	// Propagate M3 colors and metrics to upstream MyGo ui.Theme
	c.SetTheme(&ui.Theme{
		Dark:           opts.Dark,
		Background:     scheme.Background,
		Surface:        scheme.SurfaceContainer,
		SurfaceHover:   scheme.SurfaceContainerHigh,
		SurfacePressed: scheme.SurfaceContainerHighest,
		Border:         scheme.OutlineVariant,
		Text:           scheme.OnSurface,
		TextMuted:      scheme.OnSurfaceVariant,
		Accent:         scheme.Primary,
		AccentHover:    scheme.PrimaryContainer,
		AccentPressed:  scheme.PrimaryContainer,
		AccentText:     scheme.OnPrimary,
		Danger:         scheme.Error,
		Warning:        ui.Hex("#FF9800"),
		Success:        ui.Hex("#4CAF50"),
		Selection:      scheme.PrimaryContainer,
		Focus:          scheme.Primary,
		Inverse:        scheme.InverseSurface,
		InverseText:    scheme.InverseOnSurface,
		Radius:         opts.Density.ControlHeight() / 2,
		Spacing:        opts.Density.SpacingUnit(),
		FontSize:       opts.Density.BaseFontSize(),
	})

	// Inject into ui.Local at root for child Mayu components
	*ui.Local(c.Root(), themeKey{}, func() *theme.Theme { return t }) = t

	return t
}

// CurrentTheme retrieves the active Mayu Theme from the Context, or creates a default one.
func CurrentTheme(c *ui.Context) *theme.Theme {
	loc := ui.Local(c.Root(), themeKey{}, func() *theme.Theme {
		scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), c.Theme().Dark)
		return theme.NewTheme(scheme, theme.DensityStandard)
	})
	if loc != nil && *loc != nil {
		return *loc
	}
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), c.Theme().Dark)
	return theme.NewTheme(scheme, theme.DensityStandard)
}

// Scheme returns the active M3 ColorScheme from the Context.
func Scheme(c *ui.Context) theme.ColorScheme {
	return CurrentTheme(c).Scheme
}

// FilledButton constructs an M3 high-emphasis Filled Button.
func FilledButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantFilled, CurrentTheme(c))
}

// ElevatedButton constructs an M3 Elevated Button with tonal elevation.
func ElevatedButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantElevated, CurrentTheme(c))
}

// FilledTonalButton constructs an M3 medium-emphasis Filled Tonal Button.
func FilledTonalButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantFilledTonal, CurrentTheme(c))
}

// OutlinedButton constructs an M3 medium-low emphasis Outlined Button.
func OutlinedButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantOutlined, CurrentTheme(c))
}

// TextButton constructs an M3 low-emphasis Text Button.
func TextButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantText, CurrentTheme(c))
}
