package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
)

func TestGallery_InteractiveFeatures(t *testing.T) {
	app := newGalleryApp()

	tt := ui.NewTester(app.view, 1200, 1800)

	// 1. Verify desktop skeleton and navigation
	if !tt.HasText("Overview") || !tt.HasText("Favorites") || !tt.HasText("Config") {
		t.Fatal("NavigationRail destinations not found")
	}

	// 2. Verify segmented button
	if !tt.HasText("Daily") || !tt.HasText("Weekly") {
		t.Fatal("SegmentedButton labels not found")
	}

	// 3. Verify capsule sliders
	if !tt.HasText("开灯") {
		t.Fatal("Horizontal capsule slider '开灯' not found")
	}

	// 4. Verify list items & avatars
	if !tt.HasText("Sarah Connor") {
		t.Fatal("ListItem 'Sarah Connor' not found")
	}

	// 5. Verify accordion
	if !tt.HasText("Advanced System Configuration") {
		t.Fatal("Accordion header not found")
	}

	// 6. Click Segmented option "Weekly"
	if err := tt.Click("Weekly"); err != nil {
		t.Fatalf("failed clicking Segmented option: %v", err)
	}
	if app.segmentedIndex != 1 {
		t.Errorf("expected segmentedIndex == 1, got %d", app.segmentedIndex)
	}

	// 7. Click Navigation Rail destination "Favorites"
	if err := tt.Click("Favorites"); err != nil {
		t.Fatalf("failed clicking NavigationRail destination: %v", err)
	}
	if app.railIndex != 1 {
		t.Errorf("expected railIndex == 1, got %d", app.railIndex)
	}

	// 8. Click Dark Mode toggle
	if err := tt.Click("Dark Mode"); err != nil {
		t.Fatalf("failed clicking Dark Mode: %v", err)
	}
	if !app.dark {
		t.Error("expected app.dark to be true after click")
	}

	// 9. Click Density toggle
	if err := tt.Click("Density: Standard"); err != nil {
		t.Fatalf("failed clicking Density toggle: %v", err)
	}
	if app.density != mayu.DensityCompact {
		t.Errorf("expected density to change to Compact, got %v", app.density)
	}

	// 10. Click ListItem "Sarah Connor"
	if err := tt.Click("Sarah Connor"); err != nil {
		t.Fatalf("failed clicking ListItem: %v", err)
	}
	if app.lastAction != "Clicked Sarah Connor ListItem" {
		t.Errorf("expected lastAction updated, got %s", app.lastAction)
	}

	// 11. Open Alert Dialog
	if err := tt.Click("Open Alert Dialog"); err != nil {
		t.Fatalf("failed clicking Open Alert Dialog: %v", err)
	}
	if !app.dialogOpen {
		t.Error("expected dialogOpen == true after click")
	}

	// 12. Confirm dialog
	if err := tt.Click("Reset"); err != nil {
		t.Fatalf("failed clicking Reset in dialog: %v", err)
	}
	if app.dialogOpen {
		t.Error("expected dialogOpen == false after confirm")
	}
}
