// Command gallery is the interactive showcase of Mayu: every component of the
// library, on one of three pages the navigation rail switches between.
//
//	go run ./examples/gallery
package main

import (
	"fmt"
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// navPage is one destination of the navigation rail: a page of the gallery.
type navPage struct {
	title string
	icon  *ui.SVG
	badge int
}

// navPages are the rail's destinations, in order. The index selected is the
// page shown, so adding a page to this list and a case to view's switch is all
// it takes.
var navPages = []navPage{
	{title: "Components", icon: symbols.Settings, badge: 3},
	{title: "Overlays", icon: symbols.Info},
	{title: "Theme", icon: symbols.Palette},
}

type seedOption struct {
	name  string
	color ui.Color
}

var presetSeeds = []seedOption{
	{"M3 Purple", ui.Hex("#6750A4")},
	{"Coral Red", ui.Hex("#B3261E")},
	{"Ocean Blue", ui.Hex("#006492")},
	{"Forest Green", ui.Hex("#006D44")},
	{"Amber Gold", ui.Hex("#795900")},
}

type galleryApp struct {
	seedIndex  int
	dark       bool
	density    mayu.Density
	clicks     int
	lastAction string

	// page is the rail's destination: the page of the gallery shown.
	page int

	// segIdx and refreshIdx are the choices of two segmented button groups,
	// each its own.
	segIdx     int
	refreshIdx int

	searchQuery string

	// Selection state
	checkboxVal bool
	switchVal   bool
	radioChoice string

	// Text input state
	username string
	email    string
	notes    string
	city     string

	// Choosing state
	tabIdx int
	zone   string

	// Slider & progress state
	sliderVal      float64
	priceLow       float64
	priceHigh      float64
	brightnessVal  float64
	volumeVal      float64
	lightSwitchVal float64

	// Overlay states
	accordionOpen bool
	menuOpen      bool
	dialogOpen    bool
	snackbarOpen  bool
	snackbarMsg   string
	sideSheetOpen bool
	// toastShape and toastColor pick the host's toast outline and colors.
	toastShape int
	toastColor int
}

func newGalleryApp() *galleryApp {
	return &galleryApp{
		seedIndex:      0,
		dark:           false,
		density:        mayu.DensityStandard,
		lastAction:     "Ready",
		page:           0,
		searchQuery:    "",
		checkboxVal:    true,
		switchVal:      true,
		radioChoice:    "Standard",
		username:       "Alex",
		email:          "alex@example.com",
		notes:          "A new component batch: tabs, select, autocomplete and a progress ring.",
		city:           "",
		tabIdx:         0,
		zone:           "UTC+8",
		sliderVal:      60.0,
		priceLow:       30.0,
		priceHigh:      75.0,
		brightnessVal:  65.0,
		volumeVal:      35.0,
		lightSwitchVal: 80.0,
		accordionOpen:  true,
		snackbarMsg:    "Settings saved successfully",
	}
}

func (a *galleryApp) pageTitle() string {
	if a.page < 0 || a.page >= len(navPages) {
		return navPages[0].title
	}
	return navPages[a.page].title
}

func (a *galleryApp) view(c *ui.Context) {
	// 1. Initialize M3 Theme from the chosen seed.
	t := mayu.Use(c, mayu.ThemeOptions{
		SeedColor: presetSeeds[a.seedIndex].color,
		Dark:      a.dark,
		Density:   a.density,
	})

	// The toast host covers the window; the rail cannot, so it goes first.
	a.toastHost(c, t)

	ui.Row(c).Fill().Children(func() {
		a.rail(c)

		// The page keeps its own scroll offset: switching pages starts at the
		// top of the new one.
		ui.Scroll(c).Grow(1).FillHeight().Key(a.page).Children(func() {
			ui.Column(c).FillWidth().Padding(24).Gap(20).Children(func() {
				a.header(c, t)

				switch a.page {
				case 0:
					a.componentsPage(c, t)
				case 1:
					a.overlaysPage(c, t)
				default:
					a.themePage(c, t)
				}

				a.statusFooter(c, t)
			})
		})
	})

	a.modals(c, t)
}

// rail is the desktop skeleton: the navigation rail whose destinations are the
// pages, with a FAB over them and the user at the foot.
func (a *galleryApp) rail(c *ui.Context) {
	destinations := make([]mayu.NavDestination, len(navPages))
	for i, p := range navPages {
		destinations[i] = mayu.NavDestination{Icon: p.icon, Label: p.title, Badge: p.badge}
	}

	mayu.NavigationRail(c, mayu.NavigationRailOptions{
		Selected:     &a.page,
		Destinations: destinations,
		Header: func() {
			mayu.FAB(c, symbols.Add).OnClick(func() {
				a.clicks++
				a.lastAction = "Clicked the rail's FAB"
			})
		},
		Footer: func() {
			mayu.InitialsAvatar(c, "Mayu User")
		},
		OnSelect: func(idx int) {
			a.lastAction = fmt.Sprintf("Switched to the %s page", navPages[idx].title)
		},
	})
}

// modals are the surfaces that belong to the window rather than to a page, so
// they open over whichever page is shown.
func (a *galleryApp) modals(c *ui.Context, t *theme.Theme) {
	mayu.AlertDialog(c, &a.dialogOpen, mayu.AlertDialogOptions{
		Title:       "Confirm Reset",
		Body:        "Do you want to reset all gallery selections to default?",
		ConfirmText: "Reset",
		OnConfirm: func() {
			a.clicks = 0
			a.lastAction = "Reset All Settings"
		},
		DismissText: "Cancel",
	})

	mayu.SideSheet(c, &a.sideSheetOpen, mayu.SideSheetOptions{
		Title:    "显示与亮度",
		Subtitle: "DP-1（主显示器）",
		Actions: func() {
			ui.Row(c).Justify(ui.End).Gap(8).Children(func() {
				mayu.TextButton(c, "取消").
					OnClick(func() { a.sideSheetOpen = false })
				mayu.FilledButton(c, "应用").
					OnClick(func() {
						a.sideSheetOpen = false
						a.lastAction = "应用了显示设置"
						mayu.ShowToast(c, mayu.ToastOptions{
							Title: "显示设置已应用",
							Kind:  mayu.ToastSuccess,
						})
					})
			})
		},
	}, func() {
		mayu.LabeledSlider(c, &a.brightnessVal, 0, 100, mayu.LabeledSliderOptions{
			Label: "亮度 (Brightness)",
			Icon:  symbols.Lightbulb,
		})

		mayu.LabeledSlider(c, &a.volumeVal, 0, 100, mayu.LabeledSliderOptions{
			Label: "音量 (Volume)",
			Icon:  symbols.VolumeUp,
		})

		ui.Column(c).Gap(8).Children(func() {
			ui.Text(c, "显示选项").FontSize(theme.TitleSmall.Size).Bold().TextColor(t.Scheme.OnSurface)
			mayu.Switch(c, &a.switchVal, "夜间模式")
			mayu.Checkbox(c, &a.checkboxVal, "硬件加速合成")
		})

		mayu.SegmentedButton(c, &a.refreshIdx, []mayu.SegmentOption{
			{Label: "60 Hz"},
			{Label: "120 Hz"},
			{Label: "180 Hz"},
		})

		mayu.OutlinedTextField(c, &a.username, "配置文件名").
			Placeholder("my-profile")
	})
}

func main() {
	app := newGalleryApp()
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Mayu — Material 3 UI for MyGo",
			Width:   1080,
			Height:  780,
			Content: ui.View(app.view),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
