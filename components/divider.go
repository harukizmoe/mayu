package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// HorizontalDivider creates a 1dp horizontal separator line in OutlineVariant.
func HorizontalDivider(c *ui.Context, t *theme.Theme) ui.Element {
	return ui.Box(c).Height(1).FillWidth().Background(t.Scheme.OutlineVariant)
}

// VerticalDivider creates a 1dp vertical separator line in OutlineVariant.
func VerticalDivider(c *ui.Context, t *theme.Theme) ui.Element {
	return ui.Box(c).Width(1).FillHeight().Background(t.Scheme.OutlineVariant)
}
