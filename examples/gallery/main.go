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

	// Selection state
	checkboxVal bool
	switchVal   bool
	radioChoice string

	// Text input state
	username string
	email    string

	// Slider & progress state
	sliderVal float64

	// Dialog state
	dialogOpen bool
}

func newGalleryApp() *galleryApp {
	return &galleryApp{
		seedIndex:   0,
		dark:        false,
		density:     mayu.DensityStandard,
		lastAction:  "Ready",
		checkboxVal: true,
		switchVal:   true,
		radioChoice: "Standard",
		username:    "Alex",
		email:       "alex@example.com",
		sliderVal:   60.0,
		dialogOpen:  false,
	}
}

func (a *galleryApp) view(c *ui.Context) {
	currentSeed := presetSeeds[a.seedIndex].color

	// 1. Initialize M3 Theme
	t := mayu.Use(c, mayu.ThemeOptions{
		SeedColor: currentSeed,
		Dark:      a.dark,
		Density:   a.density,
	})
	scheme := t.Scheme

	// Outer scrollable container
	ui.Scroll(c).Fill().Children(func() {
		ui.Column(c).FillWidth().Padding(24).Gap(20).Children(func() {
			// Header row
			ui.Row(c).AlignItems(ui.Center).Justify(ui.SpaceBetween).Children(func() {
				ui.Column(c).Gap(4).Children(func() {
					ui.Row(c).Gap(10).AlignItems(ui.Center).Children(func() {
						ui.Text(c, "Mayu Component Gallery").FontSize(theme.HeadlineMedium.Size).Bold()
						mayu.Badge(c, 12)
					})
					ui.Text(c, "Material Design 3 (Material You) UI Toolkit for MyGo").TextColor(scheme.OnSurfaceVariant)
				})

				// Mode & Density controls
				ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
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

			// 1. Seed Color Row
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

			mayu.HorizontalDivider(c)

			// 2. Buttons & FAB Section
			ui.Column(c).Gap(12).Children(func() {
				ui.Text(c, "Buttons & Floating Actions").FontSize(theme.TitleMedium.Size).Bold()

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
						OnClick(func() {
							a.clicks++
							a.lastAction = "Clicked Text Button"
						})
				})

				// FABs & IconButtons Row
				ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
					mayu.ExtendedFAB(c, symbols.Add, "Compose").
						OnClick(func() {
							a.clicks++
							a.lastAction = "Clicked Extended FAB"
						})

					mayu.FAB(c, symbols.Favorite).
						OnClick(func() {
							a.clicks++
							a.lastAction = "Clicked FAB"
						})

					mayu.IconButton(c, symbols.Search).
						OnClick(func() {
							a.lastAction = "Clicked IconButton"
						})

					mayu.FilledIconButton(c, symbols.Settings).
						OnClick(func() {
							a.lastAction = "Clicked FilledIconButton"
						})

					mayu.FilledTonalIconButton(c, symbols.Send).
						OnClick(func() {
							a.lastAction = "Clicked FilledTonalIconButton"
						})

					mayu.OutlinedIconButton(c, symbols.Close).
						OnClick(func() {
							a.lastAction = "Clicked OutlinedIconButton"
						})
				})
			})

			mayu.HorizontalDivider(c)

			// 3. Selection Controls Section
			ui.Column(c).Gap(12).Children(func() {
				ui.Text(c, "Selection Controls").FontSize(theme.TitleMedium.Size).Bold()

				ui.Row(c).Gap(24).AlignItems(ui.Center).Children(func() {
					mayu.Checkbox(c, &a.checkboxVal, "Enable notifications")

					mayu.Switch(c, &a.switchVal, "Airplane mode")

					ui.Row(c).Gap(16).Children(func() {
						mayu.Radio(c, &a.radioChoice, "Standard", "Standard Tier")
						mayu.Radio(c, &a.radioChoice, "Pro", "Pro Tier")
					})
				})
			})

			mayu.HorizontalDivider(c)

			// 4. Cards Section
			ui.Column(c).Gap(12).Children(func() {
				ui.Text(c, "Cards (Tonal Elevation)").FontSize(theme.TitleMedium.Size).Bold()

				ui.Row(c).Gap(16).Children(func() {
					mayu.ElevatedCard(c).Width(240).Children(func() {
						ui.Column(c).Gap(6).Children(func() {
							ui.Text(c, "Elevated Card").FontSize(theme.TitleMedium.Size).Bold()
							ui.Text(c, "SurfaceContainerLow with resting elevation shadow.").TextColor(scheme.OnSurfaceVariant).FontSize(13)
						})
					}).OnClick(func() {
						a.lastAction = "Clicked Elevated Card"
					})

					mayu.FilledCard(c).Width(240).Children(func() {
						ui.Column(c).Gap(6).Children(func() {
							ui.Text(c, "Filled Card").FontSize(theme.TitleMedium.Size).Bold()
							ui.Text(c, "SurfaceContainerHighest without elevation shadow.").TextColor(scheme.OnSurfaceVariant).FontSize(13)
						})
					}).OnClick(func() {
						a.lastAction = "Clicked Filled Card"
					})

					mayu.OutlinedCard(c).Width(240).Children(func() {
						ui.Column(c).Gap(6).Children(func() {
							ui.Text(c, "Outlined Card").FontSize(theme.TitleMedium.Size).Bold()
							ui.Text(c, "Surface with 1dp OutlineVariant clean border.").TextColor(scheme.OnSurfaceVariant).FontSize(13)
						})
					}).OnClick(func() {
						a.lastAction = "Clicked Outlined Card"
					})
				})
			})

			mayu.HorizontalDivider(c)

			// 5. Text Fields Section
			ui.Column(c).Gap(12).Children(func() {
				ui.Text(c, "Text Fields").FontSize(theme.TitleMedium.Size).Bold()

				ui.Row(c).Gap(20).FillWidth().Children(func() {
					mayu.FilledTextField(c, &a.username, "Username").
						Placeholder("Enter username")

					mayu.OutlinedTextField(c, &a.email, "Email Address").
						Placeholder("Enter email")
				})
			})

			mayu.HorizontalDivider(c)

			// 6. Chips, Slider & Progress Section
			ui.Column(c).Gap(12).Children(func() {
				ui.Text(c, "Chips, Sliders & Progress").FontSize(theme.TitleMedium.Size).Bold()

				// Chips row
				ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
					mayu.Chip(c, "Action Chip").
						Icon(symbols.Settings).
						OnClick(func() { a.lastAction = "Clicked Action Chip" })

					mayu.FilterChip(c, "Filter (Active)", true).
						OnClick(func() { a.lastAction = "Toggled Filter Chip" })

					mayu.FilterChip(c, "Filter (Inactive)", false).
						OnClick(func() { a.lastAction = "Toggled Inactive Filter Chip" })
				})

				// Slider & Progress
				ui.Column(c).Gap(8).Children(func() {
					ui.Row(c).AlignItems(ui.Center).Justify(ui.SpaceBetween).Children(func() {
						ui.Text(c, fmt.Sprintf("Volume Level: %.0f%%", a.sliderVal))
						mayu.Button(c, "Open Alert Dialog", mayu.VariantFilledTonal).
							OnClick(func() {
								a.dialogOpen = true
							})
					})

					mayu.Slider(c, &a.sliderVal, 0, 100)
					mayu.LinearProgress(c, float32(a.sliderVal/100.0))
				})
			})

			// Footer Status Row
			ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
				ui.Text(c, fmt.Sprintf("Clicks: %d", a.clicks)).Bold()
				ui.Text(c, "•").TextColor(scheme.Outline)
				ui.Text(c, fmt.Sprintf("Selection: Checkbox=%v, Switch=%v, Tier=%s", a.checkboxVal, a.switchVal, a.radioChoice)).TextColor(scheme.OnSurfaceVariant)
				ui.Text(c, "•").TextColor(scheme.Outline)
				ui.Text(c, fmt.Sprintf("Status: %s", a.lastAction)).TextColor(scheme.Primary)
			})
		})
	})

	// Modal Dialog when dialogOpen is true
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
}

func main() {
	app := newGalleryApp()
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Mayu — Material 3 UI for MyGo",
			Width:   960,
			Height:  720,
			Content: ui.View(app.view),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
