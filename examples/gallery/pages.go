package main

import (
	"fmt"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// componentsPage shows the widgets themselves: what each one looks like, and
// what it does when it is used.
func (a *galleryApp) componentsPage(c *ui.Context, t *theme.Theme) {
	shown := 0
	sec := func(title string, body func()) {
		if a.section(c, title, body) {
			shown++
		}
	}

	sec("Buttons & Floating Actions", func() {
		ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
			for _, b := range []struct {
				label   string
				icon    *ui.SVG
				variant mayu.ButtonVariant
			}{
				{"Filled Button", symbols.Send, mayu.VariantFilled},
				{"Elevated Button", symbols.Add, mayu.VariantElevated},
				{"Filled Tonal", symbols.Favorite, mayu.VariantFilledTonal},
				{"Outlined", symbols.Settings, mayu.VariantOutlined},
				{"Text Button", nil, mayu.VariantText},
			} {
				name := b.label
				btn := mayu.Button(c, name, b.variant).Icon(b.icon)
				btn.OnClick(func() {
					a.clicks++
					a.lastAction = "Clicked the " + name
				})
			}
		})

		ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
			mayu.ExtendedFAB(c, symbols.Add, "Compose").
				OnClick(func() {
					a.clicks++
					a.lastAction = "Clicked the extended FAB"
				})

			mayu.FAB(c, symbols.Favorite).
				OnClick(func() { a.lastAction = "Clicked the FAB" })

			mayu.IconButton(c, symbols.Search).
				OnClick(func() { a.lastAction = "Clicked a standard icon button" })
			mayu.FilledIconButton(c, symbols.Settings).
				OnClick(func() { a.lastAction = "Clicked a filled icon button" })
			mayu.FilledTonalIconButton(c, symbols.Add).
				OnClick(func() { a.lastAction = "Clicked a filled tonal icon button" })
			mayu.OutlinedIconButton(c, symbols.Close).
				OnClick(func() { a.lastAction = "Clicked an outlined icon button" })
		})
	})

	sec("Cards", func() {
		ui.Row(c).Gap(16).Children(func() {
			mayu.ElevatedCard(c).Width(240).Children(func() {
				ui.Column(c).Gap(6).Children(func() {
					ui.Text(c, "Elevated Card").FontSize(theme.TitleMedium.Size).Bold()
					ui.Text(c, "SurfaceContainerLow with a resting elevation shadow.").
						TextColor(t.Scheme.OnSurfaceVariant).FontSize(13)
				})
			}).OnClick(func() { a.lastAction = "Clicked an elevated card" })

			mayu.FilledCard(c).Width(240).Children(func() {
				ui.Column(c).Gap(6).Children(func() {
					ui.Text(c, "Filled Card").FontSize(theme.TitleMedium.Size).Bold()
					ui.Text(c, "SurfaceContainerHighest without an elevation shadow.").
						TextColor(t.Scheme.OnSurfaceVariant).FontSize(13)
				})
			}).OnClick(func() { a.lastAction = "Clicked a filled card" })

			mayu.OutlinedCard(c).Width(240).Children(func() {
				ui.Column(c).Gap(6).Children(func() {
					ui.Text(c, "Outlined Card").FontSize(theme.TitleMedium.Size).Bold()
					ui.Text(c, "Surface with a 1dp OutlineVariant border.").
						TextColor(t.Scheme.OnSurfaceVariant).FontSize(13)
				})
			}).OnClick(func() { a.lastAction = "Clicked an outlined card" })
		})
	})

	sec("Selection Controls", func() {
		ui.Row(c).Gap(24).AlignItems(ui.Center).Children(func() {
			mayu.Checkbox(c, &a.checkboxVal, "Enable notifications")
			mayu.Switch(c, &a.switchVal, "Airplane mode")
			mayu.Radio(c, &a.radioChoice, "Standard", "Standard Tier")
			mayu.Radio(c, &a.radioChoice, "Pro", "Pro Tier")
		})
	})

	sec("Text Fields", func() {
		ui.Row(c).Gap(20).FillWidth().Children(func() {
			mayu.FilledTextField(c, &a.username, "Username").
				Placeholder("Enter username")
			mayu.OutlinedTextField(c, &a.email, "Email Address").
				Placeholder("Enter email")
		})
	})

	sec("Chips", func() {
		ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
			mayu.Chip(c, "Action Chip").
				Icon(symbols.Settings).
				OnClick(func() { a.lastAction = "Clicked an action chip" })
			mayu.FilterChip(c, "Filter (Active)", true).
				OnClick(func() { a.lastAction = "Toggled a chosen filter chip" })
			mayu.FilterChip(c, "Filter (Inactive)", false).
				OnClick(func() { a.lastAction = "Toggled an unchosen filter chip" })
		})
	})

	sec("Sliders & Progress", func() {
		ui.Column(c).Gap(8).Children(func() {
			ui.Text(c, fmt.Sprintf("Standard Slider: %.0f%%", a.sliderVal)).
				TextColor(t.Scheme.OnSurfaceVariant)
			mayu.Slider(c, &a.sliderVal, 0, 100)
			mayu.LinearProgress(c, float32(a.sliderVal/100.0))
		})
	})

	sec("List Items & Avatars", func() {
		ui.Column(c).Gap(8).Children(func() {
			mayu.ListItem(c, "Sarah Connor", mayu.ListItemOptions{
				SupportingText: "Online • Cyberdyne Systems",
				Leading: func() {
					mayu.InitialsAvatar(c, "Sarah Connor")
				},
				Trailing: func() {
					mayu.Badge(c, 2)
				},
				OnClick: func() { a.lastAction = "Clicked the Sarah Connor row" },
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
	})

	sec("Accordion", func() {
		mayu.Accordion(c, &a.accordionOpen, "Advanced System Configuration", func() {
			ui.Column(c).Gap(10).Children(func() {
				ui.Text(c, "Detailed parameters for the graphics pipeline and the font rendering engine.").
					TextColor(t.Scheme.OnSurfaceVariant).FontSize(13)
				mayu.Checkbox(c, &a.checkboxVal, "Enable hardware acceleration")
			})
		})
	})

	a.emptyState(c, shown, t)
}

// overlaysPage shows the surfaces that float over a page: what they look like,
// how they are anchored, and how they close.
func (a *galleryApp) overlaysPage(c *ui.Context, t *theme.Theme) {
	shown := 0
	sec := func(title string, body func()) {
		if a.section(c, title, body) {
			shown++
		}
	}

	sec("Tooltips & Dividers", func() {
		ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
			for _, tip := range []struct {
				icon *ui.SVG
				text string
			}{
				{symbols.VolumeUp, "音量"},
				{symbols.Settings, "设置"},
				{symbols.Favorite, "收藏"},
				{symbols.Search, "搜索"},
			} {
				anchor := mayu.FilledTonalIconButton(c, tip.icon)
				mayu.Tooltip(c, anchor.Element, tip.text)
			}

			mayu.VerticalDivider(c).Height(28)

			chip := mayu.Chip(c, "悬停查看提示").Icon(symbols.Info)
			mayu.Tooltip(c, chip.Element, "这是 M3 Tooltip")
		})
	})

	sec("Dropdown Menu", func() {
		ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
			ui.Text(c, "Anchored to the button beside it:").
				TextColor(t.Scheme.OnSurfaceVariant)

			anchor := mayu.OutlinedIconButton(c, symbols.Close).
				OnClick(func() { a.menuOpen = !a.menuOpen })

			mayu.Menu(c, anchor.Element, &a.menuOpen, func(m *mayu.MenuPanel) {
				m.Item("Copy Link", mayu.MenuItemOptions{
					Icon:     symbols.Send,
					Shortcut: "⌘C",
					OnClick: func() {
						a.lastAction = "Chose Copy Link"
						mayu.ShowToast(c, mayu.ToastOptions{
							Title: "已复制到剪贴板",
							Kind:  mayu.ToastSuccess,
						})
					},
				})
				m.Divider()
				m.Item("Settings", mayu.MenuItemOptions{
					Icon:    symbols.Settings,
					OnClick: func() { a.lastAction = "Chose Settings" },
				})
			})
		})
	})

	sec("Toast & Snackbar", func() {
		ui.Row(c).Gap(16).AlignItems(ui.End).Children(func() {
			ui.Column(c).Gap(8).Width(240).Children(func() {
				ui.Text(c, "Toast shape").FontSize(theme.LabelLarge.Size).
					TextColor(t.Scheme.OnSurfaceVariant)
				mayu.SegmentedButton(c, &a.toastShape, []mayu.SegmentOption{
					{Label: "Capsule"},
					{Label: "Corner"},
				})
			})

			ui.Column(c).Gap(8).Width(240).Children(func() {
				ui.Text(c, "Toast color").FontSize(theme.LabelLarge.Size).
					TextColor(t.Scheme.OnSurfaceVariant)
				mayu.SegmentedButton(c, &a.toastColor, []mayu.SegmentOption{
					{Label: "Surface"},
					{Label: "Inverse"},
				})
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

			mayu.TextButton(c, "Close all toasts").
				OnClick(func() { mayu.CloseToast(c, "") })

			mayu.OutlinedButton(c, "Show Snackbar").
				OnClick(func() {
					a.snackbarOpen = true
					a.snackbarMsg = "Settings saved successfully"
				})
		})

		if a.snackbarOpen {
			mayu.Snackbar(c, a.snackbarMsg, mayu.SnackbarOptions{
				ActionText: "Dismiss",
				OnAction:   func() { a.snackbarOpen = false },
			})
		}
	})

	sec("Modal Surfaces", func() {
		ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
			mayu.FilledTonalButton(c, "Open Side Sheet").
				Icon(symbols.Settings).
				OnClick(func() { a.sideSheetOpen = true })

			mayu.OutlinedButton(c, "Open Alert Dialog").
				OnClick(func() { a.dialogOpen = true })
		})
	})

	a.emptyState(c, shown, t)
}

// themePage shows the design system itself: the colors a seed generates, the
// scales the tokens come from, and the controls that restyle the window.
func (a *galleryApp) themePage(c *ui.Context, t *theme.Theme) {
	shown := 0
	sec := func(title string, body func()) {
		if a.section(c, title, body) {
			shown++
		}
	}

	sec("Dynamic Seed Color", func() {
		ui.Column(c).Gap(8).Children(func() {
			ui.Text(c, "Each choice carries a swatch of the primary its seed generates.").
				TextColor(t.Scheme.OnSurfaceVariant).FontSize(13)

			ui.Row(c).Gap(8).Children(func() {
				for i, s := range presetSeeds {
					idx := i
					// One element per seed, restyled when it is the chosen
					// one: building an outlined button and replacing it with a
					// filled one would draw the selected seed twice, the inert
					// copy first.
					btn := mayu.OutlinedButton(c, s.name)
					if idx == a.seedIndex {
						btn.Variant(mayu.VariantFilled)
					}
					btn.Leading(func() {
						ui.Box(c).Size(14, 14).Radius(7).Shrink(0).
							Background(theme.SchemeFromSeed(s.color, a.dark).Primary).
							Border(1, t.Scheme.OutlineVariant)
					})
					btn.OnClick(func() {
						a.seedIndex = idx
						a.clicks++
						a.lastAction = fmt.Sprintf("Switched the seed to %s", s.name)
					})
				}
			})
		})
	})

	sec("Tonal Elevation", func() {
		ui.Column(c).Gap(8).Children(func() {
			ui.Text(c, "The surface containers, from the lowest to the highest.").
				TextColor(t.Scheme.OnSurfaceVariant).FontSize(13)

			ui.Row(c).Gap(10).Children(func() {
				for _, sw := range []struct {
					name string
					bg   ui.Color
					fg   ui.Color
				}{
					{"SurfaceContainerLowest", t.Scheme.SurfaceContainerLowest, t.Scheme.OnSurface},
					{"SurfaceContainerLow", t.Scheme.SurfaceContainerLow, t.Scheme.OnSurface},
					{"SurfaceContainer", t.Scheme.SurfaceContainer, t.Scheme.OnSurface},
					{"SurfaceContainerHigh", t.Scheme.SurfaceContainerHigh, t.Scheme.OnSurface},
					{"SurfaceContainerHighest", t.Scheme.SurfaceContainerHighest, t.Scheme.OnSurface},
					{"PrimaryContainer", t.Scheme.PrimaryContainer, t.Scheme.OnPrimaryContainer},
					{"TertiaryContainer", t.Scheme.TertiaryContainer, t.Scheme.OnTertiaryContainer},
				} {
					ui.Box(c).Width(112).Height(56).Radius(theme.CornerMedium).
						Background(sw.bg).Center().Padding(6).Children(func() {
						ui.Text(c, sw.name).FontSize(10).TextColor(sw.fg)
					})
				}
			})
		})
	})

	sec("Typography & Shape", func() {
		ui.Row(c).Gap(24).AlignItems(ui.Start).Children(func() {
			ui.Column(c).Gap(4).Children(func() {
				for _, l := range []struct {
					name string
					ts   theme.TextStyle
				}{
					{"HeadlineMedium", theme.HeadlineMedium},
					{"TitleMedium", theme.TitleMedium},
					{"BodyLarge", theme.BodyLarge},
					{"LabelLarge", theme.LabelLarge},
				} {
					ui.Text(c, l.name+" — "+fmt.Sprint(l.ts.Size)+"dp").
						FontSize(l.ts.Size).TextColor(t.Scheme.OnSurface)
				}
			})

			ui.Column(c).Gap(6).Children(func() {
				for _, r := range []struct {
					name   string
					radius float32
				}{
					{"None", theme.CornerNone},
					{"ExtraSmall", theme.CornerExtraSmall},
					{"Small", theme.CornerSmall},
					{"Medium", theme.CornerMedium},
					{"Large", theme.CornerLarge},
					{"ExtraLarge", theme.CornerExtraLarge},
				} {
					ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
						ui.Box(c).Size(36, 24).Radius(r.radius).
							Background(t.Scheme.PrimaryContainer)
						ui.Text(c, r.name).FontSize(12).TextColor(t.Scheme.OnSurfaceVariant)
					})
				}
			})
		})
	})

	sec("Segmented Button", func() {
		ui.Column(c).Gap(8).Children(func() {
			ui.Text(c, "Segments share the track's width, so no gap trails the last one.").
				TextColor(t.Scheme.OnSurfaceVariant).FontSize(13)
			mayu.SegmentedButton(c, &a.segIdx, []mayu.SegmentOption{
				{Label: "Daily", Icon: symbols.LightMode},
				{Label: "Weekly"},
				{Label: "Monthly"},
				{Label: "Yearly"},
			})
		})
	})

	sec("Capsule Sliders", func() {
		ui.Row(c).Gap(24).AlignItems(ui.Center).Children(func() {
			ui.Column(c).AlignItems(ui.Center).Gap(8).Children(func() {
				mayu.VerticalCapsuleSlider(c, &a.brightnessVal, 0, 100, symbols.Lightbulb)
				ui.Text(c, fmt.Sprintf("亮度 %.0f%%", a.brightnessVal)).
					FontSize(12).TextColor(t.Scheme.OnSurfaceVariant)
			})

			ui.Column(c).AlignItems(ui.Center).Gap(8).Children(func() {
				mayu.VerticalCapsuleSlider(c, &a.volumeVal, 0, 100, symbols.VolumeUp)
				ui.Text(c, fmt.Sprintf("音量 %.0f%%", a.volumeVal)).
					FontSize(12).TextColor(t.Scheme.OnSurfaceVariant)
			})

			ui.Column(c).Gap(12).Justify(ui.Center).Children(func() {
				ui.Text(c, "Wide Horizontal Capsule Slider").
					FontSize(theme.TitleSmall.Size).Bold()
				mayu.HorizontalCapsuleSlider(c, &a.lightSwitchVal, 0, 100, symbols.LightMode, "开灯")
				ui.Text(c, fmt.Sprintf("开关滑块档位: %.0f%%", a.lightSwitchVal)).
					FontSize(12).TextColor(t.Scheme.OnSurfaceVariant)
			})
		})
	})

	sec("Labeled Slider", func() {
		ui.Box(c).MaxWidth(560).Children(func() {
			mayu.LabeledSlider(c, &a.brightnessVal, 0, 100, mayu.LabeledSliderOptions{
				Label: "亮度 (Brightness)",
				Icon:  symbols.Lightbulb,
			})
		})
	})

	a.emptyState(c, shown, t)
}
