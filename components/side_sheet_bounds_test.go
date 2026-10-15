package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// 抽屉的底栏必须留在面板里面。这一条只有量坐标才发现得了：AX 树里三颗按钮都在，
// 但内容一长，Column 就把 footer 顶出面板底部，主操作在屏幕上根本看不见。
func TestSideSheet_FooterStaysInsideThePanel(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	const winW, winH = 1200, 700
	open := true

	tt := ui.NewTester(func(c *ui.Context) {
		SideSheet(c, &open, SideSheetOptions{
			Title:    "Wallpaper Settings",
			Subtitle: "DP-1 (Main)",
			Actions: func() {
				NewButton(c, "Apply Now", VariantFilled, th)
				ui.Row(c).Gap(8).Children(func() {
					NewButton(c, "Save", VariantFilledTonal, th)
					NewButton(c, "Queue", VariantFilledTonal, th)
				})
			},
		}, th, func() {
			// 足够长的内容，把 Column 撑到超过面板高度
			for range 40 {
				ui.Text(c, "Scale 1.20x")
			}
		})
	}, winW, winH)

	apply, ok := tt.Find("Apply Now")
	if !ok {
		t.Fatal("the sheet's primary action is missing from the tree")
	}
	t.Logf("面板窗口 %dx%d，主操作按钮 rect=%+v", winW, winH, apply)

	if apply.Y+apply.H > float32(winH) {
		t.Fatalf("the primary action is pushed out of the window: bottom=%.1f > window height %d",
			apply.Y+apply.H, winH)
	}
	if apply.H <= 0 {
		t.Fatalf("the primary action has no height: %+v", apply)
	}

	// 内容再长也不能把 footer 顶出去：标题同理要留在顶部。
	title, ok := tt.Find("Wallpaper Settings")
	if !ok {
		t.Fatal("the sheet title is missing from the tree")
	}
	if title.Y < 0 {
		t.Fatalf("the title is pushed above the panel: y=%.1f", title.Y)
	}
	if title.Y+title.H > apply.Y {
		t.Fatalf("the title overlaps the footer: title bottom=%.1f, footer top=%.1f",
			title.Y+title.H, apply.Y)
	}
}
