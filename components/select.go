package components

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// SelectOption is one choice of a Select: the value it stores and the label it
// shows.
type SelectOption[T comparable] struct {
	Value T
	Label string
}

// Select creates a Material 3 select: an outlined field showing the chosen
// option, over a popup listing them all. The popup's height is capped, and the
// option the highlight moves to scrolls into view.
func Select[T comparable](c *ui.Context, selected *T, label string, options []SelectOption[T], t *theme.Theme) ui.Element {
	parts := ui.SelectBase(c, selected)
	trigger := parts.Trigger

	trigger.MinHeight(56).Padding(0, 16).Gap(12).AlignItems(ui.Center).FillWidth().
		Radius(theme.CornerSmall).Border(1, t.Scheme.Outline)

	trigger.Children(func() {
		ui.Column(c).Grow(1).Justify(ui.Center).Children(func() {
			if label != "" {
				ui.Text(c, label).FontSize(11).TextColor(t.Scheme.Primary).SingleLine()
			}
			ui.Text(c, selectLabel(*selected, options)).
				FontSize(theme.BodyLarge.Size).TextColor(t.Scheme.OnSurface).SingleLine()
		})
		ui.Icon(c, symbols.ExpandMore).Size(20, 20).TextColor(t.Scheme.OnSurfaceVariant)
	})

	parts.Popup(func(panel ui.Element) {
		panel.MinWidth(180).MaxWidth(360).MaxHeight(280).Radius(theme.CornerSmall).
			Padding(8).Background(t.Scheme.SurfaceContainerHigh).
			Shadow(0, 3, 9, 3, t.Scheme.Shadow.Alpha(0.22))

		for _, opt := range options {
			item := parts.Item(opt.Value).MinHeight(40).Padding(0, 12).Gap(12).
				AlignItems(ui.Center).Radius(theme.CornerExtraSmall).Focusable()

			chosen := *selected == opt.Value
			switch {
			case item.Highlighted():
				item.Background(t.Scheme.SecondaryContainer)
			case chosen:
				item.Background(ApplyStateLayer(ui.Transparent, t.Scheme.OnSurface,
					item.Hovered(), item.Focused(), item.Pressed()))
			default:
				item.Background(ApplyStateLayer(ui.Transparent, t.Scheme.OnSurface,
					item.Hovered(), item.Focused(), item.Pressed()))
			}

			fg := t.Scheme.OnSurface
			if chosen || item.Highlighted() {
				fg = t.Scheme.OnSecondaryContainer
			}
			item.TextColor(fg)

			item.Children(func() {
				// The chosen option shows where it is, so the list reads at a
				// glance; the others keep its room so the labels line up.
				if chosen {
					ui.Icon(c, symbols.Check).Size(18, 18).TextColor(fg).Shrink(0)
				} else {
					ui.Box(c).Width(18).Shrink(0)
				}
				ui.Text(c, opt.Label).FontSize(theme.BodyLarge.Size).SingleLine()
			})
		}
	})

	return trigger
}

// selectLabel is the label of the option a select holds, or the value itself
// when no option carries it.
func selectLabel[T comparable](value T, options []SelectOption[T]) string {
	for _, opt := range options {
		if opt.Value == value {
			return opt.Label
		}
	}
	return fmt.Sprint(value)
}
