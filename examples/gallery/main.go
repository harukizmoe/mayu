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
}

func newGalleryApp() *galleryApp {
	return &galleryApp{
		seedIndex:  0,
		dark:       false,
		density:    mayu.DensityStandard,
		lastAction: "Ready",
	}
}

func (a *galleryApp) view(c *ui.Context) {
	currentSeed := presetSeeds[a.seedIndex].color

	// Apply Mayu M3 Theme
	t := mayu.Use(c, mayu.ThemeOptions{
		SeedColor: currentSeed,
		Dark:      a.dark,
		Density:   a.density,
	})
	scheme := t.Scheme

	ui.Column(c).Fill().Padding(24).Gap(18).Children(func() {
		// Header row
		ui.Row(c).AlignItems(ui.Center).Justify(ui.SpaceBetween).Children(func() {
			ui.Column(c).Gap(4).Children(func() {
				ui.Text(c, "Mayu UI Gallery").FontSize(theme.HeadlineMedium.Size).Bold()
				ui.Text(c, "Material Design 3 Design System for MyGo").TextColor(scheme.OnSurfaceVariant)
			})

			// Controls: Mode & Density
			ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
				// Density toggle
				densityLabel := "Standard"
				if a.density == mayu.DensityCompact {
					densityLabel = "Compact"
				} else if a.density == mayu.DensityComfortable {
					densityLabel = "Comfortable"
				}

				mayu.OutlinedButton(c, "Density: "+densityLabel).
					OnClick(func() {
						if a.density == mayu.DensityStandard {
							a.density = mayu.DensityCompact
						} else if a.density == mayu.DensityCompact {
							a.density = mayu.DensityComfortable
						} else {
							a.density = mayu.DensityStandard
						}
					})

				// Dark/Light toggle
				modeIcon := symbols.DarkMode
				modeLabel := "Dark Mode"
				if a.dark {
					modeIcon = symbols.LightMode
					modeLabel = "Light Mode"
				}
				mayu.FilledTonalButton(c, modeLabel).
					Icon(modeIcon).
					OnClick(func() {
						a.dark = !a.dark
					})
			})
		})

		// Seed Color Selector
		ui.Column(c).Gap(8).Children(func() {
			ui.Text(c, "Dynamic Seed Color").FontSize(theme.TitleMedium.Size).Bold()
			ui.Row(c).Gap(8).Children(func() {
				for i, s := range presetSeeds {
					idx := i
					btn := mayu.OutlinedButton(c, s.name)
					if idx == a.seedIndex {
						btn = mayu.FilledButton(c, s.name)
					}
					btn.OnClick(func() {
						a.seedIndex = idx
						a.lastAction = fmt.Sprintf("Switched seed to %s", s.name)
					})
				}
			})
		})

		// Dynamic Color Palette preview (Surface Container Swatches)
		ui.Column(c).Gap(8).Children(func() {
			ui.Text(c, "Tonal Elevation & Surface Containers").FontSize(theme.TitleSmall.Size).Bold()
			ui.Row(c).Gap(10).Children(func() {
				swatches := []struct {
					name string
					c    ui.Color
					fg   ui.Color
				}{
					{"Primary", scheme.Primary, scheme.OnPrimary},
					{"PrimaryContainer", scheme.PrimaryContainer, scheme.OnPrimaryContainer},
					{"Secondary", scheme.Secondary, scheme.OnSecondary},
					{"SecondaryContainer", scheme.SecondaryContainer, scheme.OnSecondaryContainer},
					{"TertiaryContainer", scheme.TertiaryContainer, scheme.OnTertiaryContainer},
					{"SurfaceContainer", scheme.SurfaceContainer, scheme.OnSurface},
					{"SurfaceHigh", scheme.SurfaceContainerHigh, scheme.OnSurface},
				}

				for _, sw := range swatches {
					ui.Box(c).Width(100).Height(50).Radius(10).Background(sw.c).Center().Children(func() {
						ui.Text(c, sw.name).FontSize(10).TextColor(sw.fg).SingleLine()
					})
				}
			})
		})

		// Component Showcase: 5 M3 Button Variants
		ui.Column(c).Gap(12).Children(func() {
			ui.Text(c, "M3 Button Family (5 Variants)").FontSize(theme.TitleMedium.Size).Bold()

			// Row 1: Active buttons
			ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
				mayu.FilledButton(c, "Filled Button").
					Icon(symbols.Send).
					OnClick(func() {
						a.clicks++
						a.lastAction = "Clicked Filled Button"
					})

				mayu.ElevatedButton(c, "Elevated Button").
					Icon(symbols.Add).
					OnClick(func() {
						a.clicks++
						a.lastAction = "Clicked Elevated Button"
					})

				mayu.FilledTonalButton(c, "Filled Tonal").
					Icon(symbols.Favorite).
					OnClick(func() {
						a.clicks++
						a.lastAction = "Clicked Filled Tonal Button"
					})

				mayu.OutlinedButton(c, "Outlined").
					Icon(symbols.Settings).
					OnClick(func() {
						a.clicks++
						a.lastAction = "Clicked Outlined Button"
					})

				mayu.TextButton(c, "Text Button").
					Icon(symbols.Close).
					OnClick(func() {
						a.clicks++
						a.lastAction = "Clicked Text Button"
					})
			})

			// Row 2: Disabled buttons to show M3 disabled state layers
			ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
				mayu.FilledButton(c, "Filled Disabled").
					Disabled(true)

				mayu.ElevatedButton(c, "Elevated Disabled").
					Disabled(true)

				mayu.FilledTonalButton(c, "Tonal Disabled").
					Disabled(true)

				mayu.OutlinedButton(c, "Outlined Disabled").
					Disabled(true)

				mayu.TextButton(c, "Text Disabled").
					Disabled(true)
			})
		})

		// Interaction Status Footer
		ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
			ui.Text(c, fmt.Sprintf("Clicks: %d", a.clicks)).Bold()
			ui.Text(c, "•").TextColor(scheme.Outline)
			ui.Text(c, fmt.Sprintf("Status: %s", a.lastAction)).TextColor(scheme.OnSurfaceVariant)
		})
	})
}

func main() {
	app := newGalleryApp()
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Mayu — Material 3 UI for MyGo",
			Width:   880,
			Height:  640,
			Content: ui.View(app.view),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
