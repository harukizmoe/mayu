package theme

// TextStyle defines standard font metrics for a typography role.
type TextStyle struct {
	Size       float32
	LineHeight float32
}

// Standard Material Design 3 Typography styles.
var (
	DisplayLarge  = TextStyle{Size: 57, LineHeight: 64}
	DisplayMedium = TextStyle{Size: 45, LineHeight: 52}
	DisplaySmall  = TextStyle{Size: 36, LineHeight: 44}

	HeadlineLarge  = TextStyle{Size: 32, LineHeight: 40}
	HeadlineMedium = TextStyle{Size: 28, LineHeight: 36}
	HeadlineSmall  = TextStyle{Size: 24, LineHeight: 32}

	TitleLarge  = TextStyle{Size: 22, LineHeight: 28}
	TitleMedium = TextStyle{Size: 16, LineHeight: 24}
	TitleSmall  = TextStyle{Size: 14, LineHeight: 20}

	BodyLarge  = TextStyle{Size: 16, LineHeight: 24}
	BodyMedium = TextStyle{Size: 14, LineHeight: 20}
	BodySmall  = TextStyle{Size: 12, LineHeight: 16}

	LabelLarge  = TextStyle{Size: 14, LineHeight: 20}
	LabelMedium = TextStyle{Size: 12, LineHeight: 16}
	LabelSmall  = TextStyle{Size: 11, LineHeight: 16}
)
