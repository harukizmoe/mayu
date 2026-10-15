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
	// MaxWidth caps a toast's width, 420 DIPs when zero.
	MaxWidth float32
}

// ToastOptions configures one toast.
type ToastOptions struct {
	// ID names the toast: adding one with the ID of a toast showing changes
	// that toast, which shows anew. Made up when empty.
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

		for _, toast := range toasts {
			styleToast(c, ui.ToastBase(c, toast), toast, maxW, t)
		}
	})
}

// styleToast gives one toast the Material 3 look: an inverse-surface pill
// holding a kind icon, the text, the action and the close button.
func styleToast(c *ui.Context, parts ui.ToastParts, toast ui.Toast, maxW float32, t *theme.Theme) {
	scheme := t.Scheme

	root := parts.Root.Row().MaxWidth(maxW).MinHeight(48).Radius(theme.CornerExtraSmall).
		Padding(8, 16).Gap(12).AlignItems(ui.Center).
		Background(scheme.InverseSurface).
		Shadow(0, 3, 9, 3, scheme.Shadow.Alpha(0.20))

	opts, _ := toast.Data.(ToastOptions)
	icon, iconColor := toastIcon(opts, toast.Type, t)

	root.Children(func() {
		if icon != nil {
			ui.Icon(c, icon).Size(20, 20).TextColor(iconColor)
		}

		ui.Column(c).Grow(1).Justify(ui.Center).Children(func() {
			if toast.Title != "" {
				ui.Text(c, toast.Title).FontSize(theme.BodyMedium.Size).
					TextColor(scheme.InverseOnSurface)
			}
			if toast.Description != "" {
				ui.Text(c, toast.Description).FontSize(theme.BodySmall.Size).
					TextColor(scheme.InverseOnSurface.Alpha(0.78))
			}
		})

		if toast.Action != "" {
			action := parts.ActionButton().Height(36).Padding(0, 8).Center()
			action.TextColor(scheme.InversePrimary)
			action.Children(func() {
				ui.Text(c, toast.Action).FontSize(theme.LabelLarge.Size).Bold().SingleLine()
			})
		}

		if opts.Close {
			closeBtn := parts.CloseButton().Size(32, 32).Radius(16).Center().Label("Close")
			closeBtn.Children(func() {
				ui.Icon(c, symbols.Close).Size(18, 18).TextColor(scheme.InverseOnSurface.Alpha(0.78))
			})
		}
	})
}

// toastIcon returns the icon and its color for a toast: the caller's when it
// gave one, else the kind's, in a color that reads on the inverse surface.
func toastIcon(opts ToastOptions, kind string, t *theme.Theme) (*ui.SVG, ui.Color) {
	switch {
	case opts.HideIcon:
		return nil, ui.Color{}
	case opts.Icon != nil:
		return opts.Icon, toastIconColor(opts.Kind, t)
	}
	switch kind {
	case ToastSuccess.kindName():
		return symbols.Check, theme.InverseSuccess
	case ToastError.kindName():
		return symbols.Close, theme.InverseError
	default:
		return symbols.Info, t.Scheme.InversePrimary
	}
}

// toastIconColor colors a caller's own icon by the kind of its toast.
func toastIconColor(kind ToastKind, t *theme.Theme) ui.Color {
	switch kind {
	case ToastSuccess:
		return theme.InverseSuccess
	case ToastError:
		return theme.InverseError
	default:
		return t.Scheme.InversePrimary
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
	return c.AddToast(ui.Toast{
		ID:          opts.ID,
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
