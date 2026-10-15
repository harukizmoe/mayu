package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// NavDestination defines a destination entry in an M3 Navigation Rail.
type NavDestination struct {
	Icon  *ui.SVG
	Label string
	Badge int
}

// NavigationRailOptions configures an M3 NavigationRail.
type NavigationRailOptions struct {
	Header       func()
	Footer       func()
	Destinations []NavDestination
	Selected     *int
	OnSelect     func(index int)
}

// NavigationRail constructs a Material 3 vertical desktop navigation rail (80dp width).
func NavigationRail(c *ui.Context, opts NavigationRailOptions, t *theme.Theme) ui.Element {
	rail := ui.Column(c).Width(80).FillHeight().Padding(16, 0).Justify(ui.SpaceBetween).AlignItems(ui.Center).
		Background(t.Scheme.Surface).
		BorderWidth(0, 1, 0, 0).
		BorderColor(t.Scheme.OutlineVariant)

	rail.Children(func() {
		// Top section: Optional Header + Destinations
		ui.Column(c).FillWidth().AlignItems(ui.Center).Gap(16).Children(func() {
			if opts.Header != nil {
				ui.Box(c).Center().Padding(0, 0, 8, 0).Children(opts.Header)
			}

			// Destination Items
			for i, dest := range opts.Destinations {
				idx := i
				isSelected := opts.Selected != nil && *opts.Selected == idx

				item := ui.Column(c).Width(68).AlignItems(ui.Center).Gap(4).Focusable()
				if item.Clicked() && !item.IsDisabled() {
					if opts.Selected != nil {
						*opts.Selected = idx
					}
					if opts.OnSelect != nil {
						opts.OnSelect(idx)
					}
				}

				// Active pill indicator container (56x32, radius 16)
				indicator := ui.Box(c).Size(56, 32).Radius(16).Center().Shrink(0)

				var indBg ui.Color
				var iconFg ui.Color
				var labelFg ui.Color

				if isSelected {
					base := t.Scheme.SecondaryContainer
					iconFg = t.Scheme.OnSecondaryContainer
					labelFg = t.Scheme.OnSurface
					indBg = ApplyStateLayer(base, iconFg, item.Hovered(), item.Focused(), item.Pressed())
				} else {
					iconFg = t.Scheme.OnSurfaceVariant
					labelFg = t.Scheme.OnSurfaceVariant
					indBg = ApplyStateLayer(ui.Transparent, iconFg, item.Hovered(), item.Focused(), item.Pressed())
				}

				indicator.Background(indBg)
				indicator.Children(func() {
					if dest.Icon != nil {
						ui.Icon(c, dest.Icon).Size(22, 22).TextColor(iconFg)
					}
				})

				item.Children(func() {
					indicatorNode := indicator
					_ = indicatorNode

					if dest.Label != "" {
						label := ui.Text(c, dest.Label).FontSize(11).TextColor(labelFg).SingleLine()
						if isSelected {
							label.Bold()
						}
					}
				})
			}
		})

		// Bottom section: Optional Footer
		if opts.Footer != nil {
			ui.Box(c).Center().Padding(8, 0, 0, 0).Children(opts.Footer)
		}
	})

	return rail
}
