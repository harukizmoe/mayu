package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

func TestToastHostShowsAndActs(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	shown, acted, closed := false, false, false

	tt := ui.NewTester(func(c *ui.Context) {
		ToastHost(c, ToastHostOptions{Placement: ToastBottomEnd, MaxWidth: 380}, th)

		if !shown {
			shown = true
			ShowToast(c, ToastOptions{
				Title:       "已保存",
				Description: "改动已写入磁盘",
				Kind:        ToastSuccess,
				Action:      "撤销",
				OnAction:    func() { acted = true },
				OnClose:     func() { closed = true },
				Close:       true,
			})
		}
	}, 700, 300)

	// The toast is added during the first build; a frame shows it.
	tt.Frame()

	if !tt.HasText("已保存") {
		t.Fatalf("toast title not shown, texts: %q", tt.Texts())
	}
	if !tt.HasText("改动已写入磁盘") {
		t.Fatalf("toast description not shown, texts: %q", tt.Texts())
	}
	if !tt.HasText("撤销") {
		t.Fatalf("toast action not shown, texts: %q", tt.Texts())
	}

	if err := tt.Click("撤销"); err != nil {
		t.Fatalf("clicking the toast action: %v", err)
	}
	if !acted {
		t.Error("the toast action did not run OnAction")
	}
	if !closed {
		t.Error("the toast action did not run OnClose")
	}

	tt.Frame()
	if tt.HasText("已保存") {
		t.Errorf("the toast still shows after its action: %q", tt.Texts())
	}
}

func TestToastShapes(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	for _, shape := range []ToastShape{ToastCapsule, ToastCorner} {
		name := "capsule"
		if shape == ToastCorner {
			name = "corner"
		}
		t.Run(name, func(t *testing.T) {
			shown, acted := false, false
			tt := ui.NewTester(func(c *ui.Context) {
				ToastHost(c, ToastHostOptions{Shape: shape}, th)
				if !shown {
					shown = true
					ShowToast(c, ToastOptions{
						Title:       "已保存",
						Description: "改动已写入磁盘",
						Kind:        ToastSuccess,
						Action:      "撤销",
						OnAction:    func() { acted = true },
						Close:       true,
					})
				}
			}, 700, 300)
			tt.Frame()

			for _, want := range []string{"已保存", "改动已写入磁盘", "撤销", "Close"} {
				if !tt.HasText(want) {
					t.Fatalf("the %s toast lacks %q: %q", name, want, tt.Texts())
				}
			}

			if err := tt.Click("撤销"); err != nil {
				t.Fatalf("clicking the %s toast action: %v", name, err)
			}
			if !acted {
				t.Errorf("the %s toast action did not run OnAction", name)
			}
		})
	}
}

// The default toast must read as part of a light app rather than as a dark
// slab, and the inverse one must stay dark: this pins the Material You
// surface palette against a regression to the classic black snackbar.
func TestToastColorsReadOnTheirTheme(t *testing.T) {
	for _, dark := range []bool{false, true} {
		for _, color := range []ToastColor{ToastSurface, ToastInverse} {
			scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), dark)
			th := theme.NewTheme(scheme, theme.DensityStandard)

			shown := false
			tt := ui.NewTester(func(c *ui.Context) {
				ToastHost(c, ToastHostOptions{Placement: ToastBottomEnd, Color: color}, th)
				if !shown {
					shown = true
					ShowToast(c, ToastOptions{Title: "已保存", Kind: ToastSuccess})
				}
			}, 420, 120)
			tt.Frame()

			// Sample just inside the toast's fill, past the title's box.
			r, ok := tt.Find("已保存")
			if !ok {
				t.Fatal("the toast's title is not laid out")
			}
			img := tt.Image()
			x, y := int(r.X+r.W+6), int(r.Y+r.H/2)
			i := img.PixOffset(x, y)
			pixel := ui.RGB(img.Pix[i], img.Pix[i+1], img.Pix[i+2])
			lum := 0.2126*float64(pixel.R) + 0.7152*float64(pixel.G) + 0.0722*float64(pixel.B)

			wantLight := color == ToastSurface
			if dark {
				wantLight = !wantLight
			}
			if wantLight && lum < 160 {
				t.Errorf("dark=%v %v: the toast is dark (luminance %.0f), not a surface",
					dark, color, lum)
			}
			if !wantLight && lum > 110 {
				t.Errorf("dark=%v %v: the inverse toast is light (luminance %.0f)",
					dark, color, lum)
			}
		}
	}
}

func TestToastCloseButtonAndIconOverrides(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	shown, closed := false, false

	tt := ui.NewTester(func(c *ui.Context) {
		ToastHost(c, ToastHostOptions{}, th)

		if !shown {
			shown = true
			ShowToast(c, ToastOptions{
				Title:   "Upload failed",
				Kind:    ToastError,
				Icon:    symbols.Favorite,
				Close:   true,
				Timeout: -1, // until it is closed
				OnClose: func() { closed = true },
			})
		}
	}, 600, 300)

	tt.Frame()

	if !tt.HasText("Upload failed") {
		t.Fatalf("toast not shown, texts: %q", tt.Texts())
	}

	if err := tt.Click("Close"); err != nil {
		t.Fatalf("clicking the toast close button: %v", err)
	}
	if !closed {
		t.Error("the close button did not run OnClose")
	}

	tt.Frame()
	if tt.HasText("Upload failed") {
		t.Errorf("the toast still shows after Close: %q", tt.Texts())
	}
}
