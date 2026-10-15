package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// Tabs creates a Material 3 tab list: one row of labels over a shared
// indicator line, the chosen tab's segment thicker and in the primary color.
// A click chooses a tab, as do the arrows, Home and End while one has the
// keyboard focus.
func Tabs(c *ui.Context, selected *int, labels []string, t *theme.Theme) ui.Element {
	n := len(labels)
	if n == 0 {
		return ui.Box(c)
	}

	parts := ui.TabsBase(c, selected, n)
	parts.List.FillWidth().Height(48).AlignItems(ui.Stretch)

	parts.List.Children(func() {
		for i, name := range labels {
			idx := i
			chosen := *selected == idx

			tab := parts.Tab(idx).Height(48).Padding(0, 12).Grow(1).Focusable()

			fg := t.Scheme.OnSurfaceVariant
			if chosen {
				fg = t.Scheme.Primary
			}
			tab.Background(ApplyStateLayer(ui.Transparent, fg,
				tab.Hovered(), tab.Focused(), tab.Pressed())).TextColor(fg)

			// The indicator line: every tab contributes its own segment, so
			// the row reads as one line with the chosen segment raised.
			tab.Draw(func(p *ui.Painter, r ui.Rect) {
				if chosen {
					p.Fill(ui.Rect{X: r.X, Y: r.Y + r.H - 3, W: r.W, H: 3}, t.Scheme.Primary, 1.5)
					return
				}
				p.Fill(ui.Rect{X: r.X, Y: r.Y + r.H - 1, W: r.W, H: 1}, t.Scheme.OutlineVariant, 0)
			})

			tab.Children(func() {
				// Lifted off the line so a descender does not touch it.
				ui.Column(c).FillHeight().Justify(ui.Center).Padding(0, 0, 6, 0).Children(func() {
					label := ui.Text(c, name).FontSize(theme.LabelLarge.Size).SingleLine()
					if chosen {
						label.Bold()
					}
				})
			})
		}
	})

	return parts.List
}
