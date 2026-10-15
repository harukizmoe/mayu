package components

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// LinearProgress creates a Material 3 linear progress bar.
// progress is a fraction between 0.0 and 1.0.
func LinearProgress(c *ui.Context, progress float32, t *theme.Theme) ui.Element {
	progress = min(max(progress, 0), 1)

	track := ui.Box(c).Height(4).Radius(2).FillWidth().Background(t.Scheme.SecondaryContainer).Clip()
	track.Children(func() {
		ui.Box(c).FillHeight().WidthPercent(progress * 100).Radius(2).Background(t.Scheme.Primary)
	})
	return track
}

// Badge creates a Material 3 small notification badge.
// If count <= 0, renders a small 6x6 dot. If count > 0, renders a pill with the count.
func Badge(c *ui.Context, count int, t *theme.Theme) ui.Element {
	if count <= 0 {
		return ui.Box(c).Size(6, 6).Radius(3).Background(t.Scheme.Error)
	}

	label := fmt.Sprintf("%d", count)
	if count > 99 {
		label = "99+"
	}

	pill := ui.Row(c).Height(16).Radius(8).Padding(0, 4).Center().Background(t.Scheme.Error)
	pill.Children(func() {
		ui.Text(c, label).FontSize(10).TextColor(t.Scheme.OnError).Bold().SingleLine()
	})
	return pill
}
