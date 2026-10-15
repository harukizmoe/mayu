package components

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// ToastKind names what a toast is about, which chooses its icon and its color.
type ToastKind int

const (
	// ToastInfo is a plain notice.
	ToastInfo ToastKind = iota
	// ToastSuccess marks what went well.
	ToastSuccess
	// ToastError marks what failed.
	ToastError
)

// kindName is the ui.Toast.Type a kind is stored under, which the look reads
// back.
func (k ToastKind) kindName() string {
	switch k {
	case ToastSuccess:
		return "success"
	case ToastError:
		return "error"
	default:
		return "info"
	}
}

// ToastShape is the outline a toast is drawn with.
type ToastShape int

const (
	// ToastCorner is the Material 3 rounded rectangle of a snackbar: 4 DIP
	// corners, 48 DIPs high for one line and 68 for two.
	ToastCorner ToastShape = iota
	// ToastCapsule is the Material 3 Expressive pill: fully rounded, 52 DIPs
	// high for one line and 72 for two, with its kind icon in a tinted disc.
	ToastCapsule
)

// ToastColor is which color roles a host fills its toasts with.
type ToastColor int

const (
	// ToastSurface fills a toast with SurfaceContainerHigh, its text with
	// OnSurface and its action with Primary: a Material You toast that sits in
	// the app's own colors, as its menus and dialogs do.
	ToastSurface ToastColor = iota
	// ToastInverse fills a toast with InverseSurface and its text with
	// InverseOnSurface: the classic Material 3 snackbar, which stands out from
	// the app by turning the theme over.
	ToastInverse
)

// ToastPlacement is where a ToastHost holds its toasts across the window.
type ToastPlacement int

const (
	// ToastBottomCenter centers toasts along the bottom edge.
	ToastBottomCenter ToastPlacement = iota
	// ToastBottomStart puts them at the start, the left on a left-to-right
	// layout.
	ToastBottomStart
	// ToastBottomEnd puts them at the end, the right on a left-to-right
	// layout.
	ToastBottomEnd
)

// ToastHostOptions configures a ToastHost.
type ToastHostOptions struct {
	// Placement is where the toasts sit across the window.
	Placement ToastPlacement
	// Shape is the outline every toast of the host is drawn with.
	Shape ToastShape
	// Color is the color roles every toast of the host is filled with.
	Color ToastColor
	// MaxWidth caps a toast's width, 420 DIPs when zero. A toast hugs its
	// content up to that width.
	MaxWidth float32
}

// toastPalette is the roles one ToastColor resolves to.
type toastPalette struct {
	container ui.Color // the toast's fill
	on        ui.Color // its title
	muted     ui.Color // its description
	action    ui.Color // its action's label
	close     ui.Color // its close button
	border    ui.Color // a hairline edge, none when zero
	discAlpha float32  // how strong the icon's disc is over the container
	// icon colors the kind icon and its disc for this palette.
	icon func(kind ToastKind, t *theme.Theme) ui.Color
}

// paletteFor resolves the roles of a toast color in a theme.
func paletteFor(color ToastColor, t *theme.Theme) toastPalette {
	s := t.Scheme
	if color == ToastInverse {
		return toastPalette{
			container: s.InverseSurface,
			on:        s.InverseOnSurface,
			muted:     s.InverseOnSurface.Alpha(0.75),
			action:    s.InversePrimary,
			close:     s.InverseOnSurface.Alpha(0.75),
			discAlpha: 0.20,
			icon: func(kind ToastKind, t *theme.Theme) ui.Color {
				switch kind {
				case ToastSuccess:
					return theme.InverseSuccess
				case ToastError:
					return theme.InverseError
				default:
					return t.Scheme.InversePrimary
				}
			},
		}
	}
	return toastPalette{
		container: s.SurfaceContainerHigh,
		on:        s.OnSurface,
		muted:     s.OnSurfaceVariant,
		action:    s.Primary,
		close:     s.OnSurfaceVariant,
		border:    s.OutlineVariant,
		discAlpha: 0.14,
		icon: func(kind ToastKind, t *theme.Theme) ui.Color {
			switch kind {
			case ToastSuccess:
				return theme.Success
			case ToastError:
				return t.Scheme.Error
			default:
				return t.Scheme.Primary
			}
		},
	}
}

// ToastOptions configures one toast.
type ToastOptions struct {
	// ID names the toast: adding one with the ID of a toast showing changes
	// that toast, which shows anew. The title when empty, so a repeated toast
	// refreshes instead of piling up.
	ID string
	// Title is what the toast says, and Description a second line under it.
	Title       string
	Description string
	// Kind chooses the icon and its color.
	Kind ToastKind
	// Icon replaces the kind's icon; HideIcon shows none.
	Icon     *ui.SVG
	HideIcon bool
	// Action labels a button that runs OnAction and closes the toast.
	Action   string
	OnAction func()
	// OnClose runs as the toast closes, however it closed.
	OnClose func()
	// Timeout is how long the toast shows: 8 seconds with an action and 4
	// without when zero, until it is closed when negative.
	Timeout time.Duration
	// Close adds a close button.
	Close bool
}

// ToastHost installs the Material 3 look of the window's toasts. Call it once
// in the view, anywhere; ShowToast then shows toasts in it.
func ToastHost(c *ui.Context, opts ToastHostOptions, t *theme.Theme) ui.Element {
	maxW := opts.MaxWidth
	if maxW <= 0 {
		maxW = 420
	}

	align := ui.Center
	switch opts.Placement {
	case ToastBottomStart:
		align = ui.Start
	case ToastBottomEnd:
		align = ui.End
	}

	return ui.ToastViewportBase(c, func(viewport ui.Element, toasts []ui.Toast) {
		viewport.Padding(16).Gap(8).AlignItems(align)

		pal := paletteFor(opts.Color, t)
		for _, toast := range toasts {
			styleToast(c, ui.ToastBase(c, toast), toast, opts.Shape, pal, maxW, t)
		}
	})
}

// toastMetrics are the lengths a toast's shape gives it.
type toastMetrics struct {
	height  float32 // the toast's height
	radius  float32 // its corner radius, half the height for a capsule
	padX    float32 // its horizontal padding
	gap     float32 // between its parts
	disc    float32 // the diameter of the icon's disc, 0 for a bare icon
	icon    float32 // the icon's size
	textMax float32 // the widest its text may be
	actionH float32 // the height of its action button
}

// metricsFor returns the lengths of shape for a toast of lines lines.
func metricsFor(shape ToastShape, lines int, maxW float32) toastMetrics {
	m := toastMetrics{}
	if shape == ToastCapsule {
		m.height, m.padX, m.gap, m.disc, m.icon, m.actionH = 52, 20, 10, 32, 18, 40
		m.textMax = maxW - 220
		if lines > 1 {
			m.height, m.textMax = 72, maxW-240
		}
		m.radius = m.height / 2
		return m
	}

	// ToastCorner: the Material 3 snackbar's rounded rectangle, with a bare
	// icon instead of the capsule's tinted disc.
	m.height, m.padX, m.gap, m.icon, m.actionH = 48, 16, 12, 20, 36
	m.textMax = maxW - 150
	if lines > 1 {
		m.height, m.textMax = 68, maxW-170
	}
	m.radius = theme.CornerExtraSmall
	return m
}

// styleToast gives one toast the look of shape and pal: a container hugging
// its content, holding a kind icon, the text, the action and the close button.
func styleToast(c *ui.Context, parts ui.ToastParts, toast ui.Toast, shape ToastShape, pal toastPalette, maxW float32, t *theme.Theme) {
	lines := 1
	if toast.Description != "" {
		lines = 2
	}
	m := metricsFor(shape, lines, maxW)

	root := parts.Root.Row().MaxWidth(maxW).Height(m.height).Radius(m.radius).
		Padding(0, m.padX).Gap(m.gap).AlignItems(ui.Center).
		Background(pal.container).
		Shadow(0, 3, 8, 1, t.Scheme.Shadow.Alpha(0.16))
	if pal.border.A > 0 {
		// A hairline keeps a surface-colored toast off a busy background.
		root.Border(1, pal.border)
	}

	opts, _ := toast.Data.(ToastOptions)
	icon := toastIcon(opts, toast, pal)

	root.Children(func() {
		if icon != nil {
			iconColor := pal.icon(opts.Kind, t)
			if m.disc > 0 {
				// A tinted disc behind the icon, as Material 3 draws it.
				ui.Box(c).Size(m.disc, m.disc).Radius(m.disc / 2).Center().Shrink(0).
					Background(iconColor.Alpha(pal.discAlpha)).Children(func() {
					ui.Icon(c, icon).Size(m.icon, m.icon).TextColor(iconColor)
				})
			} else {
				ui.Icon(c, icon).Size(m.icon, m.icon).TextColor(iconColor).Shrink(0)
			}
		}

		ui.Column(c).MaxWidth(m.textMax).Justify(ui.Center).Gap(2).Children(func() {
			if toast.Title != "" {
				ui.Text(c, toast.Title).FontSize(theme.BodyMedium.Size).
					TextColor(pal.on).SingleLine().Ellipsis("…")
			}
			if toast.Description != "" {
				ui.Text(c, toast.Description).FontSize(theme.BodySmall.Size).
					TextColor(pal.muted).SingleLine().Ellipsis("…")
			}
		})

		if toast.Action != "" {
			action := parts.ActionButton().Height(m.actionH).Radius(m.actionH/2).
				Padding(0, 14).Center().Shrink(0)
			action.Background(ApplyStateLayer(ui.Transparent, pal.action,
				action.Hovered(), action.Focused(), action.Pressed()))
			action.TextColor(pal.action)
			action.Children(func() {
				ui.Text(c, toast.Action).FontSize(theme.LabelLarge.Size).Bold().SingleLine()
			})
		}

		if opts.Close {
			closeBtn := parts.CloseButton().Size(32, 32).Radius(16).Center().
				Shrink(0).Label("Close")
			closeBtn.Background(ApplyStateLayer(ui.Transparent, pal.close,
				closeBtn.Hovered(), closeBtn.Focused(), closeBtn.Pressed()))
			closeBtn.Children(func() {
				ui.Icon(c, symbols.Close).Size(18, 18).TextColor(pal.close)
			})
		}
	})
}

// toastIcon returns the icon for a toast: the caller's when it gave one, else
// the one of its kind, or nil for none. Its color comes from the palette.
func toastIcon(opts ToastOptions, toast ui.Toast, pal toastPalette) *ui.SVG {
	switch {
	case opts.HideIcon:
		return nil
	case opts.Icon != nil:
		return opts.Icon
	}
	switch toast.Type {
	case ToastSuccess.kindName():
		return symbols.Check
	case ToastError.kindName():
		return symbols.Close
	default:
		return symbols.Info
	}
}

// ShowToast shows a toast in the window's host, which ToastHost installs for
// the Material 3 look, and returns its ID, to close it with CloseToast.
func ShowToast(c *ui.Context, opts ToastOptions) string {
	timeout := opts.Timeout
	if timeout == 0 && opts.Action != "" {
		// An action needs time to reach, as the standard look allows.
		timeout = 8 * time.Second
	}
	id := opts.ID
	if id == "" {
		// Naming it by its title makes a repeated toast refresh the one
		// showing instead of piling up.
		id = opts.Title
	}
	return c.AddToast(ui.Toast{
		ID:          id,
		Title:       opts.Title,
		Description: opts.Description,
		Type:        opts.Kind.kindName(),
		Timeout:     timeout,
		Action:      opts.Action,
		OnAction:    opts.OnAction,
		OnClose:     opts.OnClose,
		Data:        opts,
	})
}

// CloseToast closes the toast with id, or every toast when id is empty.
func CloseToast(c *ui.Context, id string) {
	c.CloseToast(id)
}
