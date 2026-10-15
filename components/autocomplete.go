package components

import (
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// AutocompleteBuilder is a Material 3 autocomplete: the field, and what the
// user chose from it.
type AutocompleteBuilder struct {
	ui.Element
	c      *ui.Context
	text   *string
	chosen string
	picked bool
}

// NewAutocomplete creates a Material 3 autocomplete: an outlined field whose
// text filters the suggestions listed under it, those starting with the text
// first. A click or Enter takes a suggestion into the field, which Chosen
// reports.
func NewAutocomplete(c *ui.Context, text *string, label string, suggestions []string, t *theme.Theme) *AutocompleteBuilder {
	parts := ui.ComboboxBase(c, text)
	input := parts.Input

	builder := &AutocompleteBuilder{Element: input, c: c, text: text}

	input.MinHeight(56).Padding(8, 16).FillWidth().
		Radius(theme.CornerSmall).Border(1, t.Scheme.Outline).
		TextColor(t.Scheme.OnSurface)
	if label != "" {
		input.Placeholder(label).Label(label)
	}

	// The suggestions the text filters, or all of them when the popup was
	// opened by a click rather than by typing.
	shown := suggestions
	if parts.Filtering() {
		shown = matching(suggestions, *text)
	}

	parts.Popup(func(panel ui.Element) {
		panel.MinWidth(180).MaxWidth(360).MaxHeight(280).Radius(theme.CornerSmall).
			Padding(8).Background(t.Scheme.SurfaceContainerHigh).
			Shadow(0, 3, 9, 3, t.Scheme.Shadow.Alpha(0.22))

		for _, opt := range shown {
			item := parts.Item(opt).MinHeight(40).Padding(0, 12).
				AlignItems(ui.Center).Radius(theme.CornerExtraSmall)

			fg, bg := t.Scheme.OnSurface, ui.Color{}
			switch {
			case item.Highlighted():
				fg, bg = t.Scheme.OnSecondaryContainer, t.Scheme.SecondaryContainer
			case opt == *text:
				bg = t.Scheme.SurfaceContainerHighest
			}
			item.Background(bg).TextColor(fg)

			item.Children(func() {
				ui.Text(c, opt).FontSize(theme.BodyLarge.Size).SingleLine()
			})
		}
	})

	// The choice is read after Popup, which takes the options' input as it
	// returns, so it can be applied in the same build pass: without this the
	// field would keep what was typed rather than the suggestion chosen.
	if v, ok := parts.Chosen(); ok {
		*text = v
		builder.chosen, builder.picked = v, true
	}

	return builder
}

// Chosen returns the suggestion taken in this build pass, and whether one was.
func (b *AutocompleteBuilder) Chosen() (string, bool) {
	return b.chosen, b.picked
}

// matching returns the suggestions containing query, ignoring case, those
// starting with it first: all of them for a blank query.
func matching(options []string, query string) []string {
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return options
	}
	var starts, contains []string
	for _, o := range options {
		switch l := strings.ToLower(o); {
		case strings.HasPrefix(l, q):
			starts = append(starts, o)
		case strings.Contains(l, q):
			contains = append(contains, o)
		}
	}
	return append(starts, contains...)
}
