package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// MenuItemOptions configures a single option in an M3 Menu.
type MenuItemOptions struct {
	Icon     *ui.SVG
	Shortcut string
	Disabled bool
	OnClick  func()
}

// Menu renders an M3 Dropdown Menu popup anchored to an element when *open is true.
func Menu(c *ui.Context, anchor ui.Element, open *bool, items func(m *MenuPanel), t *theme.Theme) ui.Element {
	return ui.PopoverBase(c, anchor, open, func(panel ui.Element) {
		panel.MinWidth(160).MaxWidth(280).Radius(theme.CornerSmall).Padding(8, 0).
			Background(t.Scheme.SurfaceContainerHigh).
			Shadow(0, 3, 9, 3, t.Scheme.Shadow.Alpha(0.22))

		mp := &MenuPanel{c: c, open: open, theme: t}
		panel.Children(func() {
			ui.Column(c).FillWidth().Children(func() {
				items(mp)
			})
		})
	})
}

// MenuPanel provides methods to populate items inside a Menu.
type MenuPanel struct {
	c     *ui.Context
	open  *bool
	theme *theme.Theme
}

// Item adds a clickable row item to the menu.
func (m *MenuPanel) Item(label string, opts MenuItemOptions) {
	c := m.c
	t := m.theme

	row := ui.ButtonBase(c).Height(40).Padding(0, 16).Gap(12).AlignItems(ui.Center).FillWidth()
	if opts.Disabled {
		row.Disabled(true)
	}

	if row.Clicked() && !row.IsDisabled() {
		*m.open = false
		if opts.OnClick != nil {
			opts.OnClick()
		}
	}

	// State layer
	fg := t.Scheme.OnSurface
	if opts.Disabled {
		fg = t.Scheme.OnSurface.Alpha(DisabledContentAlpha)
	}
	bg := ApplyStateLayer(ui.Transparent, t.Scheme.OnSurface, row.Hovered(), row.Focused(), row.Pressed())

	row.Background(bg).TextColor(fg)
	row.Children(func() {
		if opts.Icon != nil {
			ui.Icon(c, opts.Icon).Size(20, 20).TextColor(t.Scheme.OnSurfaceVariant)
		}

		ui.Text(c, label).FontSize(theme.BodyLarge.Size).TextColor(fg).Grow(1).SingleLine()

		if opts.Shortcut != "" {
			ui.Text(c, opts.Shortcut).FontSize(theme.LabelMedium.Size).TextColor(t.Scheme.OnSurfaceVariant).SingleLine()
		}
	})
}

// Divider adds a hairline horizontal separator inside the menu.
func (m *MenuPanel) Divider() {
	HorizontalDivider(m.c, m.theme)
}
