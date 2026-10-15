package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
)

// The navigation rail switches pages: each destination shows its own page and
// no other, so the gallery is three browsable pages rather than one long one.
func TestGallery_Navigation(t *testing.T) {
	app := newGalleryApp()
	tt := ui.NewTester(app.view, 1400, 1800)

	onPage := func(title string, in, out []string) {
		t.Helper()
		if !tt.HasText(title) {
			t.Fatalf("the %s page is not the one shown", title)
		}
		for _, s := range in {
			if !tt.HasText(s) {
				t.Errorf("%s: the page lacks %q", title, s)
			}
		}
		for _, s := range out {
			if tt.HasText(s) {
				t.Errorf("%s: the page still shows %q from another page", title, s)
			}
		}
	}

	components := []string{"Buttons & Floating Actions", "Cards", "Selection Controls", "Accordion"}
	overlays := []string{"Tooltips & Dividers", "Dropdown Menu", "Toast & Snackbar", "Modal Surfaces"}
	themePage := []string{"Dynamic Seed Color", "Tonal Elevation", "Typography & Shape", "Capsule Sliders"}

	onPage("Components", components, append(overlays, themePage...))

	if err := tt.Click("Overlays"); err != nil {
		t.Fatalf("failed clicking the rail's Overlays destination: %v", err)
	}
	if app.page != 1 {
		t.Fatalf("expected the rail to select page 1, got %d", app.page)
	}
	onPage("Overlays", overlays, append(components, themePage...))

	if err := tt.Click("Theme"); err != nil {
		t.Fatalf("failed clicking the rail's Theme destination: %v", err)
	}
	if app.page != 2 {
		t.Fatalf("expected the rail to select page 2, got %d", app.page)
	}
	onPage("Theme", themePage, append(components, overlays...))

	// And back, which is where the page's own scroll offset is kept.
	if err := tt.Click("Components"); err != nil {
		t.Fatalf("failed clicking the rail's Components destination: %v", err)
	}
	if app.page != 0 {
		t.Fatalf("expected the rail to select page 0, got %d", app.page)
	}
	onPage("Components", components, append(overlays, themePage...))
}

// The search box filters the sections of the page being shown, so it does what
// its placeholder says rather than sitting there.
func TestGallery_SearchFiltersThePage(t *testing.T) {
	app := newGalleryApp()
	tt := ui.NewTester(app.view, 1400, 1800)

	// Typing into the box finds it by the name it gives assistive technology.
	if err := tt.Click("Filter this page's components..."); err != nil {
		t.Fatalf("the search box is not reachable: %v", err)
	}
	tt.Type("Chips")
	tt.Frame()

	if app.searchQuery != "Chips" {
		t.Fatalf("the search box did not take the text, query=%q", app.searchQuery)
	}
	if !tt.HasText("Chips") {
		t.Fatal("the matching section is not shown")
	}
	for _, gone := range []string{"Buttons & Floating Actions", "Cards", "Accordion"} {
		if tt.HasText(gone) {
			t.Errorf("a search for Chips still shows %q", gone)
		}
	}

	// A query nothing matches says so instead of leaving the page blank.
	app.searchQuery = "nothing matches this"
	tt.Frame()
	if !tt.HasText("No component matches") {
		t.Fatalf("the empty state is missing, texts: %q", tt.Texts())
	}

	// Clearing it brings the page back whole.
	app.searchQuery = ""
	tt.Frame()
	for _, want := range []string{"Buttons & Floating Actions", "Cards", "Accordion"} {
		if !tt.HasText(want) {
			t.Errorf("clearing the search did not restore %q", want)
		}
	}
}

// The Components page's widgets act on the state the view reads.
func TestGallery_ComponentsPage(t *testing.T) {
	app := newGalleryApp()
	tt := ui.NewTester(app.view, 1400, 1800)

	if err := tt.Click("Filled Button"); err != nil {
		t.Fatalf("failed clicking the filled button: %v", err)
	}
	if app.clicks != 1 || app.lastAction != "Clicked the Filled Button" {
		t.Fatalf("the filled button did not act: %d clicks, %q", app.clicks, app.lastAction)
	}

	if err := tt.Click("Compose"); err != nil {
		t.Fatalf("failed clicking the extended FAB: %v", err)
	}
	if app.lastAction != "Clicked the extended FAB" {
		t.Fatalf("the extended FAB did not act: %q", app.lastAction)
	}

	if err := tt.Click("Elevated Card"); err != nil {
		t.Fatalf("failed clicking the elevated card: %v", err)
	}
	if app.lastAction != "Clicked an elevated card" {
		t.Fatalf("the card did not act: %q", app.lastAction)
	}

	before := app.checkboxVal
	if err := tt.Click("Enable notifications"); err != nil {
		t.Fatalf("failed clicking the checkbox: %v", err)
	}
	if app.checkboxVal == before {
		t.Error("the checkbox did not toggle")
	}

	if err := tt.Click("Action Chip"); err != nil {
		t.Fatalf("failed clicking the action chip: %v", err)
	}
	if app.lastAction != "Clicked an action chip" {
		t.Fatalf("the chip did not act: %q", app.lastAction)
	}

	if err := tt.Click("Sarah Connor"); err != nil {
		t.Fatalf("failed clicking the list item: %v", err)
	}
	if app.lastAction != "Clicked the Sarah Connor row" {
		t.Fatalf("the list item did not act: %q", app.lastAction)
	}

	// The accordion's own content is hidden until it is opened.
	if !tt.HasText("Advanced System Configuration") {
		t.Fatal("the accordion header is missing")
	}
}

// The Overlays page raises the surfaces, in the shapes and colours it lets you
// pick.
func TestGallery_OverlaysPage(t *testing.T) {
	app := newGalleryApp()
	app.page = 1
	tt := ui.NewTester(app.view, 1400, 1800)

	// A toast, in the capsule the host starts with.
	if err := tt.Click("Toast: Success"); err != nil {
		t.Fatalf("failed clicking the toast trigger: %v", err)
	}
	tt.Frame()
	if !tt.HasText("已保存") {
		t.Fatalf("the toast does not show, texts: %q", tt.Texts())
	}

	// Switching the host's shape and colour applies to the next toast.
	if err := tt.Click("Corner"); err != nil {
		t.Fatalf("failed clicking the Corner segment: %v", err)
	}
	if err := tt.Click("Inverse"); err != nil {
		t.Fatalf("failed clicking the Inverse segment: %v", err)
	}
	if app.toastShape != 1 || app.toastColor != 1 {
		t.Fatalf("the toast options did not change: shape=%d color=%d", app.toastShape, app.toastColor)
	}

	// Closing every toast clears them.
	if err := tt.Click("Close all toasts"); err != nil {
		t.Fatalf("failed clicking Close all toasts: %v", err)
	}
	tt.Frame()
	if tt.HasText("已保存") {
		t.Errorf("the toasts are still there: %q", tt.Texts())
	}

	// A tooltip appears once its anchor has been hovered for its delay.
	tip, ok := tt.Find("悬停查看提示")
	if !ok {
		t.Fatal("the chip with a tooltip is not laid out")
	}
	tt.Move(tip.X+tip.W/2, tip.Y+tip.H/2)
	time.Sleep(700 * time.Millisecond) // the tooltip's 600ms delay
	tt.Frame()
	if !tt.HasText("这是 M3 Tooltip") {
		t.Fatalf("the tooltip does not show, texts: %q", tt.Texts())
	}

	// The side sheet opens from the page and closes from its own footer.
	if err := tt.Click("Open Side Sheet"); err != nil {
		t.Fatalf("failed clicking the side sheet trigger: %v", err)
	}
	if !app.sideSheetOpen {
		t.Fatal("the side sheet did not open")
	}
	tt.Frame()
	if !tt.HasText("显示与亮度") {
		t.Fatalf("the side sheet does not show, texts: %q", tt.Texts())
	}
	if err := tt.Click("应用"); err != nil {
		t.Fatalf("failed clicking the sheet's footer action: %v", err)
	}
	if app.sideSheetOpen {
		t.Error("the sheet's action did not close it")
	}
	tt.Frame()
	if !tt.HasText("显示设置已应用") {
		t.Fatalf("the sheet did not raise its toast, texts: %q", tt.Texts())
	}
}

// The Theme page restyles the whole window and shows the tokens it is built
// from.
func TestGallery_ThemePage(t *testing.T) {
	app := newGalleryApp()
	app.page = 2
	tt := ui.NewTester(app.view, 1400, 1800)

	// Every seed is drawn once, and the chosen one still acts.
	if err := tt.Click("Ocean Blue"); err != nil {
		t.Fatalf("failed clicking a seed: %v", err)
	}
	if app.seedIndex != 2 || app.clicks != 1 {
		t.Fatalf("the seed did not change: seed=%d clicks=%d", app.seedIndex, app.clicks)
	}
	if err := tt.Click("Ocean Blue"); err != nil {
		t.Fatalf("failed clicking the chosen seed: %v", err)
	}
	if app.clicks != 2 {
		t.Fatalf("the chosen seed's button does not act: %d clicks", app.clicks)
	}

	// The elevation swatches and the type scale are shown.
	for _, want := range []string{"SurfaceContainerLowest", "SurfaceContainerHighest", "HeadlineMedium"} {
		if !tt.HasText(want) {
			t.Errorf("the Theme page lacks %q", want)
		}
	}

	// The header's switches restyle the window.
	if err := tt.Click("Dark Mode"); err != nil {
		t.Fatalf("failed clicking the mode switch: %v", err)
	}
	if !app.dark {
		t.Error("the mode switch did not turn dark on")
	}
	if err := tt.Click("Density: Standard"); err != nil {
		t.Fatalf("failed clicking the density switch: %v", err)
	}
	if app.density != mayu.DensityCompact {
		t.Errorf("the density switch did not change, got %v", app.density)
	}

	// The segmented group is its own: the one in the sheet must not share it.
	if err := tt.Click("Weekly"); err != nil {
		t.Fatalf("failed clicking a segment: %v", err)
	}
	if app.segIdx != 1 {
		t.Errorf("the segmented group did not change, got %d", app.segIdx)
	}
	if app.refreshIdx != 0 {
		t.Errorf("the sheet's own segmented group moved instead, got %d", app.refreshIdx)
	}
}
