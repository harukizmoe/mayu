package components

import (
	"strings"
	"unicode/utf8"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// AvatarBuilder provides a fluent builder for M3 Avatars.
type AvatarBuilder struct {
	ui.Element
	c       *ui.Context
	size    float32
	online  bool
	scheme  theme.ColorScheme
	density theme.Density
}

// Avatar creates a circular Avatar showing an icon.
func Avatar(c *ui.Context, icon *ui.SVG, t *theme.Theme) *AvatarBuilder {
	box := ui.Box(c).Size(40, 40).Radius(20).Center().Background(t.Scheme.PrimaryContainer).Clip()

	builder := &AvatarBuilder{
		Element: box,
		c:       c,
		size:    40,
		scheme:  t.Scheme,
		density: t.Density,
	}

	box.Children(func() {
		if icon != nil {
			ui.Icon(c, icon).Size(20, 20).TextColor(t.Scheme.OnPrimaryContainer)
		}
	})

	return builder
}

// InitialsAvatar creates an M3 Avatar displaying 1-2 initials extracted from name.
func InitialsAvatar(c *ui.Context, name string, t *theme.Theme) *AvatarBuilder {
	initials := extractInitials(name)
	box := ui.Box(c).Size(40, 40).Radius(20).Center().Background(t.Scheme.PrimaryContainer).Clip()

	builder := &AvatarBuilder{
		Element: box,
		c:       c,
		size:    40,
		scheme:  t.Scheme,
		density: t.Density,
	}

	box.Children(func() {
		ui.Text(c, initials).FontSize(14).Bold().TextColor(t.Scheme.OnPrimaryContainer).SingleLine()
	})

	return builder
}

// Size sets a custom size in DIPs.
func (b *AvatarBuilder) Size(s float32) *AvatarBuilder {
	b.size = s
	b.Element.Size(s, s).Radius(s / 2)
	return b
}

func extractInitials(name string) string {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "?"
	}
	if len(parts) == 1 {
		r, _ := utf8.DecodeRuneInString(parts[0])
		return strings.ToUpper(string(r))
	}
	r1, _ := utf8.DecodeRuneInString(parts[0])
	r2, _ := utf8.DecodeRuneInString(parts[len(parts)-1])
	return strings.ToUpper(string(r1) + string(r2))
}
