package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

func TestSelect(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	speed := "120 Hz"
	options := []SelectOption[string]{
		{Value: "60 Hz", Label: "60 Hz — power saving"},
		{Value: "120 Hz", Label: "120 Hz"},
		{Value: "180 Hz", Label: "180 Hz — overclocked"},
	}

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Padding(16).Width(320).Children(func() {
			Select(c, &speed, "Refresh rate", options, th)
		})
	}, 360, 420)

	// The trigger shows the chosen option's label, not its value.
	if !tt.HasText("Refresh rate") {
		t.Fatalf("the field's label is missing: %q", tt.Texts())
	}
	if !tt.HasText("120 Hz") {
		t.Fatalf("the chosen label is missing: %q", tt.Texts())
	}
	if tt.HasText("180 Hz — overclocked") {
		t.Fatal("the options are listed before the field is opened")
	}

	if err := tt.Click("Refresh rate"); err != nil {
		t.Fatalf("failed opening the select: %v", err)
	}
	if !tt.HasText("180 Hz — overclocked") {
		t.Fatalf("the options do not show when opened: %q", tt.Texts())
	}

	if err := tt.Click("180 Hz — overclocked"); err != nil {
		t.Fatalf("failed choosing an option: %v", err)
	}
	if speed != "180 Hz" {
		t.Errorf("expected the chosen value to be stored, got %q", speed)
	}

	tt.Frame()
	if tt.HasText("60 Hz — power saving") {
		t.Errorf("the popup stayed open: %q", tt.Texts())
	}
}

func TestAutocomplete(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	query := ""
	fonts := []string{"Helvetica", "Times", "Optima", "Courier"}

	// Chosen is a within-the-frame question, so it is asked from the view.
	picked := ""
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Padding(16).Width(320).Children(func() {
			ac := NewAutocomplete(c, &query, "Font", fonts, th)
			if v, ok := ac.Chosen(); ok {
				picked = v
			}
		})
	}, 360, 420)

	// A click opens the suggestions, all of them.
	if err := tt.Click("Font"); err != nil {
		t.Fatalf("failed focusing the field: %v", err)
	}
	if !tt.HasText("Courier") {
		t.Fatalf("a click does not show the suggestions: %q", tt.Texts())
	}

	// Typing filters them, those starting with the text first.
	tt.Type("ti")
	if query != "ti" {
		t.Fatalf("the field did not take the text, query=%q", query)
	}
	if !tt.HasText("Times") || tt.HasText("Courier") {
		t.Fatalf("typing does not filter: %q", tt.Texts())
	}

	// Enter takes the first suggestion into the field.
	tt.Key(0, ui.KeyEnter)
	tt.Frame()
	if query != "Times" {
		t.Errorf("Enter did not take the suggestion, got %q", query)
	}
	if picked != "Times" {
		t.Errorf("Chosen reported %q, want the suggestion taken", picked)
	}
}

// A click on a suggestion takes it too, when the popup was opened by a click.
func TestAutocompleteClickChooses(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	query := "Helvetica"
	fonts := []string{"Helvetica", "Times", "Optima", "Courier"}

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Padding(16).Width(320).Children(func() {
			NewAutocomplete(c, &query, "Font", fonts, th)
		})
	}, 360, 420)

	if err := tt.Click("Font"); err != nil {
		t.Fatalf("failed focusing the field: %v", err)
	}
	// Opened by a click, the popup shows every suggestion rather than
	// filtering to the text already in the field.
	if !tt.HasText("Courier") || !tt.HasText("Optima") {
		t.Fatalf("a click does not show them all: %q", tt.Texts())
	}

	r, ok := tt.Find("Courier")
	if !ok {
		t.Fatal("the suggestion is not laid out")
	}
	tt.ClickAt(r.X+r.W/2, r.Y+r.H/2)
	tt.Frame()
	if query != "Courier" {
		t.Errorf("the click did not take the suggestion, got %q", query)
	}
}

// The suggestions are filtered case-insensitively, those starting with the
// text first, and a blank query admits them all in order.
func TestAutocompleteMatching(t *testing.T) {
	all := []string{"Helvetica", "Times", "Optima", "Courier"}

	for _, tc := range []struct {
		query string
		want  []string
	}{
		{"", all},
		// Helvetica contains "ti" too, so it follows those starting with it.
		{"ti", []string{"Times", "Helvetica", "Optima"}},
		{"TI", []string{"Times", "Helvetica", "Optima"}},
		{"  ", all},
		{"zz", nil},
	} {
		got := matching(all, tc.query)
		if len(got) != len(tc.want) {
			t.Errorf("matching(%q) = %v, want %v", tc.query, got, tc.want)
			continue
		}
		for i := range got {
			if got[i] != tc.want[i] {
				t.Errorf("matching(%q) = %v, want %v", tc.query, got, tc.want)
				break
			}
		}
	}
}
