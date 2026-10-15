package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
)

func TestGallery_InteractiveFeatures(t *testing.T) {
	app := newGalleryApp()

	tt := ui.NewTester(app.view, 1200, 1600)
	if !tt.HasText("Mayu Component Gallery") {
		t.Fatal("gallery header not found")
	}
	if !tt.HasText("Filled Button") {
		t.Fatal("Filled Button not found")
	}
	if !tt.HasText("Elevated Card") {
		t.Fatal("Elevated Card not found")
	}
	if !tt.HasText("Enable notifications") {
		t.Fatal("Checkbox label not found")
	}
	if !tt.HasText("Airplane mode") {
		t.Fatal("Switch label not found")
	}

	// 1. Click "Filled Button" and verify clicks incremented
	if err := tt.Click("Filled Button"); err != nil {
		t.Fatalf("failed clicking Filled Button: %v", err)
	}
	if app.clicks != 1 {
		t.Errorf("expected clicks == 1, got %d", app.clicks)
	}

	// 2. Click "Ocean Blue" seed button
	if err := tt.Click("Ocean Blue"); err != nil {
		t.Fatalf("failed clicking Ocean Blue seed: %v", err)
	}
	if app.seedIndex != 2 {
		t.Errorf("expected seedIndex == 2, got %d", app.seedIndex)
	}

	// 3. Click Dark Mode toggle
	if err := tt.Click("Dark Mode"); err != nil {
		t.Fatalf("failed clicking Dark Mode: %v", err)
	}
	if !app.dark {
		t.Error("expected app.dark to be true after click")
	}

	// 4. Click Density toggle
	if err := tt.Click("Density: Standard"); err != nil {
		t.Fatalf("failed clicking Density toggle: %v", err)
	}
	if app.density != mayu.DensityCompact {
		t.Errorf("expected density to change to Compact, got %v", app.density)
	}

	// 5. Click Checkbox to toggle
	if err := tt.Click("Enable notifications"); err != nil {
		t.Fatalf("failed clicking Checkbox: %v", err)
	}
	if app.checkboxVal {
		t.Error("expected checkboxVal to toggle to false")
	}

	// 6. Click Elevated Card
	if err := tt.Click("Elevated Card"); err != nil {
		t.Fatalf("failed clicking Elevated Card: %v", err)
	}
	if app.lastAction != "Clicked Elevated Card" {
		t.Errorf("expected lastAction == 'Clicked Elevated Card', got %s", app.lastAction)
	}

	// 7. Click Extended FAB
	if err := tt.Click("Compose"); err != nil {
		t.Fatalf("failed clicking Extended FAB: %v", err)
	}
	if app.lastAction != "Clicked Extended FAB" {
		t.Errorf("expected lastAction == 'Clicked Extended FAB', got %s", app.lastAction)
	}

	// 8. Open Alert Dialog
	if err := tt.Click("Open Alert Dialog"); err != nil {
		t.Fatalf("failed clicking Open Alert Dialog: %v", err)
	}
	if !app.dialogOpen {
		t.Error("expected dialogOpen == true after click")
	}

	// 9. Confirm dialog
	if err := tt.Click("Reset"); err != nil {
		t.Fatalf("failed clicking Reset in dialog: %v", err)
	}
	if app.clicks != 0 {
		t.Errorf("expected clicks reset to 0, got %d", app.clicks)
	}
	if app.dialogOpen {
		t.Error("expected dialogOpen == false after confirm")
	}
}
