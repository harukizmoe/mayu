package main

import (
	"fmt"
	"log"
	"time"

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

	// Navigation & segmented state
	railIndex      int
	segmentedIndex int
	searchQuery    string

	// Selection state
	checkboxVal bool
	switchVal   bool
	radioChoice string

	// Text input state
	username string
	email    string

	// Slider & progress state
	sliderVal      float64
	brightnessVal  float64
	volumeVal      float64
	lightSwitchVal float64

	// Overlay states
	accordionOpen bool
	menuOpen      bool
	dialogOpen    bool
	snackbarOpen  bool
	snackbarMsg   string

	// toastShapeIdx picks the host's toast outline: 0 capsule, 1 corner.
	toastShapeIdx int
	// toastColorIdx picks the host's toast colors: 0 surface, 1 inverse.
	toastColorIdx int
	sideSheetOpen bool
	// refreshIdx is the side sheet's refresh-rate choice, its own group.
	refreshIdx int
}

func newGalleryApp() *galleryApp {
	return &galleryApp{
		seedIndex:      0,
		dark:           false,
		density:        mayu.DensityStandard,
		lastAction:     "Ready",
		railIndex:      0,
		segmentedIndex: 0,
		searchQuery:    "",
		checkboxVal:    true,
		switchVal:      true,
		radioChoice:    "Standard",
		username:       "Alex",
		email:          "alex@example.com",
		sliderVal:      60.0,
		brightnessVal:  65.0,
		volumeVal:      35.0,
		lightSwitchVal: 80.0,
		accordionOpen:  true,
		menuOpen:       false,
		dialogOpen:     false,
		snackbarOpen:   false,
		snackbarMsg:    "Settings saved successfully",
		toastShapeIdx:  0,
		toastColorIdx:  0,
		sideSheetOpen:  false,
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

	// Material 3 toasts show in this host, over everything else.
	toastShape := mayu.ToastCapsule
	if a.toastShapeIdx == 1 {
		toastShape = mayu.ToastCorner
	}
	toastColor := mayu.ToastSurface
	if a.toastColorIdx == 1 {
		toastColor = mayu.ToastInverse
	}
	mayu.ToastHostWith(c, mayu.ToastHostOptions{
		Placement: mayu.ToastBottomEnd,
		Shape:     toastShape,
		Color:     toastColor,
	})

	// Top-level layout: Navigation Rail on left, Scrollable content on right
	ui.Row(c).Fill().Children(func() {
		// ==========================================
		// Left: Navigation Rail (Desktop Skeleton)
		// ==========================================
		mayu.NavigationRail(c, mayu.NavigationRailOptions{
			Selected: &a.railIndex,
			Header: func() {
				mayu.FAB(c, symbols.Add).OnClick(func() {
					a.clicks++
					a.lastAction = "Clicked Rail FAB"
				})
			},
			Destinations: []mayu.NavDestination{
				{Icon: symbols.Send, Label: "Overview", Badge: 3},
				{Icon: symbols.Favorite, Label: "Favorites"},
				{Icon: symbols.Settings, Label: "Config"},
			},
			Footer: func() {
				mayu.InitialsAvatar(c, "Mayu User")
			},
			OnSelect: func(idx int) {
				a.lastAction = fmt.Sprintf("Switched to Tab %d", idx)
			},
		})

		// ==========================================
		// Right: Scrollable Main Content
		// ==========================================
		ui.Scroll(c).Grow(1).FillHeight().Children(func() {
			ui.Column(c).FillWidth().Padding(24).Gap(20).Children(func() {
				// Search Bar & Header
				ui.Row(c).AlignItems(ui.Center).Gap(16).FillWidth().Children(func() {
					mayu.SearchBar(c, &a.searchQuery, "Search Mayu components...").
						Grow(1)

					// Density & Theme Mode Controls
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

				// 1. Dynamic Seed Colors
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

				// 2. Segmented Buttons (Timeline / Filter)
				ui.Column(c).Gap(10).Children(func() {
					ui.Text(c, "Segmented Button (M3 View Switcher)").FontSize(theme.TitleMedium.Size).Bold()
					mayu.SegmentedButton(c, &a.segmentedIndex, []mayu.SegmentOption{
						{Label: "Daily", Icon: symbols.LightMode},
						{Label: "Weekly"},
						{Label: "Monthly"},
						{Label: "Yearly"},
					})
				})

				mayu.HorizontalDivider(c)

				// 3. Capsule Sliders (Control Center / 宽体胶囊滑动条)
				ui.Column(c).Gap(12).Children(func() {
					ui.Text(c, "Capsule Sliders (Control Center / 宽体胶囊滑动条)").FontSize(theme.TitleMedium.Size).Bold()

					ui.Row(c).Gap(24).AlignItems(ui.Center).Children(func() {
						// Vertical Brightness Capsule
						ui.Column(c).AlignItems(ui.Center).Gap(8).Children(func() {
							mayu.VerticalCapsuleSlider(c, &a.brightnessVal, 0, 100, symbols.Lightbulb)
							ui.Text(c, fmt.Sprintf("亮度 %.0f%%", a.brightnessVal)).FontSize(12).TextColor(scheme.OnSurfaceVariant)
						})

						// Vertical Volume Capsule
						ui.Column(c).AlignItems(ui.Center).Gap(8).Children(func() {
							mayu.VerticalCapsuleSlider(c, &a.volumeVal, 0, 100, symbols.VolumeUp)
							ui.Text(c, fmt.Sprintf("音量 %.0f%%", a.volumeVal)).FontSize(12).TextColor(scheme.OnSurfaceVariant)
						})

						// Horizontal Wide Capsule ("开灯")
						ui.Column(c).Gap(12).Justify(ui.Center).Children(func() {
							ui.Text(c, "Wide Horizontal Capsule Slider").FontSize(theme.TitleSmall.Size).Bold()
							mayu.HorizontalCapsuleSlider(c, &a.lightSwitchVal, 0, 100, symbols.LightMode, "开灯")
							ui.Text(c, fmt.Sprintf("开关滑块档位: %.0f%%", a.lightSwitchVal)).FontSize(12).TextColor(scheme.OnSurfaceVariant)
						})
					})
				})

				mayu.HorizontalDivider(c)

				// Labeled Slider (标题滑块卡片)
				ui.Column(c).Gap(12).MaxWidthPercent(100).Children(func() {
					ui.Text(c, "Labeled Slider (Card)").FontSize(theme.TitleMedium.Size).Bold()

					ui.Box(c).MaxWidth(560).Children(func() {
						mayu.LabeledSlider(c, &a.brightnessVal, 0, 100, mayu.LabeledSliderOptions{
							Label: "亮度 (Brightness)",
							Icon:  symbols.Lightbulb,
						})
					})
				})

				mayu.HorizontalDivider(c)

				// 4. Buttons & Floating Actions
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

					// Extended FAB & Icon Buttons Row
					ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
						mayu.ExtendedFAB(c, symbols.Add, "Compose").
							OnClick(func() {
								a.clicks++
								a.lastAction = "Clicked Extended FAB"
							})

						mayu.IconButton(c, symbols.Search).
							OnClick(func() { a.lastAction = "Clicked IconButton" })

						mayu.FilledIconButton(c, symbols.Settings).
							OnClick(func() { a.lastAction = "Clicked FilledIconButton" })

						menuAnchor := mayu.OutlinedIconButton(c, symbols.Close).
							OnClick(func() {
								a.menuOpen = !a.menuOpen
							})

						// Dropdown Menu anchored to icon button
						mayu.Menu(c, menuAnchor.Element, &a.menuOpen, func(m *mayu.MenuPanel) {
							m.Item("Copy Link", mayu.MenuItemOptions{
								Icon:     symbols.Send,
								Shortcut: "⌘C",
								OnClick: func() {
									a.lastAction = "Menu Item Copy Link"
									mayu.ShowToast(c, mayu.ToastOptions{
										Title: "已复制到剪贴板",
										Kind:  mayu.ToastSuccess,
									})
								},
							})
							m.Divider()
							m.Item("Settings", mayu.MenuItemOptions{
								Icon: symbols.Settings,
								OnClick: func() {
									a.lastAction = "Menu Item Settings"
								},
							})
						})
					})
				})

				mayu.HorizontalDivider(c)

				// 5. List Items (Settings / Content items)
				ui.Column(c).Gap(8).Children(func() {
					ui.Text(c, "List Items & Avatars").FontSize(theme.TitleMedium.Size).Bold()

					mayu.ListItem(c, "Sarah Connor", mayu.ListItemOptions{
						SupportingText: "Online • Cyberdyne Systems",
						Leading: func() {
							mayu.InitialsAvatar(c, "Sarah Connor")
						},
						Trailing: func() {
							mayu.Badge(c, 2)
						},
						OnClick: func() {
							a.lastAction = "Clicked Sarah Connor ListItem"
						},
					})

					mayu.ListItem(c, "Airplane Mode", mayu.ListItemOptions{
						SupportingText: "Disable Wi-Fi, Bluetooth and Cellular",
						Leading: func() {
							mayu.Avatar(c, symbols.Settings)
						},
						Trailing: func() {
							mayu.Switch(c, &a.switchVal, "")
						},
					})
				})

				mayu.HorizontalDivider(c)

				// 6. Accordion (Collapsible Panel)
				ui.Column(c).Gap(10).Children(func() {
					ui.Text(c, "Accordion / Collapsible Panel").FontSize(theme.TitleMedium.Size).Bold()

					mayu.Accordion(c, &a.accordionOpen, "Advanced System Configuration", func() {
						ui.Column(c).Gap(10).Children(func() {
							ui.Text(c, "Detailed parameters for graphics pipeline and font rendering engine.").
								TextColor(scheme.OnSurfaceVariant).FontSize(13)

							ui.Row(c).Gap(16).Children(func() {
								mayu.Checkbox(c, &a.checkboxVal, "Enable hardware acceleration")
							})
						})
					})
				})

				mayu.HorizontalDivider(c)

				// 7. Cards & Selection Controls
				ui.Column(c).Gap(12).Children(func() {
					ui.Text(c, "Cards & Selection Controls").FontSize(theme.TitleMedium.Size).Bold()

					ui.Row(c).Gap(16).Children(func() {
						mayu.ElevatedCard(c).Width(240).Children(func() {
							ui.Column(c).Gap(6).Children(func() {
								ui.Text(c, "Elevated Card").FontSize(theme.TitleMedium.Size).Bold()
								ui.Text(c, "SurfaceContainerLow with resting elevation shadow.").TextColor(scheme.OnSurfaceVariant).FontSize(13)
							})
						}).OnClick(func() { a.lastAction = "Clicked Elevated Card" })

						mayu.FilledCard(c).Width(240).Children(func() {
							ui.Column(c).Gap(6).Children(func() {
								ui.Text(c, "Filled Card").FontSize(theme.TitleMedium.Size).Bold()
								ui.Text(c, "SurfaceContainerHighest without elevation shadow.").TextColor(scheme.OnSurfaceVariant).FontSize(13)
							})
						}).OnClick(func() { a.lastAction = "Clicked Filled Card" })

						mayu.OutlinedCard(c).Width(240).Children(func() {
							ui.Column(c).Gap(6).Children(func() {
								ui.Text(c, "Outlined Card").FontSize(theme.TitleMedium.Size).Bold()
								ui.Text(c, "Surface with 1dp OutlineVariant clean border.").TextColor(scheme.OnSurfaceVariant).FontSize(13)
							})
						}).OnClick(func() { a.lastAction = "Clicked Outlined Card" })
					})

					ui.Row(c).Gap(20).AlignItems(ui.Center).Children(func() {
						mayu.Checkbox(c, &a.checkboxVal, "Enable notifications")
						mayu.Switch(c, &a.switchVal, "Airplane mode")
						mayu.Radio(c, &a.radioChoice, "Standard", "Standard Tier")
						mayu.Radio(c, &a.radioChoice, "Pro", "Pro Tier")
					})
				})

				mayu.HorizontalDivider(c)

				// 8. Text Fields Section
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

				// 9. Chips, Sliders & Progress Section
				ui.Column(c).Gap(12).Children(func() {
					ui.Text(c, "Chips, Sliders & Progress").FontSize(theme.TitleMedium.Size).Bold()

					ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
						mayu.Chip(c, "Action Chip").
							Icon(symbols.Settings).
							OnClick(func() { a.lastAction = "Clicked Action Chip" })

						mayu.FilterChip(c, "Filter (Active)", true).
							OnClick(func() { a.lastAction = "Toggled Filter Chip" })

						mayu.FilterChip(c, "Filter (Inactive)", false).
							OnClick(func() { a.lastAction = "Toggled Inactive Filter Chip" })
					})

					ui.Column(c).Gap(8).Children(func() {
						ui.Text(c, fmt.Sprintf("Standard Slider: %.0f%%", a.sliderVal)).TextColor(scheme.OnSurfaceVariant)
						mayu.Slider(c, &a.sliderVal, 0, 100)
						mayu.LinearProgress(c, float32(a.sliderVal/100.0))
					})
				})

				// Toast & Snackbar notifications
				ui.Column(c).Gap(12).Children(func() {
					ui.Text(c, "Toast & Snackbar Notifications").FontSize(theme.TitleMedium.Size).Bold()

					ui.Column(c).Gap(8).Children(func() {
						ui.Text(c, "Toast shape").FontSize(theme.LabelLarge.Size).TextColor(scheme.OnSurfaceVariant)
						mayu.SegmentedButton(c, &a.toastShapeIdx, []mayu.SegmentOption{
							{Label: "Capsule 胶囊"},
							{Label: "Corner 圆角矩形"},
						})
					})

					ui.Column(c).Gap(8).Children(func() {
						ui.Text(c, "Toast color").FontSize(theme.LabelLarge.Size).TextColor(scheme.OnSurfaceVariant)
						mayu.SegmentedButton(c, &a.toastColorIdx, []mayu.SegmentOption{
							{Label: "Surface 应用表面"},
							{Label: "Inverse 反色"},
						})
					})

					ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
						mayu.FilledTonalButton(c, "Toast: Success").
							Icon(symbols.Check).
							OnClick(func() {
								mayu.ShowToast(c, mayu.ToastOptions{
									Title:       "已保存",
									Description: "改动已写入磁盘",
									Kind:        mayu.ToastSuccess,
									Close:       true,
								})
							})

						mayu.FilledTonalButton(c, "Toast: Error").
							Icon(symbols.Close).
							OnClick(func() {
								mayu.ShowToast(c, mayu.ToastOptions{
									Title:       "上传失败",
									Description: "网络连接不可用，请稍后重试",
									Kind:        mayu.ToastError,
									Close:       true,
								})
							})

						mayu.FilledTonalButton(c, "Toast: Info").
							Icon(symbols.Info).
							OnClick(func() {
								mayu.ShowToast(c, mayu.ToastOptions{
									Title:   "Material You 已启用",
									Kind:    mayu.ToastInfo,
									Timeout: 3 * time.Second,
								})
							})

						mayu.FilledTonalButton(c, "Toast: Undo").
							OnClick(func() {
								mayu.ShowToast(c, mayu.ToastOptions{
									Title:    "已将 3 个项目移至废纸篓",
									Kind:     mayu.ToastInfo,
									Action:   "撤销",
									OnAction: func() { a.lastAction = "撤销了删除操作" },
									Close:    true,
								})
							})

						mayu.OutlinedButton(c, "Show Snackbar").
							OnClick(func() {
								a.snackbarOpen = true
								a.snackbarMsg = "Settings saved successfully"
							})
					})
				})

				// Modal surfaces: a side sheet anchored to the right edge.
				ui.Column(c).Gap(12).Children(func() {
					ui.Text(c, "Modal Surfaces").FontSize(theme.TitleMedium.Size).Bold()

					ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
						mayu.FilledTonalButton(c, "Open Side Sheet").
							Icon(symbols.Settings).
							OnClick(func() {
								a.sideSheetOpen = true
							})

						mayu.OutlinedButton(c, "Open Alert Dialog").
							OnClick(func() {
								a.dialogOpen = true
							})
					})
				})

				// Snackbar trigger
				if a.snackbarOpen {
					mayu.Snackbar(c, a.snackbarMsg, mayu.SnackbarOptions{
						ActionText: "Dismiss",
						OnAction: func() {
							a.snackbarOpen = false
						},
					})
				}

				// Status footer
				ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
					ui.Text(c, fmt.Sprintf("Clicks: %d", a.clicks)).Bold()
					ui.Text(c, "•").TextColor(scheme.Outline)
					ui.Text(c, fmt.Sprintf("Tab=%d, Segment=%d, Tier=%s", a.railIndex, a.segmentedIndex, a.radioChoice)).TextColor(scheme.OnSurfaceVariant)
					ui.Text(c, "•").TextColor(scheme.Outline)
					ui.Text(c, fmt.Sprintf("Status: %s", a.lastAction)).TextColor(scheme.Primary)
				})
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

	// Modal side sheet anchored to the right edge while sideSheetOpen.
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
			ui.Text(c, "显示选项").FontSize(theme.TitleSmall.Size).Bold().TextColor(scheme.OnSurface)
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
