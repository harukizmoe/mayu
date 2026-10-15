package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// SnackbarOptions configures an M3 Snackbar banner.
type SnackbarOptions struct {
	ActionText string
	OnAction   func()
}

// Snackbar renders an M3 Snackbar floating feedback pill.
func Snackbar(c *ui.Context, message string, opts SnackbarOptions, t *theme.Theme) ui.Element {
	bar := ui.Row(c).MinHeight(48).Radius(theme.CornerExtraSmall).Padding(8, 16).Gap(16).
		AlignItems(ui.Center).Justify(ui.SpaceBetween).
		Background(t.Scheme.InverseSurface).
		Shadow(0, 3, 9, 3, t.Scheme.Shadow.Alpha(0.20))

	bar.Children(func() {
		// Message Text
		ui.Text(c, message).FontSize(theme.BodyMedium.Size).TextColor(t.Scheme.InverseOnSurface).Grow(1)

		// Action Button
		if opts.ActionText != "" {
			btn := ui.ButtonBase(c).Height(36).Padding(0, 8).Center().Focusable()
			if btn.Clicked() && !btn.IsDisabled() && opts.OnAction != nil {
				opts.OnAction()
			}

			fg := t.Scheme.InversePrimary
			bg := ApplyStateLayer(ui.Transparent, fg, btn.Hovered(), btn.Focused(), btn.Pressed())

			btn.Background(bg).TextColor(fg)
			btn.Children(func() {
				ui.Text(c, opts.ActionText).FontSize(theme.LabelLarge.Size).Bold().SingleLine()
			})
		}
	})

	return bar
}
