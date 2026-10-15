package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
)

func TestGallery_InteractiveFeatures(t *testing.T) {
	app := newGalleryApp()

	tt := ui.NewTester(app.view, 900, 650)

	// Verify initial elements
	if !tt.HasText("Mayu UI Gallery") {
		t.Fatal("gallery header not found")
	}
	if !tt.HasText("Filled Button") {
		t.Fatal("Filled Button not found")
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
}
