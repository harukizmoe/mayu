package components

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// ListItemOptions configures an M3 List Item.
type ListItemOptions struct {
	Leading        func()
	SupportingText string
	Trailing       func()
	OnClick        func()
}

// ListItemBuilder provides a fluent builder for list items.
type ListItemBuilder struct {
	ui.Element
	c              *ui.Context
	headline       string
	supportingText string
	leading        func()
	trailing       func()
	onClick        func()
	scheme         theme.ColorScheme
	density        theme.Density
}

// NewListItem creates a Material 3 List Item.
func NewListItem(c *ui.Context, headline string, opts ListItemOptions, t *theme.Theme) *ListItemBuilder {
	row := ui.Row(c).FillWidth().MinHeight(56).Padding(12, 16).Gap(16).AlignItems(ui.Center).Radius(theme.CornerSmall)

	builder := &ListItemBuilder{
		Element:        row,
		c:              c,
		headline:       headline,
		supportingText: opts.SupportingText,
		leading:        opts.Leading,
		trailing:       opts.Trailing,
		onClick:        opts.OnClick,
		scheme:         t.Scheme,
		density:        t.Density,
	}

	if opts.OnClick != nil {
		row.Focusable()
		if row.Clicked() && !row.IsDisabled() {
			opts.OnClick()
		}
	}

	builder.style()

	row.Children(func() {
		// 1. Leading Slot (Avatar, Icon, Checkbox)
		if builder.leading != nil {
			ui.Box(c).Shrink(0).Children(builder.leading)
		}

		// 2. Middle Content Slot (Headline + SupportingText)
		ui.Column(c).Grow(1).Justify(ui.Center).Gap(2).Children(func() {
			if builder.headline != "" {
				ui.Text(c, builder.headline).FontSize(theme.BodyLarge.Size).TextColor(builder.scheme.OnSurface).SingleLine()
			}
			if builder.supportingText != "" {
				ui.Text(c, builder.supportingText).FontSize(theme.BodyMedium.Size).TextColor(builder.scheme.OnSurfaceVariant).SingleLine()
			}
		})

		// 3. Trailing Slot (Switch, Badge, Icon, Timestamp)
		if builder.trailing != nil {
			ui.Box(c).Shrink(0).Children(builder.trailing)
		}
	})

	return builder
}

func (b *ListItemBuilder) style() {
	clickable := b.onClick != nil
	hovered := b.Element.Hovered()
	focused := b.Element.Focused()
	pressed := b.Element.Pressed()

	var bg ui.Color
	if clickable {
		bg = ApplyStateLayer(ui.Transparent, b.scheme.OnSurface, hovered, focused, pressed)
	} else {
		bg = ui.Transparent
	}

	b.Element.Background(bg)
}
