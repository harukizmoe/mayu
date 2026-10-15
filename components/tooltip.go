package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// Tooltip attaches an M3 floating tooltip to an anchor element.
// Renders with InverseSurface background and InverseOnSurface text.
func Tooltip(c *ui.Context, anchor ui.Element, text string, t *theme.Theme) ui.Element {
	return ui.TooltipBase(c, anchor, func(tip ui.Element) {
		tip.Padding(4, 8).Radius(theme.CornerExtraSmall).
			Background(t.Scheme.InverseSurface).
			TextColor(t.Scheme.InverseOnSurface)

		tip.Children(func() {
			ui.Text(c, text).FontSize(theme.LabelSmall.Size).SingleLine()
		})
	})
}
