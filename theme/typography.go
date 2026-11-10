package theme

// TextStyle defines standard font metrics for a typography role.
type TextStyle struct {
	Size       float32
	LineHeight float32
	Weight     int
}

// Standard Material Design 3 Typography styles. Material 3 sets display,
// headline, title-large and body at regular, and title-medium, title-small and
// every label at medium: weight is part of the role, not emphasis added on top.
var (
	DisplayLarge  = TextStyle{Size: 57, LineHeight: 64, Weight: 400}
	DisplayMedium = TextStyle{Size: 45, LineHeight: 52, Weight: 400}
	DisplaySmall  = TextStyle{Size: 36, LineHeight: 44, Weight: 400}

	HeadlineLarge  = TextStyle{Size: 32, LineHeight: 40, Weight: 400}
	HeadlineMedium = TextStyle{Size: 28, LineHeight: 36, Weight: 400}
	HeadlineSmall  = TextStyle{Size: 24, LineHeight: 32, Weight: 400}

	TitleLarge  = TextStyle{Size: 22, LineHeight: 28, Weight: 400}
	TitleMedium = TextStyle{Size: 16, LineHeight: 24, Weight: 500}
	TitleSmall  = TextStyle{Size: 14, LineHeight: 20, Weight: 500}

	BodyLarge  = TextStyle{Size: 16, LineHeight: 24, Weight: 400}
	BodyMedium = TextStyle{Size: 14, LineHeight: 20, Weight: 400}
	BodySmall  = TextStyle{Size: 12, LineHeight: 16, Weight: 400}

	LabelLarge  = TextStyle{Size: 14, LineHeight: 20, Weight: 500}
	LabelMedium = TextStyle{Size: 12, LineHeight: 16, Weight: 500}
	LabelSmall  = TextStyle{Size: 11, LineHeight: 16, Weight: 500}
)
