package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// SearchBarBuilder provides a builder for M3 Search Bars.
type SearchBarBuilder struct {
	ui.Element
	c           *ui.Context
	query       *string
	placeholder string
	input       ui.Element
	trailing    func()
	scheme      theme.ColorScheme
	density     theme.Density
}

// NewSearchBar creates a Material 3 Search Bar (56dp height pill).
func NewSearchBar(c *ui.Context, query *string, placeholder string, t *theme.Theme) *SearchBarBuilder {
	row := ui.Row(c).Height(56).Radius(28).Padding(0, 16).Gap(12).AlignItems(ui.Center).FillWidth().
		Background(t.Scheme.SurfaceContainerHigh).
		Shadow(0, 1, 3, 1, t.Scheme.Shadow.Alpha(0.15))

	builder := &SearchBarBuilder{
		Element:     row,
		c:           c,
		query:       query,
		placeholder: placeholder,
		scheme:      t.Scheme,
		density:     t.Density,
	}

	row.Children(func() {
		// Leading search icon
		ui.Icon(c, symbols.Search).Size(22, 22).TextColor(t.Scheme.OnSurfaceVariant)

		// Middle editable text input, named for assistive technology: without
		// a name a screen reader reads the box as an unnamed text field.
		input := ui.TextInputBase(c, query).
			Grow(1).
			TextColor(t.Scheme.OnSurface)

		if placeholder != "" {
			input.Placeholder(placeholder)
			input.Label(placeholder)
		}
		builder.input = input

		// Trailing clear button when query has text
		if *query != "" {
			clearBtn := ui.ButtonBase(c).Size(32, 32).Radius(16).Center().Focusable()
			if clearBtn.Clicked() {
				*query = ""
			}
			bg := ApplyStateLayer(ui.Transparent, t.Scheme.OnSurface, clearBtn.Hovered(), clearBtn.Focused(), clearBtn.Pressed())
			clearBtn.Background(bg)
			clearBtn.Children(func() {
				ui.Icon(c, symbols.Close).Size(18, 18).TextColor(t.Scheme.OnSurfaceVariant)
			})
		}
	})

	return builder
}
