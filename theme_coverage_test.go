package mayu_test

import (
	"image"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
	"github.com/harukizmoe/mayu/symbols"
)

// backdrop is the neutral a component is painted on, so that the pixels it
// paints can be told apart from the window's.
var backdrop = ui.Hex("#808080")

// backdropSlack is how far a pixel may differ from the backdrop in one channel
// and still count as backdrop: antialiasing leaves a rim around every shape.
const backdropSlack = 6

type sample struct {
	name  string
	build func(c *ui.Context)
	// after runs against the tester once the first frame is up, to open a
	// popup or reach a hover.
	after func(tt *ui.Tester)
}

// samples is one entry per component the library exports, so that a component
// carrying a color of its own is caught here rather than by eye. One component
// per entry, deliberately: a sample holding a whole family would let one
// miscolored member hide behind the others.
func samples() []sample {
	return []sample{
		pass("FilledButton", func(c *ui.Context) { mayu.FilledButton(c, "Filled") }),
		pass("ElevatedButton", func(c *ui.Context) { mayu.ElevatedButton(c, "Elevated") }),
		pass("FilledTonalButton", func(c *ui.Context) { mayu.FilledTonalButton(c, "Tonal") }),
		pass("OutlinedButton", func(c *ui.Context) { mayu.OutlinedButton(c, "Outlined") }),
		pass("TextButton", func(c *ui.Context) { mayu.TextButton(c, "Text") }),

		pass("FAB", func(c *ui.Context) { mayu.FAB(c, symbols.Add) }),
		pass("ExtendedFAB", func(c *ui.Context) { mayu.ExtendedFAB(c, symbols.Send, "Send") }),
		pass("IconButton", func(c *ui.Context) { mayu.IconButton(c, symbols.Search) }),
		pass("FilledIconButton", func(c *ui.Context) { mayu.FilledIconButton(c, symbols.Settings) }),
		pass("FilledTonalIconButton", func(c *ui.Context) { mayu.FilledTonalIconButton(c, symbols.Favorite) }),
		pass("OutlinedIconButton", func(c *ui.Context) { mayu.OutlinedIconButton(c, symbols.Close) }),

		pass("ElevatedCard", func(c *ui.Context) {
			mayu.ElevatedCard(c).Width(200).Children(func() { ui.Text(c, "Elevated") })
		}),
		pass("FilledCard", func(c *ui.Context) {
			mayu.FilledCard(c).Width(200).Children(func() { ui.Text(c, "Filled") })
		}),
		pass("OutlinedCard", func(c *ui.Context) {
			mayu.OutlinedCard(c).Width(200).Children(func() { ui.Text(c, "Outlined") })
		}),

		pass("Checkbox", func(c *ui.Context) {
			on := true
			mayu.Checkbox(c, &on, "Checkbox")
		}),
		pass("Switch", func(c *ui.Context) {
			on := true
			mayu.Switch(c, &on, "Switch")
		}),
		pass("Radio", func(c *ui.Context) {
			pick := "a"
			mayu.Radio(c, &pick, "a", "Radio")
		}),

		pass("FilledTextField", func(c *ui.Context) {
			var v string
			ui.Box(c).Width(240).Children(func() { mayu.FilledTextField(c, &v, "Filled") })
		}),
		pass("OutlinedTextField", func(c *ui.Context) {
			var v string
			ui.Box(c).Width(240).Children(func() { mayu.OutlinedTextField(c, &v, "Outlined") })
		}),
		pass("FilledTextArea", func(c *ui.Context) {
			var v string
			ui.Box(c).Width(240).Children(func() { mayu.FilledTextArea(c, &v, "Area") })
		}),
		pass("OutlinedTextArea", func(c *ui.Context) {
			var v string
			ui.Box(c).Width(240).Children(func() { mayu.OutlinedTextArea(c, &v, "Area") })
		}),

		pass("Chip", func(c *ui.Context) { mayu.Chip(c, "Assist") }),
		pass("FilterChip", func(c *ui.Context) { mayu.FilterChip(c, "Filter", true) }),
		pass("Badge", func(c *ui.Context) { mayu.Badge(c, 5) }),

		pass("Tabs", func(c *ui.Context) {
			i := 1
			ui.Box(c).Width(300).Children(func() {
				mayu.Tabs(c, &i, []string{"One", "Two", "Three"})
			})
		}),
		pass("SegmentedButton", func(c *ui.Context) {
			i := 1
			ui.Box(c).Width(300).Children(func() {
				mayu.SegmentedButton(c, &i, []mayu.SegmentOption{
					{Label: "Day"}, {Label: "Week"}, {Label: "Month"},
				})
			})
		}),
		pass("Select", func(c *ui.Context) {
			v := "b"
			ui.Box(c).Width(240).Children(func() {
				mayu.Select(c, &v, "Choose", []mayu.SelectOption[string]{
					{Value: "a", Label: "Alpha"}, {Value: "b", Label: "Beta"},
				})
			})
		}),
		pass("Autocomplete", func(c *ui.Context) {
			var q string
			ui.Box(c).Width(240).Children(func() {
				mayu.Autocomplete(c, &q, "Font", []string{"Helvetica", "Times"})
			})
		}),

		pass("Slider", func(c *ui.Context) {
			v := 40.0
			ui.Box(c).Width(240).Children(func() { mayu.Slider(c, &v, 0, 100) })
		}),
		pass("LabeledSlider", func(c *ui.Context) {
			v := 40.0
			ui.Box(c).Width(320).Children(func() {
				mayu.LabeledSlider(c, &v, 0, 100, mayu.LabeledSliderOptions{
					Label: "Brightness", Icon: symbols.Lightbulb,
				})
			})
		}),
		pass("VerticalCapsuleSlider", func(c *ui.Context) {
			v := 60.0
			mayu.VerticalCapsuleSlider(c, &v, 0, 100, symbols.Lightbulb)
		}),
		pass("HorizontalCapsuleSlider", func(c *ui.Context) {
			v := 60.0
			mayu.HorizontalCapsuleSlider(c, &v, 0, 100, symbols.LightMode, "开灯")
		}),

		pass("LinearProgress", func(c *ui.Context) {
			ui.Box(c).Width(200).Children(func() { mayu.LinearProgress(c, 0.6) })
		}),
		pass("CircularProgress", func(c *ui.Context) {
			mayu.CircularProgress(c, 0.6, mayu.CircularProgressOptions{ShowValue: true})
		}),
		pass("HorizontalDivider", func(c *ui.Context) {
			ui.Box(c).Width(200).Children(func() { mayu.HorizontalDivider(c) })
		}),
		pass("VerticalDivider", func(c *ui.Context) { mayu.VerticalDivider(c).Height(80) }),

		pass("ListItem", func(c *ui.Context) {
			ui.Box(c).Width(320).Children(func() {
				mayu.ListItem(c, "Sarah Connor", mayu.ListItemOptions{SupportingText: "Online"})
			})
		}),
		pass("Avatar", func(c *ui.Context) { mayu.Avatar(c, symbols.Settings) }),
		pass("InitialsAvatar", func(c *ui.Context) { mayu.InitialsAvatar(c, "Sarah Connor") }),
		pass("Accordion", func(c *ui.Context) {
			open := true
			ui.Box(c).Width(320).Children(func() {
				mayu.Accordion(c, &open, "Advanced", func() { ui.Text(c, "Inside") })
			})
		}),
		pass("SearchBar", func(c *ui.Context) {
			var q string
			ui.Box(c).Width(280).Children(func() { mayu.SearchBar(c, &q, "Search") })
		}),
		pass("NavigationRail", func(c *ui.Context) {
			i := 0
			ui.Row(c).Fill().Children(func() {
				mayu.NavigationRail(c, mayu.NavigationRailOptions{
					Selected: &i,
					Destinations: []mayu.NavDestination{
						{Icon: symbols.Settings, Label: "One", Badge: 3},
						{Icon: symbols.Palette, Label: "Two"},
					},
				})
			})
		}),

		{
			name: "Menu",
			after: func(tt *ui.Tester) {
				if err := tt.Click("Open"); err != nil {
					return
				}
			},
			build: func(c *ui.Context) {
				open := true
				anchor := mayu.FilledButton(c, "Open")
				mayu.Menu(c, anchor.Element, &open, func(m *mayu.MenuPanel) {
					m.Item("Copy", mayu.MenuItemOptions{Icon: symbols.Send, Shortcut: "⌘C"})
					m.Divider()
					m.Item("Delete", mayu.MenuItemOptions{Icon: symbols.Close})
				})
			},
		},
		pass("Snackbar", func(c *ui.Context) {
			ui.Box(c).Width(360).Children(func() {
				mayu.Snackbar(c, "Message archived", mayu.SnackbarOptions{ActionText: "Undo"})
			})
		}),
		pass("Toast", func(c *ui.Context) {
			mayu.ToastHostWith(c, mayu.ToastHostOptions{Placement: mayu.ToastBottomCenter})
			mayu.ShowToast(c, mayu.ToastOptions{
				Title: "已保存", Description: "改动已写入磁盘", Close: true,
			})
		}),
		pass("AlertDialog", func(c *ui.Context) {
			open := true
			mayu.AlertDialog(c, &open, mayu.AlertDialogOptions{
				Title: "Confirm", Body: "Are you sure?", ConfirmText: "Yes", DismissText: "No",
			})
		}),
		pass("SideSheet", func(c *ui.Context) {
			open := true
			mayu.SideSheet(c, &open, mayu.SideSheetOptions{
				Title: "显示与亮度", Subtitle: "DP-1",
			}, func() { ui.Text(c, "Inside") })
		}),

		{
			name: "Tooltip",
			after: func(tt *ui.Tester) {
				r, ok := tt.Find("Hover me")
				if !ok {
					return
				}
				tt.Move(r.X+r.W/2, r.Y+r.H/2)
				time.Sleep(700 * time.Millisecond) // the tooltip's delay
			},
			build: func(c *ui.Context) {
				anchor := mayu.FilledTonalButton(c, "Hover me")
				mayu.Tooltip(c, anchor.Element, "A tooltip")
			},
		},
	}
}

// pass is a sample of one component that needs nothing done to it beyond
// building.
func pass(name string, build func(c *ui.Context)) sample {
	return sample{name: name, build: build}
}

// Every component must paint from the theme it is given. Rendered on a neutral
// backdrop in both modes, the pixels it paints have to change with the mode: a
// component carrying a color of its own — the black slab a light app once
// showed — paints the same either way and fails here.
//
// This asks only that the paint follows the theme, not which way it moves,
// because a transparent component paints its ink rather than a surface, and its
// ink is dark in a light theme and light in a dark one. Contrast is a property
// of the tokens instead, and theme/contrast_test.go checks it there.
func TestEveryComponentPaintsFromItsTheme(t *testing.T) {
	for _, s := range samples() {
		t.Run(s.name, func(t *testing.T) {
			light := paint(t, s, false)
			dark := paint(t, s, true)

			// The pixels the component painted, in either frame.
			var n, diff int
			b := light.Bounds()
			for y := b.Min.Y; y < b.Max.Y; y++ {
				for x := b.Min.X; x < b.Max.X; x++ {
					l, d := colorAt(light, x, y), colorAt(dark, x, y)
					if sameAsBackdrop(l) && sameAsBackdrop(d) {
						continue
					}
					n++
					if l != d {
						diff++
					}
				}
			}

			if n < 30 {
				t.Fatalf("the component painted %d pixels: nothing to measure", n)
			}
			if diff < n/4 {
				t.Errorf("only %d of its %d pixels differ between the themes: it does not use the theme",
					diff, n)
			}
		})
	}
}

// paint renders one sample in one mode on the neutral backdrop.
func paint(t *testing.T, s sample, dark bool) *image.RGBA {
	t.Helper()
	tt := ui.NewTester(func(c *ui.Context) {
		mayu.Use(c, mayu.ThemeOptions{Dark: dark, Density: mayu.DensityStandard})
		ui.Box(c).Fill().Background(backdrop).Children(func() {
			ui.Column(c).Padding(16).Children(func() { s.build(c) })
		})
	}, 420, 260)

	if s.after != nil {
		s.after(tt)
	}
	// A component showing on the frame after its first, as a toast does, is
	// painted by now.
	tt.Frame()
	return tt.Image()
}

func colorAt(img *image.RGBA, x, y int) ui.Color {
	i := img.PixOffset(x, y)
	return ui.RGB(img.Pix[i], img.Pix[i+1], img.Pix[i+2])
}

// sameAsBackdrop reports a pixel the component left alone.
func sameAsBackdrop(c ui.Color) bool {
	return near(c.R, backdrop.R) && near(c.G, backdrop.G) && near(c.B, backdrop.B)
}

func near(a, b uint8) bool {
	d := int(a) - int(b)
	return d <= backdropSlack && d >= -backdropSlack
}
