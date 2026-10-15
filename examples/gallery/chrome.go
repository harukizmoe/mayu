package main

import (
	"fmt"
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// matches reports whether the search box admits a section title. An empty box
// admits every section, so a page shows whole.
func (a *galleryApp) matches(title string) bool {
	q := strings.ToLower(strings.TrimSpace(a.searchQuery))
	return q == "" || strings.Contains(strings.ToLower(title), q)
}

// section renders one titled section of a page: its title, its content, and
// the divider that closes it. A title the search box does not admit is left
// out whole, divider and all. It reports whether the section was shown.
func (a *galleryApp) section(c *ui.Context, title string, body func()) bool {
	if !a.matches(title) {
		return false
	}
	ui.Column(c).Gap(12).Children(func() {
		ui.Text(c, title).FontSize(theme.TitleMedium.Size).Bold()
		body()
	})
	mayu.HorizontalDivider(c)
	return true
}

// emptyState says so when the search box admits no section of a page, rather
// than leaving the page blank.
func (a *galleryApp) emptyState(c *ui.Context, shown int, t *theme.Theme) {
	if shown > 0 || strings.TrimSpace(a.searchQuery) == "" {
		return
	}
	ui.Row(c).Padding(24).AlignItems(ui.Center).Gap(12).Children(func() {
		ui.Icon(c, symbols.Search).Size(20, 20).TextColor(t.Scheme.OnSurfaceVariant)
		ui.Text(c, fmt.Sprintf("No component matches %q", a.searchQuery)).
			TextColor(t.Scheme.OnSurfaceVariant)
	})
}

// header is the top of every page: which page this is, the box that filters
// its sections, and the two switches that restyle the whole window.
func (a *galleryApp) header(c *ui.Context, t *theme.Theme) {
	ui.Row(c).AlignItems(ui.Center).Gap(16).FillWidth().Children(func() {
		ui.Column(c).Gap(2).Shrink(0).Children(func() {
			ui.Text(c, "Mayu Gallery").FontSize(theme.HeadlineSmall.Size).Bold()
			ui.Text(c, a.pageTitle()).FontSize(theme.BodyMedium.Size).
				TextColor(t.Scheme.OnSurfaceVariant)
		})

		ui.Spacer(c)

		mayu.SearchBar(c, &a.searchQuery, "Filter this page's components...").Width(280)

		density := "Standard"
		switch a.density {
		case mayu.DensityCompact:
			density = "Compact"
		case mayu.DensityComfortable:
			density = "Comfortable"
		}
		mayu.OutlinedButton(c, "Density: "+density).
			OnClick(func() {
				switch a.density {
				case mayu.DensityStandard:
					a.density = mayu.DensityCompact
				case mayu.DensityCompact:
					a.density = mayu.DensityComfortable
				default:
					a.density = mayu.DensityStandard
				}
			})

		modeIcon, modeLabel := symbols.DarkMode, "Dark Mode"
		if a.dark {
			modeIcon, modeLabel = symbols.LightMode, "Light Mode"
		}
		mayu.FilledTonalButton(c, modeLabel).
			Icon(modeIcon).
			OnClick(func() { a.dark = !a.dark })
	})
}

// statusFooter reports what the last interaction was, so a click that changed
// nothing visible is still legible.
func (a *galleryApp) statusFooter(c *ui.Context, t *theme.Theme) {
	ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
		ui.Text(c, fmt.Sprintf("Clicks: %d", a.clicks)).Bold()
		ui.Text(c, "•").TextColor(t.Scheme.Outline)
		ui.Text(c, fmt.Sprintf("Page=%s, Segment=%d, Tier=%s", a.pageTitle(), a.segIdx, a.radioChoice)).
			TextColor(t.Scheme.OnSurfaceVariant)
		ui.Text(c, "•").TextColor(t.Scheme.Outline)
		ui.Text(c, fmt.Sprintf("Status: %s", a.lastAction)).TextColor(t.Scheme.Primary)
	})
}

// toastHost installs the window's toasts with the shapes and colors the
// Overlays page lets you pick.
func (a *galleryApp) toastHost(c *ui.Context, t *theme.Theme) {
	shape := mayu.ToastCapsule
	if a.toastShape == 1 {
		shape = mayu.ToastCorner
	}
	color := mayu.ToastSurface
	if a.toastColor == 1 {
		color = mayu.ToastInverse
	}
	mayu.ToastHostWith(c, mayu.ToastHostOptions{
		Placement: mayu.ToastBottomEnd,
		Shape:     shape,
		Color:     color,
	})
}
