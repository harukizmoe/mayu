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
