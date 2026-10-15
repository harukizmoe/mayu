package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

func TestTextField_Render(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	filledVal := "Alice"
	outlinedVal := ""

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Gap(12).Children(func() {
			NewTextField(c, &filledVal, "Username", TextFieldFilled, th)
			NewTextField(c, &outlinedVal, "Email", TextFieldOutlined, th).
				Placeholder("user@example.com")
		})
	}, 400, 300)

	if !tt.HasText("Username") {
		t.Fatal("filled text field label 'Username' not found")
	}
	if !tt.HasText("Email") {
		t.Fatal("outlined text field label 'Email' not found")
	}
	if filledVal != "Alice" {
		t.Fatalf("expected filledVal == Alice, got %s", filledVal)
	}
}

func TestTextArea(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	notes := "first line\nsecond line"

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Padding(16).Gap(12).Children(func() {
			NewTextArea(c, &notes, "Release notes", TextFieldFilled, th).
				Placeholder("What changed?")
			NewTextArea(c, &notes, "Notes", TextFieldOutlined, th)
		})
	}, 420, 360)

	for _, want := range []string{"Release notes", "Notes"} {
		if !tt.HasText(want) {
			t.Fatalf("the text area lacks %q: %q", want, tt.Texts())
		}
	}

	// A text area is taller than a field, so the two stack without either
	// shrinking to a single line's height.
	first, ok := tt.Find("Release notes")
	if !ok {
		t.Fatal("the first text area is not laid out")
	}
	second, ok := tt.Find("Notes")
	if !ok {
		t.Fatal("the second text area is not laid out")
	}
	if gap := second.Y - (first.Y + first.H); gap < 60 {
		t.Errorf("the text areas are %v apart, want a field's own height between them", gap)
	}
}
