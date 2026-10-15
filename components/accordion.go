package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// Accordion creates an M3 collapsible panel with a clickable header and expandable body.
func Accordion(c *ui.Context, open *bool, title string, content func(), t *theme.Theme) ui.Element {
	col := ui.Column(c).Radius(theme.CornerMedium).Background(t.Scheme.SurfaceContainerLow).
		Border(1, t.Scheme.OutlineVariant).Clip().FillWidth()

	isOpen := *open

	col.Children(func() {
		// Header row
		header := ui.ButtonBase(c).Height(48).Padding(0, 16).AlignItems(ui.Center).Justify(ui.SpaceBetween).FillWidth()
		if header.Clicked() && !header.IsDisabled() {
			*open = !*open
		}

		bg := ApplyStateLayer(ui.Transparent, t.Scheme.OnSurface, header.Hovered(), header.Focused(), header.Pressed())
		header.Background(bg)

		header.Children(func() {
			ui.Text(c, title).FontSize(theme.TitleMedium.Size).Bold().TextColor(t.Scheme.OnSurface).SingleLine()

			// Trailing indicator: Close if open, Add if closed
			icon := symbols.Add
			if isOpen {
				icon = symbols.Close
			}
			ui.Icon(c, icon).Size(18, 18).TextColor(t.Scheme.OnSurfaceVariant)
		})

		// Body content
		if isOpen && content != nil {
			ui.Column(c).Padding(4, 16, 16, 16).FillWidth().Children(content)
		}
	})

	return col
}
