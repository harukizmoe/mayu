package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

func TestTier3Components(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	accordionOpen := false
	searchQuery := "Material Design"

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Gap(16).Children(func() {
			// 1. Avatar
			Avatar(c, symbols.Settings, th)
			InitialsAvatar(c, "John Doe", th)

			// 2. Accordion
			Accordion(c, &accordionOpen, "System Settings", func() {
				ui.Text(c, "Accordion Detailed Content")
			}, th)

			// 3. Search Bar
			NewSearchBar(c, &searchQuery, "Search settings...", th)
		})
	}, 500, 500)

	// Verify text rendering
	if !tt.HasText("JD") {
		t.Fatal("initials 'JD' not found in avatar")
	}
	if !tt.HasText("System Settings") {
		t.Fatal("accordion header 'System Settings' not found")
	}
	if tt.HasText("Accordion Detailed Content") {
		t.Fatal("accordion content should be hidden initially")
	}

	// Click Accordion header to expand
	if err := tt.Click("System Settings"); err != nil {
		t.Fatalf("failed clicking Accordion header: %v", err)
	}
	if !accordionOpen {
		t.Error("expected accordionOpen == true")
	}

	// Re-render tester to verify content shown
	tt.Frame()
	if !tt.HasText("Accordion Detailed Content") {
		t.Fatal("accordion content not shown after expansion")
	}
}
