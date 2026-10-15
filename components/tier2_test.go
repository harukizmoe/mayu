package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

func TestTooltipAndSnackbar(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	snackbarActionClicked := false

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Gap(16).Children(func() {
			// 1. Tooltip attached to a button
			btn := ui.ButtonBase(c).Label("AnchorBtn")
			Tooltip(c, btn, "Helpful Hint", th)

			// 2. Snackbar feedback banner
			Snackbar(c, "Message archived", SnackbarOptions{
				ActionText: "Undo",
				OnAction: func() {
					snackbarActionClicked = true
				},
			}, th)
		})
	}, 400, 300)

	// Verify text rendering
	if !tt.HasText("Message archived") || !tt.HasText("Undo") {
		t.Fatal("Snackbar content missing")
	}

	// Click Snackbar Undo
	if err := tt.Click("Undo"); err != nil {
		t.Fatalf("failed clicking Snackbar action: %v", err)
	}
	if !snackbarActionClicked {
		t.Error("expected snackbarActionClicked == true")
	}
}

func TestDropdownMenu(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	menuOpen := true
	menuItemClicked := false

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Children(func() {
			btn := ui.ButtonBase(c).Label("AnchorBtn")
			Menu(c, btn, &menuOpen, func(m *MenuPanel) {
				m.Item("Copy Link", MenuItemOptions{
					Icon:     symbols.Send,
					Shortcut: "⌘C",
					OnClick: func() {
						menuItemClicked = true
					},
				})
				m.Divider()
				m.Item("Delete", MenuItemOptions{
					Icon:     symbols.Close,
					Disabled: true,
				})
			}, th)
		})
	}, 400, 300)

	if !tt.HasText("Copy Link") || !tt.HasText("Delete") {
		t.Fatal("Menu items missing")
	}

	if err := tt.Click("Copy Link"); err != nil {
		t.Fatalf("failed clicking Menu item: %v", err)
	}
	if !menuItemClicked {
		t.Error("expected menuItemClicked == true")
	}
	if menuOpen {
		t.Error("expected menuOpen == false after item chosen")
	}
}
