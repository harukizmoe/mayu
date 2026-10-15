package theme

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/color"
)

// ColorScheme defines the full set of Material Design 3 semantic color tokens.
type ColorScheme struct {
	Dark bool

	// Primary roles
	Primary            ui.Color
	OnPrimary          ui.Color
	PrimaryContainer   ui.Color
	OnPrimaryContainer ui.Color
	InversePrimary     ui.Color

	// Secondary roles
	Secondary            ui.Color
	OnSecondary          ui.Color
	SecondaryContainer   ui.Color
	OnSecondaryContainer ui.Color

	// Tertiary roles
	Tertiary            ui.Color
	OnTertiary          ui.Color
	TertiaryContainer   ui.Color
	OnTertiaryContainer ui.Color

	// Error roles
	Error            ui.Color
	OnError          ui.Color
	ErrorContainer   ui.Color
	OnErrorContainer ui.Color

	// Neutral / Surface roles
	Background       ui.Color
	OnBackground     ui.Color
	Surface          ui.Color
	OnSurface        ui.Color
	SurfaceVariant   ui.Color
	OnSurfaceVariant ui.Color
	Outline          ui.Color
	OutlineVariant   ui.Color
	Shadow           ui.Color
	Scrim            ui.Color
	InverseSurface   ui.Color
	InverseOnSurface ui.Color

	// M3 Tonal Elevation Surface Containers
	SurfaceContainerLowest  ui.Color
	SurfaceContainerLow     ui.Color
	SurfaceContainer        ui.Color
	SurfaceContainerHigh    ui.Color
	SurfaceContainerHighest ui.Color
}

// LightScheme creates a light ColorScheme from an M3 DynamicScheme.
func LightScheme(dyn *color.DynamicScheme) ColorScheme {
	return ColorScheme{
		Dark: false,

		Primary:            dyn.Primary.Tone(40),
		OnPrimary:          dyn.Primary.Tone(100),
		PrimaryContainer:   dyn.Primary.Tone(90),
		OnPrimaryContainer: dyn.Primary.Tone(10),
		InversePrimary:     dyn.Primary.Tone(80),

		Secondary:            dyn.Secondary.Tone(40),
		OnSecondary:          dyn.Secondary.Tone(100),
		SecondaryContainer:   dyn.Secondary.Tone(90),
		OnSecondaryContainer: dyn.Secondary.Tone(10),

		Tertiary:            dyn.Tertiary.Tone(40),
		OnTertiary:          dyn.Tertiary.Tone(100),
		TertiaryContainer:   dyn.Tertiary.Tone(90),
		OnTertiaryContainer: dyn.Tertiary.Tone(10),

		Error:            dyn.Error.Tone(40),
		OnError:          dyn.Error.Tone(100),
		ErrorContainer:   dyn.Error.Tone(90),
		OnErrorContainer: dyn.Error.Tone(10),

		Background:       dyn.Neutral.Tone(98),
		OnBackground:     dyn.Neutral.Tone(10),
		Surface:          dyn.Neutral.Tone(98),
		OnSurface:        dyn.Neutral.Tone(10),
		SurfaceVariant:   dyn.NeutralVariant.Tone(90),
		OnSurfaceVariant: dyn.NeutralVariant.Tone(30),
		Outline:          dyn.NeutralVariant.Tone(50),
		OutlineVariant:   dyn.NeutralVariant.Tone(80),
		Shadow:           dyn.Neutral.Tone(0),
		Scrim:            dyn.Neutral.Tone(0),
		InverseSurface:   dyn.Neutral.Tone(20),
		InverseOnSurface: dyn.Neutral.Tone(95),

		SurfaceContainerLowest:  dyn.Neutral.Tone(100),
		SurfaceContainerLow:     dyn.Neutral.Tone(96),
		SurfaceContainer:        dyn.Neutral.Tone(94),
		SurfaceContainerHigh:    dyn.Neutral.Tone(92),
		SurfaceContainerHighest: dyn.Neutral.Tone(90),
	}
}

// DarkScheme creates a dark ColorScheme from an M3 DynamicScheme.
func DarkScheme(dyn *color.DynamicScheme) ColorScheme {
	return ColorScheme{
		Dark: true,

		Primary:            dyn.Primary.Tone(80),
		OnPrimary:          dyn.Primary.Tone(20),
		PrimaryContainer:   dyn.Primary.Tone(30),
		OnPrimaryContainer: dyn.Primary.Tone(90),
		InversePrimary:     dyn.Primary.Tone(40),

		Secondary:            dyn.Secondary.Tone(80),
		OnSecondary:          dyn.Secondary.Tone(20),
		SecondaryContainer:   dyn.Secondary.Tone(30),
		OnSecondaryContainer: dyn.Secondary.Tone(90),

		Tertiary:            dyn.Tertiary.Tone(80),
		OnTertiary:          dyn.Tertiary.Tone(20),
		TertiaryContainer:   dyn.Tertiary.Tone(30),
		OnTertiaryContainer: dyn.Tertiary.Tone(90),

		Error:            dyn.Error.Tone(80),
		OnError:          dyn.Error.Tone(20),
		ErrorContainer:   dyn.Error.Tone(30),
		OnErrorContainer: dyn.Error.Tone(90),

		Background:       dyn.Neutral.Tone(6),
		OnBackground:     dyn.Neutral.Tone(90),
		Surface:          dyn.Neutral.Tone(6),
		OnSurface:        dyn.Neutral.Tone(90),
		SurfaceVariant:   dyn.NeutralVariant.Tone(30),
		OnSurfaceVariant: dyn.NeutralVariant.Tone(80),
		Outline:          dyn.NeutralVariant.Tone(60),
		OutlineVariant:   dyn.NeutralVariant.Tone(30),
		Shadow:           dyn.Neutral.Tone(0),
		Scrim:            dyn.Neutral.Tone(0),
		InverseSurface:   dyn.Neutral.Tone(90),
		InverseOnSurface: dyn.Neutral.Tone(20),

		SurfaceContainerLowest:  dyn.Neutral.Tone(4),
		SurfaceContainerLow:     dyn.Neutral.Tone(10),
		SurfaceContainer:        dyn.Neutral.Tone(12),
		SurfaceContainerHigh:    dyn.Neutral.Tone(17),
		SurfaceContainerHighest: dyn.Neutral.Tone(22),
	}
}

// SchemeFromSeed generates a ColorScheme directly from a seed color and dark flag.
func SchemeFromSeed(seed ui.Color, dark bool) ColorScheme {
	dyn := color.NewDynamicScheme(seed)
	if dark {
		return DarkScheme(dyn)
	}
	return LightScheme(dyn)
}

// Accents for what stands on a color scheme's InverseSurface, where the
// scheme's own roles are too dark to read. They are the tones Material 3
// reserves for that surface.
var (
	// InverseSuccess marks what went well, as "Saved".
	InverseSuccess = ui.Hex("#6DD58C")
	// InverseError marks what failed, as "Upload failed".
	InverseError = ui.Hex("#FFB4AB")
	// InverseWarning marks what calls for care.
	InverseWarning = ui.Hex("#FFD8A8")
)

// Accents for what stands on the app's own surfaces, where InverseSuccess and
// its siblings are too pale to read. Error and the rest come from the scheme;
// these are the two roles Material 3 leaves to the app.
var (
	// Success marks what went well, as "Saved".
	Success = ui.Hex("#2E6C3F")
	// Warning marks what calls for care.
	Warning = ui.Hex("#8A5300")
)
