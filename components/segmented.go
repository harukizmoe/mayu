package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// SegmentOption defines a single segment in a SegmentedButton.
type SegmentOption struct {
	Label string
	Icon  *ui.SVG
}

// SegmentedButton creates a connected M3 Segmented Button group.
func SegmentedButton(c *ui.Context, selected *int, options []SegmentOption, t *theme.Theme) ui.Element {
	n := len(options)
	if n == 0 {
		return ui.Box(c)
	}

	parts := ui.SegmentedBase(c, selected, n)
	row := parts.Track.Height(40).Radius(20).Border(1, t.Scheme.Outline).Clip().AlignItems(ui.Stretch)

	parts.Track.Children(func() {
		for i, opt := range options {
			idx := i
			isChosen := *selected == idx
			seg := parts.Segment(idx).Height(40).Padding(0, 16).Gap(8).Center().Focusable()

			// Segment background & foreground colors
			var bg ui.Color
			var fg ui.Color

			if isChosen {
				base := t.Scheme.SecondaryContainer
				fg = t.Scheme.OnSecondaryContainer
				bg = ApplyStateLayer(base, fg, seg.Hovered(), seg.Focused(), seg.Pressed())
			} else {
				fg = t.Scheme.OnSurface
				bg = ApplyStateLayer(ui.Transparent, fg, seg.Hovered(), seg.Focused(), seg.Pressed())
			}

			seg.Background(bg).TextColor(fg)

			// Separator line between segments (except first)
			if idx > 0 {
				seg.BorderWidth(0, 0, 0, 1).BorderColor(t.Scheme.Outline)
			}

			seg.Children(func() {
				// Show checkmark icon if chosen
				if isChosen {
					ui.Icon(c, symbols.Check).Size(18, 18).TextColor(fg)
				} else if opt.Icon != nil {
					ui.Icon(c, opt.Icon).Size(18, 18).TextColor(fg)
				}

				if opt.Label != "" {
					ui.Text(c, opt.Label).FontSize(theme.LabelLarge.Size).Bold().SingleLine()
				}
			})
		}
	})

	return row
}
