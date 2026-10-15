package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// AlertDialogOptions configures a Material 3 basic alert dialog.
type AlertDialogOptions struct {
	Title       string
	Body        string
	ConfirmText string
	OnConfirm   func()
	DismissText string
	OnDismiss   func()
}

// AlertDialog renders an M3 Alert Dialog when *open is true.
func AlertDialog(c *ui.Context, open *bool, opts AlertDialogOptions, t *theme.Theme) ui.Element {
	return ui.DialogBase(c, open, func(backdrop ui.Element, panel ui.Element) {
		// Backdrop darkening with Scrim
		backdrop.Background(t.Scheme.Scrim.Alpha(0.35))

		// Panel styling: 28dp radius, SurfaceContainerHigh, elevation shadow
		panel.Width(340).Radius(theme.CornerExtraLarge).
			Padding(24).
			Background(t.Scheme.SurfaceContainerHigh).
			Shadow(0, 6, 18, 6, t.Scheme.Shadow.Alpha(0.25))

		ui.Column(c).Gap(16).FillWidth().Children(func() {
			// Title
			if opts.Title != "" {
				ui.Text(c, opts.Title).FontSize(theme.HeadlineSmall.Size).Bold().TextColor(t.Scheme.OnSurface)
			}

			// Body text
			if opts.Body != "" {
				ui.Text(c, opts.Body).FontSize(theme.BodyMedium.Size).TextColor(t.Scheme.OnSurfaceVariant)
			}

			// Actions row (end-aligned)
			ui.Row(c).Justify(ui.End).Gap(8).FillWidth().Children(func() {
				if opts.DismissText != "" {
					NewButton(c, opts.DismissText, VariantText, t).
						OnClick(func() {
							*open = false
							if opts.OnDismiss != nil {
								opts.OnDismiss()
							}
						})
				}

				if opts.ConfirmText != "" {
					NewButton(c, opts.ConfirmText, VariantText, t).
						OnClick(func() {
							*open = false
							if opts.OnConfirm != nil {
								opts.OnConfirm()
							}
						})
				}
			})
		})
	})
}
