package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// SideSheetOptions configures a Material 3 modal side sheet.
type SideSheetOptions struct {
	// Title is the sheet's headline. Empty for a sheet without a header.
	Title string
	// Subtitle sits under the title, in OnSurfaceVariant.
	Subtitle string
	// Width of the panel in DIPs. Zero means 360, the M3 modal side sheet width.
	Width float32
	// Actions are laid out in a footer pinned below the content, under a
	// divider. Nil for a sheet whose content carries its own actions.
	Actions func()
}

// SideSheet renders a Material 3 modal side sheet anchored to the right edge
// of the window while *open is true.
//
// It is the side-anchored sibling of AlertDialog and reuses the same overlay:
// the content behind it is dimmed and blocked, the sheet takes focus, and it
// closes on Escape or a click on the scrim. The difference is placement — a
// full-height panel against the edge instead of a centred box — which is what
// suits a surface the user reads alongside the content it acts on, and what
// keeps a long list of controls usable without a scrolling dialog.
func SideSheet(c *ui.Context, open *bool, opts SideSheetOptions, t *theme.Theme, content func()) ui.Element {
	width := opts.Width
	if width <= 0 {
		width = 360
	}
	sc := t.Scheme

	return ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
		backdrop.Background(sc.Scrim.Alpha(0.35))

		// Absolute keeps the panel out of the backdrop's centring, and pinning
		// it to the right edge makes it a sheet rather than a dialog. Only the
		// left corners are rounded: the right ones are off-screen.
		panel.Absolute().Top(0).Right(0).Bottom(0).Width(width).
			Background(sc.SurfaceContainerLow).
			Radius(theme.CornerExtraLarge, 0, 0, theme.CornerExtraLarge).
			BorderWidth(0, 0, 0, 1).
			BorderColor(sc.OutlineVariant).
			Shadow(-6, 0, 20, 0, sc.Shadow.Alpha(0.3))

		ui.Column(c).Fill().Children(func() {
			if opts.Title != "" {
				ui.Column(c).FillWidth().Padding(24, 20).Gap(4).Children(func() {
					ui.Text(c, opts.Title).
						FontSize(theme.TitleLarge.Size).Bold().TextColor(sc.OnSurface)
					if opts.Subtitle != "" {
						ui.Text(c, opts.Subtitle).
							FontSize(theme.BodyMedium.Size).TextColor(sc.OnSurfaceVariant)
					}
				})
				HorizontalDivider(c, t)
			}

			ui.Scroll(c).Grow(1).FillWidth().Children(func() {
				ui.Column(c).FillWidth().Padding(24).Gap(16).Children(content)
			})

			if opts.Actions != nil {
				HorizontalDivider(c, t)
				ui.Column(c).FillWidth().Padding(24, 16).Gap(8).Children(opts.Actions)
			}
		})
	})
}
