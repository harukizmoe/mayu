package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

func TestSideSheet(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	open := false
	applied := false

	render := func(c *ui.Context) {
		ui.Box(c).Fill().Children(func() {
			ui.Text(c, "Gallery Behind The Sheet")

			SideSheet(c, &open, SideSheetOptions{
				Title:    "Wallpaper Settings",
				Subtitle: "DP-1 (Main)",
				Actions: func() {
					NewButton(c, "Apply", VariantFilled, th).OnClick(func() { applied = true })
				},
			}, th, func() {
				ui.Text(c, "Scale 1.20x")
			})
		})
	}

	tt := ui.NewTester(render, 900, 600)

	// Closed: the sheet contributes nothing to the tree.
	if tt.HasText("Wallpaper Settings") {
		t.Fatal("side sheet rendered while closed")
	}
	if !tt.HasText("Gallery Behind The Sheet") {
		t.Fatal("content behind the sheet should render while closed")
	}

	// Open it.
	open = true
	tt.Frame()

	if !tt.HasText("Wallpaper Settings") {
		t.Fatal("side sheet title not found after opening")
	}
	if !tt.HasText("DP-1 (Main)") {
		t.Fatal("side sheet subtitle not found")
	}
	if !tt.HasText("Scale 1.20x") {
		t.Fatal("side sheet content not found")
	}
	if !tt.HasText("Gallery Behind The Sheet") {
		t.Fatal("content behind a modal sheet should still render, just dimmed")
	}

	// Actions live in the pinned footer and must be clickable.
	if err := tt.Click("Apply"); err != nil {
		t.Fatalf("failed clicking the sheet's footer action: %v", err)
	}
	if !applied {
		t.Error("expected the footer action to run")
	}

	// Escape dismisses the sheet and clears the open flag, as AlertDialog does.
	tt.Key(0, ui.KeyEscape)
	tt.Frame()
	if open {
		t.Error("expected Escape to close the side sheet")
	}
	if tt.HasText("Wallpaper Settings") {
		t.Fatal("side sheet still rendered after Escape")
	}
}

// A sheet without a title or footer is still a sheet: content only.
func TestSideSheet_ContentOnly(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), true)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	open := true
	tt := ui.NewTester(func(c *ui.Context) {
		SideSheet(c, &open, SideSheetOptions{}, th, func() {
			ui.Text(c, "Just The Content")
		})
	}, 900, 600)

	if !tt.HasText("Just The Content") {
		t.Fatal("content-only side sheet did not render its content")
	}
}
