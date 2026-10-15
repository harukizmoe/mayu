package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

func TestExtraComponents(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	chipClicked := false
	sliderVal := 50.0

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Gap(8).Children(func() {
			NewFAB(c, symbols.Add, th)
			NewExtendedFAB(c, symbols.Send, "Send Message", th)
			NewIconButton(c, symbols.Settings, IconButtonFilled, th)

			NewChip(c, "Filter Tag", th).
				Selected(true).
				OnClick(func() {
					chipClicked = true
				})

			NewSlider(c, &sliderVal, 0, 100, th)

			HorizontalDivider(c, th)
			VerticalDivider(c, th)
			LinearProgress(c, 0.65, th)
			Badge(c, 5, th)
			Badge(c, 0, th)
		})
	}, 500, 600)

	// Verify text rendering
	if !tt.HasText("Send Message") {
		t.Fatal("extended FAB text not found")
	}
	if !tt.HasText("Filter Tag") {
		t.Fatal("chip text not found")
	}

	// Click Chip
	if err := tt.Click("Filter Tag"); err != nil {
		t.Fatalf("failed clicking Chip: %v", err)
	}
	if !chipClicked {
		t.Error("expected chipClicked == true")
	}
}

func TestAlertDialog(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	dialogOpen := true
	dialogConfirmed := false

	tt := ui.NewTester(func(c *ui.Context) {
		AlertDialog(c, &dialogOpen, AlertDialogOptions{
			Title:       "Confirm Action",
			Body:        "Are you sure you want to proceed?",
			ConfirmText: "Confirm",
			OnConfirm: func() {
				dialogConfirmed = true
			},
			DismissText: "Cancel",
		}, th)
	}, 400, 300)

	if !tt.HasText("Confirm Action") {
		t.Fatal("dialog title not found")
	}
	if !tt.HasText("Confirm") {
		t.Fatal("dialog Confirm button not found")
	}

	if err := tt.Click("Confirm"); err != nil {
		t.Fatalf("failed clicking Confirm: %v", err)
	}
	if !dialogConfirmed {
		t.Error("expected dialogConfirmed == true")
	}
	if dialogOpen {
		t.Error("expected dialogOpen == false after Confirm")
	}
}
