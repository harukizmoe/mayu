package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
)

func TestGallery_InteractiveFeatures(t *testing.T) {
	app := newGalleryApp()

	tt := ui.NewTester(app.view, 1400, 2900)

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

	// 3b. Verify the labeled slider card and its value badge
	if !tt.HasText("亮度 (Brightness)") {
		t.Fatal("labeled slider card not found")
	}
	if !tt.HasText("65 %") {
		t.Fatalf("labeled slider value badge not found, texts: %q", tt.Texts())
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

	// 13. Toast: clicking a trigger shows it in the host on the next frame.
	if err := tt.Click("Toast: Success"); err != nil {
		t.Fatalf("failed clicking the toast trigger: %v", err)
	}
	tt.Frame()
	if !tt.HasText("已保存") {
		t.Fatalf("the toast does not show, texts: %q", tt.Texts())
	}
}

func TestGallery_ToastShapes(t *testing.T) {
	app := newGalleryApp()
	tt := ui.NewTester(app.view, 1400, 2900)

	// The capsule shape is the host's default.
	if app.toastShapeIdx != 0 {
		t.Fatalf("expected the capsule shape by default, got %d", app.toastShapeIdx)
	}

	// Switch the host to the rounded rectangle.
	if err := tt.Click("Corner 圆角矩形"); err != nil {
		t.Fatalf("failed clicking the toast shape segment: %v", err)
	}
	if app.toastShapeIdx != 1 {
		t.Errorf("expected the shape to switch to corner, got %d", app.toastShapeIdx)
	}

	tt.Frame()
	if err := tt.Click("Toast: Undo"); err != nil {
		t.Fatalf("failed clicking the toast trigger: %v", err)
	}
	tt.Frame()
	if !tt.HasText("已将 3 个项目移至废纸篓") {
		t.Fatalf("the corner-shaped toast does not show, texts: %q", tt.Texts())
	}
	if !tt.HasText("撤销") {
		t.Fatalf("the corner-shaped toast lacks its action, texts: %q", tt.Texts())
	}

	// Switch the host to the inverse colors, the classic snackbar.
	if err := tt.Click("Inverse 反色"); err != nil {
		t.Fatalf("failed clicking the toast color segment: %v", err)
	}
	if app.toastColorIdx != 1 {
		t.Errorf("expected the colors to switch to inverse, got %d", app.toastColorIdx)
	}

	tt.Frame()
	if err := tt.Click("Toast: Info"); err != nil {
		t.Fatalf("failed clicking the toast trigger: %v", err)
	}
	tt.Frame()
	if !tt.HasText("Material You 已启用") {
		t.Fatalf("the inverse-colored toast does not show, texts: %q", tt.Texts())
	}
	// Every component of the library must be reachable from the gallery.
	for _, want := range []string{
		"Tooltips, Dividers & Icon Buttons",
		"悬停查看提示",
		"关闭全部 Toast",
		"Modal Surfaces",
		"Open Side Sheet",
		"Advanced System Configuration",
	} {
		if !tt.HasText(want) {
			t.Errorf("the gallery does not show %q", want)
		}
	}

	// A tooltip appears once its anchor has been hovered for its delay.
	move, ok := tt.Find("悬停查看提示")
	if !ok {
		t.Fatal("the chip with a tooltip is not laid out")
	}
	tt.Move(move.X+move.W/2, move.Y+move.H/2)
	time.Sleep(700 * time.Millisecond) // the tooltip's 600ms delay
	tt.Frame()
	if !tt.HasText("这是 M3 Tooltip") {
		t.Fatalf("the tooltip does not show, texts: %q", tt.Texts())
	}
}

func TestGallery_SeedPicker(t *testing.T) {
	app := newGalleryApp()
	tt := ui.NewTester(app.view, 1400, 2900)

	// Choosing a seed recolors the window with it.
	if err := tt.Click("Ocean Blue"); err != nil {
		t.Fatalf("failed clicking a seed: %v", err)
	}
	if app.seedIndex != 2 || app.clicks != 1 {
		t.Fatalf("expected the Ocean Blue seed from one click, got seed %d after %d clicks",
			app.seedIndex, app.clicks)
	}

	// The selected seed's button must still be the one that acts: building an
	// outlined button and replacing it with a filled one draws the selected
	// seed twice, with the inert copy first, so this click would do nothing.
	if err := tt.Click("Ocean Blue"); err != nil {
		t.Fatalf("failed clicking the selected seed: %v", err)
	}
	if app.clicks != 2 {
		t.Fatalf("the selected seed's button does not act: %d clicks", app.clicks)
	}

	// And switching away and back keeps working.
	if err := tt.Click("M3 Purple"); err != nil {
		t.Fatalf("failed clicking M3 Purple: %v", err)
	}
	if app.seedIndex != 0 || app.clicks != 3 {
		t.Fatalf("expected the M3 Purple seed from a third click, got seed %d after %d clicks",
			app.seedIndex, app.clicks)
	}
}

func TestGallery_SideSheet(t *testing.T) {
	app := newGalleryApp()
	tt := ui.NewTester(app.view, 1400, 2900)

	// A closed side sheet contributes nothing.
	if tt.HasText("显示与亮度") {
		t.Fatal("the side sheet renders while closed")
	}
	if tt.HasText("DP-1（主显示器）") {
		t.Fatal("the side sheet's subtitle renders while closed")
	}

	if err := tt.Click("Open Side Sheet"); err != nil {
		t.Fatalf("failed clicking the side sheet trigger: %v", err)
	}
	if !app.sideSheetOpen {
		t.Fatal("expected the side sheet to open")
	}

	tt.Frame()
	for _, want := range []string{"显示与亮度", "DP-1（主显示器）", "亮度 (Brightness)", "显示选项", "夜间模式", "配置文件名"} {
		if !tt.HasText(want) {
			t.Fatalf("the side sheet lacks %q: %q", want, tt.Texts())
		}
	}

	// The footer action closes the sheet and raises a toast.
	if err := tt.Click("应用"); err != nil {
		t.Fatalf("failed clicking the side sheet's footer action: %v", err)
	}
	if app.sideSheetOpen {
		t.Error("expected 应用 to close the side sheet")
	}
	tt.Frame()
	if !tt.HasText("显示设置已应用") {
		t.Fatalf("no toast after applying, texts: %q", tt.Texts())
	}
}
