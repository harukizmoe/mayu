package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

func TestTier1Components(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	segmentedChoice := 0
	itemClicked := false
	railChoice := 1

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Gap(16).Children(func() {
			// Segmented Button
			SegmentedButton(c, &segmentedChoice, []SegmentOption{
				{Label: "Day", Icon: symbols.LightMode},
				{Label: "Week"},
				{Label: "Month"},
			}, th)

			// List Item
			NewListItem(c, "Wi-Fi Settings", ListItemOptions{
				SupportingText: "Connected to Home-5G",
				Leading: func() {
					ui.Icon(c, symbols.Settings).Size(20, 20)
				},
				Trailing: func() {
					ui.Text(c, "On")
				},
				OnClick: func() {
					itemClicked = true
				},
			}, th)

			// Navigation Rail
			NavigationRail(c, NavigationRailOptions{
				Selected: &railChoice,
				Destinations: []NavDestination{
					{Label: "Inbox", Icon: symbols.Send},
					{Label: "Favorites", Icon: symbols.Favorite},
				},
			}, th)
		})
	}, 600, 500)

	// 1. Verify text rendering
	if !tt.HasText("Day") || !tt.HasText("Week") || !tt.HasText("Month") {
		t.Fatal("SegmentedButton labels missing")
	}
	if !tt.HasText("Wi-Fi Settings") || !tt.HasText("Connected to Home-5G") {
		t.Fatal("ListItem texts missing")
	}
	if !tt.HasText("Inbox") || !tt.HasText("Favorites") {
		t.Fatal("NavigationRail destinations missing")
	}

	// 2. Click Segmented option "Week"
	if err := tt.Click("Week"); err != nil {
		t.Fatalf("failed clicking Segmented option: %v", err)
	}
	if segmentedChoice != 1 {
		t.Errorf("expected segmentedChoice == 1, got %d", segmentedChoice)
	}

	// 3. Click ListItem
	if err := tt.Click("Wi-Fi Settings"); err != nil {
		t.Fatalf("failed clicking ListItem: %v", err)
	}
	if !itemClicked {
		t.Error("expected itemClicked == true after click")
	}

	// 4. Click Navigation Rail destination "Inbox"
	if err := tt.Click("Inbox"); err != nil {
		t.Fatalf("failed clicking NavigationRail item: %v", err)
	}
	if railChoice != 0 {
		t.Errorf("expected railChoice == 0, got %d", railChoice)
	}
}
