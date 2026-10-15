package mayu

import (
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/components"
	"github.com/harukizmoe/mayu/theme"
)

// Density re-exports theme.Density for convenient access.
type Density = theme.Density

const (
	DensityStandard    Density = theme.DensityStandard
	DensityCompact     Density = theme.DensityCompact
	DensityComfortable Density = theme.DensityComfortable
)

type ButtonVariant = components.ButtonVariant

const (
	VariantFilled      ButtonVariant = components.VariantFilled
	VariantElevated    ButtonVariant = components.VariantElevated
	VariantFilledTonal ButtonVariant = components.VariantFilledTonal
	VariantOutlined    ButtonVariant = components.VariantOutlined
	VariantText        ButtonVariant = components.VariantText
)

// Re-exported option types.
type AlertDialogOptions = components.AlertDialogOptions
type SideSheetOptions = components.SideSheetOptions
type CapsuleSliderOptions = components.CapsuleSliderOptions
type LabeledSliderOptions = components.LabeledSliderOptions
type SegmentOption = components.SegmentOption
type ListItemOptions = components.ListItemOptions
type NavDestination = components.NavDestination
type NavigationRailOptions = components.NavigationRailOptions
type SnackbarOptions = components.SnackbarOptions
type MenuItemOptions = components.MenuItemOptions
type MenuPanel = components.MenuPanel
type ToastOptions = components.ToastOptions
type ToastHostOptions = components.ToastHostOptions
type ToastKind = components.ToastKind
type ToastPlacement = components.ToastPlacement

// Toast kinds, which choose a toast's icon and its color.
const (
	ToastInfo    ToastKind = components.ToastInfo
	ToastSuccess ToastKind = components.ToastSuccess
	ToastError   ToastKind = components.ToastError
)

// Where a ToastHost holds its toasts across the window.
const (
	ToastBottomCenter ToastPlacement = components.ToastBottomCenter
	ToastBottomStart  ToastPlacement = components.ToastBottomStart
	ToastBottomEnd    ToastPlacement = components.ToastBottomEnd
)

// ThemeOptions configures the Material 3 design system for the window.
type ThemeOptions struct {
	SeedColor ui.Color
	Dark      bool
	Density   theme.Density
}

type themeKey struct{}

// Use applies Material 3 ThemeOptions to the window Context.
// Call it first in every frame of the root view.
func Use(c *ui.Context, opts ThemeOptions) *theme.Theme {
	seed := opts.SeedColor
	if seed == (ui.Color{}) {
		seed = ui.Hex("#6750A4") // Canonical M3 seed
	}

	scheme := theme.SchemeFromSeed(seed, opts.Dark)
	t := theme.NewTheme(scheme, opts.Density)

	// Propagate M3 colors and metrics to upstream MyGo ui.Theme
	c.SetTheme(&ui.Theme{
		Dark:           opts.Dark,
		Background:     scheme.Background,
		Surface:        scheme.SurfaceContainer,
		SurfaceHover:   scheme.SurfaceContainerHigh,
		SurfacePressed: scheme.SurfaceContainerHighest,
		Border:         scheme.OutlineVariant,
		Text:           scheme.OnSurface,
		TextMuted:      scheme.OnSurfaceVariant,
		Accent:         scheme.Primary,
		AccentHover:    scheme.PrimaryContainer,
		AccentPressed:  scheme.PrimaryContainer,
		AccentText:     scheme.OnPrimary,
		Danger:         scheme.Error,
		Warning:        ui.Hex("#FF9800"),
		Success:        ui.Hex("#4CAF50"),
		Selection:      scheme.PrimaryContainer,
		Focus:          scheme.Primary,
		Inverse:        scheme.InverseSurface,
		InverseText:    scheme.InverseOnSurface,
		Radius:         opts.Density.ControlHeight() / 2,
		Spacing:        opts.Density.SpacingUnit(),
		FontSize:       opts.Density.BaseFontSize(),
	})

	// Inject into ui.Local at root for child Mayu components
	*ui.Local(c.Root(), themeKey{}, func() *theme.Theme { return t }) = t

	return t
}

// CurrentTheme retrieves the active Mayu Theme from the Context, or creates a default one.
func CurrentTheme(c *ui.Context) *theme.Theme {
	loc := ui.Local(c.Root(), themeKey{}, func() *theme.Theme {
		scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), c.Theme().Dark)
		return theme.NewTheme(scheme, theme.DensityStandard)
	})
	if loc != nil && *loc != nil {
		return *loc
	}
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), c.Theme().Dark)
	return theme.NewTheme(scheme, theme.DensityStandard)
}

// Scheme returns the active M3 ColorScheme from the Context.
func Scheme(c *ui.Context) theme.ColorScheme {
	return CurrentTheme(c).Scheme
}

// ==========================================
// 1. Buttons
// ==========================================

// Button constructs an M3 Button with the specified variant.
func Button(c *ui.Context, label string, variant ButtonVariant) *components.ButtonBuilder {
	return components.NewButton(c, label, variant, CurrentTheme(c))
}

// FilledButton constructs an M3 high-emphasis Filled Button.
func FilledButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantFilled, CurrentTheme(c))
}

// ElevatedButton constructs an M3 Elevated Button with tonal elevation.
func ElevatedButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantElevated, CurrentTheme(c))
}

// FilledTonalButton constructs an M3 medium-emphasis Filled Tonal Button.
func FilledTonalButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantFilledTonal, CurrentTheme(c))
}

// OutlinedButton constructs an M3 medium-low emphasis Outlined Button.
func OutlinedButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantOutlined, CurrentTheme(c))
}

// TextButton constructs an M3 low-emphasis Text Button.
func TextButton(c *ui.Context, label string) *components.ButtonBuilder {
	return components.NewButton(c, label, components.VariantText, CurrentTheme(c))
}

// FAB constructs a standard 56x56 Floating Action Button.
func FAB(c *ui.Context, icon *ui.SVG) *components.FABBuilder {
	return components.NewFAB(c, icon, CurrentTheme(c))
}

// ExtendedFAB constructs an Extended FAB with icon and label.
func ExtendedFAB(c *ui.Context, icon *ui.SVG, label string) *components.FABBuilder {
	return components.NewExtendedFAB(c, icon, label, CurrentTheme(c))
}

// IconButton constructs a standard circular Icon Button.
func IconButton(c *ui.Context, icon *ui.SVG) *components.IconButtonBuilder {
	return components.NewIconButton(c, icon, components.IconButtonStandard, CurrentTheme(c))
}

// FilledIconButton constructs a filled circular Icon Button.
func FilledIconButton(c *ui.Context, icon *ui.SVG) *components.IconButtonBuilder {
	return components.NewIconButton(c, icon, components.IconButtonFilled, CurrentTheme(c))
}

// FilledTonalIconButton constructs a filled tonal circular Icon Button.
func FilledTonalIconButton(c *ui.Context, icon *ui.SVG) *components.IconButtonBuilder {
	return components.NewIconButton(c, icon, components.IconButtonFilledTonal, CurrentTheme(c))
}

// OutlinedIconButton constructs an outlined circular Icon Button.
func OutlinedIconButton(c *ui.Context, icon *ui.SVG) *components.IconButtonBuilder {
	return components.NewIconButton(c, icon, components.IconButtonOutlined, CurrentTheme(c))
}

// ==========================================
// 2. Cards
// ==========================================

// ElevatedCard constructs an M3 Elevated Card with tonal elevation shadow.
func ElevatedCard(c *ui.Context) *components.CardBuilder {
	return components.NewCard(c, components.CardElevated, CurrentTheme(c))
}

// FilledCard constructs an M3 Filled Card using SurfaceContainerHighest.
func FilledCard(c *ui.Context) *components.CardBuilder {
	return components.NewCard(c, components.CardFilled, CurrentTheme(c))
}

// OutlinedCard constructs an M3 Outlined Card with an outline border.
func OutlinedCard(c *ui.Context) *components.CardBuilder {
	return components.NewCard(c, components.CardOutlined, CurrentTheme(c))
}

// ==========================================
// 3. Selection Controls
// ==========================================

// Checkbox constructs an M3 Checkbox toggling *checked.
func Checkbox(c *ui.Context, checked *bool, label string) *components.CheckboxBuilder {
	return components.NewCheckbox(c, checked, label, CurrentTheme(c))
}

// Switch constructs an M3 Switch toggling *on.
func Switch(c *ui.Context, on *bool, label string) *components.SwitchBuilder {
	return components.NewSwitch(c, on, label, CurrentTheme(c))
}

// Radio constructs an M3 Radio selecting value into *selected.
func Radio[T comparable](c *ui.Context, selected *T, value T, label string) *components.RadioBuilder[T] {
	return components.NewRadio(c, selected, value, label, CurrentTheme(c))
}

// ==========================================
// 4. Text Fields
// ==========================================

// FilledTextField constructs an M3 Filled Text Field.
func FilledTextField(c *ui.Context, value *string, label string) *components.TextFieldBuilder {
	return components.NewTextField(c, value, label, components.TextFieldFilled, CurrentTheme(c))
}

// OutlinedTextField constructs an M3 Outlined Text Field.
func OutlinedTextField(c *ui.Context, value *string, label string) *components.TextFieldBuilder {
	return components.NewTextField(c, value, label, components.TextFieldOutlined, CurrentTheme(c))
}

// ==========================================
// 5. Chips
// ==========================================

// Chip constructs an M3 Chip.
func Chip(c *ui.Context, label string) *components.ChipBuilder {
	return components.NewChip(c, label, CurrentTheme(c))
}

// FilterChip constructs an M3 Filter Chip with selection state.
func FilterChip(c *ui.Context, label string, selected bool) *components.ChipBuilder {
	return components.NewChip(c, label, CurrentTheme(c)).Selected(selected)
}

// ==========================================
// 6. Sliders & Dividers
// ==========================================

// Slider constructs an M3 continuous Slider between lo and hi.
func Slider(c *ui.Context, value *float64, lo, hi float64) *components.SliderBuilder {
	return components.NewSlider(c, value, lo, hi, CurrentTheme(c))
}

// VerticalCapsuleSlider creates an M3/Control Center vertical wide capsule slider (e.g. brightness or volume).
func VerticalCapsuleSlider(c *ui.Context, value *float64, lo, hi float64, icon *ui.SVG) *components.CapsuleSliderBuilder {
	return components.VerticalCapsuleSlider(c, value, lo, hi, icon, CurrentTheme(c))
}

// HorizontalCapsuleSlider creates a horizontal wide capsule slider with icon and label.
func HorizontalCapsuleSlider(c *ui.Context, value *float64, lo, hi float64, icon *ui.SVG, label string) *components.CapsuleSliderBuilder {
	return components.HorizontalCapsuleSlider(c, value, lo, hi, icon, label, CurrentTheme(c))
}

// CapsuleSlider creates a wide capsule slider with customized options.
func CapsuleSlider(c *ui.Context, value *float64, lo, hi float64, opts CapsuleSliderOptions) *components.CapsuleSliderBuilder {
	return components.NewCapsuleSlider(c, value, lo, hi, opts, CurrentTheme(c))
}

// LabeledSlider creates a card-style slider: a header of an icon, a label and a
// value badge above a full-width track with a round thumb.
func LabeledSlider(c *ui.Context, value *float64, lo, hi float64, opts LabeledSliderOptions) ui.Element {
	return components.LabeledSlider(c, value, lo, hi, opts, CurrentTheme(c))
}
func HorizontalDivider(c *ui.Context) ui.Element {
	return components.HorizontalDivider(c, CurrentTheme(c))
}

// VerticalDivider creates a 1dp vertical separator line.
func VerticalDivider(c *ui.Context) ui.Element {
	return components.VerticalDivider(c, CurrentTheme(c))
}

// ==========================================
// 7. Indicators, Badges & Dialogs
// ==========================================

// LinearProgress creates an M3 Linear Progress bar (progress in 0..1).
func LinearProgress(c *ui.Context, progress float32) ui.Element {
	return components.LinearProgress(c, progress, CurrentTheme(c))
}

// Badge creates an M3 notification badge.
func Badge(c *ui.Context, count int) ui.Element {
	return components.Badge(c, count, CurrentTheme(c))
}

// AlertDialog renders an M3 Alert Dialog when *open is true.
func AlertDialog(c *ui.Context, open *bool, opts AlertDialogOptions) ui.Element {
	return components.AlertDialog(c, open, opts, CurrentTheme(c))
}

// SideSheet renders an M3 modal side sheet anchored to the right edge while
// *open is true. The side-anchored sibling of AlertDialog: same overlay, but a
// full-height panel instead of a centred box.
func SideSheet(c *ui.Context, open *bool, opts SideSheetOptions, content func()) ui.Element {
	return components.SideSheet(c, open, opts, CurrentTheme(c), content)
}

// ==========================================
// 8. Architecture & Navigation
// ==========================================

// SegmentedButton creates a connected M3 Segmented Button group.
func SegmentedButton(c *ui.Context, selected *int, options []SegmentOption) ui.Element {
	return components.SegmentedButton(c, selected, options, CurrentTheme(c))
}

// ListItem creates an M3 List Item with slots for leading, supporting text, and trailing.
func ListItem(c *ui.Context, headline string, opts ListItemOptions) *components.ListItemBuilder {
	return components.NewListItem(c, headline, opts, CurrentTheme(c))
}

// NavigationRail creates an M3 vertical desktop navigation rail (80dp width).
func NavigationRail(c *ui.Context, opts NavigationRailOptions) ui.Element {
	return components.NavigationRail(c, opts, CurrentTheme(c))
}

// ==========================================
// 9. Overlays, Menus & Tooltips
// ==========================================

// Tooltip attaches an M3 floating tooltip to an anchor element.
func Tooltip(c *ui.Context, anchor ui.Element, text string) ui.Element {
	return components.Tooltip(c, anchor, text, CurrentTheme(c))
}

// Snackbar renders an M3 Snackbar floating feedback pill.
func Snackbar(c *ui.Context, message string, opts SnackbarOptions) ui.Element {
	return components.Snackbar(c, message, opts, CurrentTheme(c))
}

// Menu renders an M3 Dropdown Menu popup anchored to an element when *open is true.
func Menu(c *ui.Context, anchor ui.Element, open *bool, items func(m *MenuPanel)) ui.Element {
	return components.Menu(c, anchor, open, items, CurrentTheme(c))
}

// ==========================================
// 9b. Toasts
// ==========================================

// ToastHost installs the Material 3 look of the window's toasts. Call it once
// in the view, anywhere; ShowToast then shows toasts in it.
func ToastHost(c *ui.Context) ui.Element {
	return components.ToastHost(c, components.ToastHostOptions{}, CurrentTheme(c))
}

// ToastHostWith installs the toast host with a chosen placement and width.
func ToastHostWith(c *ui.Context, opts ToastHostOptions) ui.Element {
	return components.ToastHost(c, opts, CurrentTheme(c))
}

// ShowToast shows a toast and returns its ID, to close it with CloseToast.
func ShowToast(c *ui.Context, opts ToastOptions) string {
	return components.ShowToast(c, opts)
}

// CloseToast closes the toast with id, or every toast when id is empty.
func CloseToast(c *ui.Context, id string) {
	components.CloseToast(c, id)
}

// ==========================================
// 10. Visual & Content Widgets
// ==========================================

// Avatar creates a circular M3 Avatar showing an icon.
func Avatar(c *ui.Context, icon *ui.SVG) *components.AvatarBuilder {
	return components.Avatar(c, icon, CurrentTheme(c))
}

// InitialsAvatar creates an M3 Avatar displaying 1-2 initials extracted from name.
func InitialsAvatar(c *ui.Context, name string) *components.AvatarBuilder {
	return components.InitialsAvatar(c, name, CurrentTheme(c))
}

// Accordion creates an M3 collapsible panel with a clickable header and expandable body.
func Accordion(c *ui.Context, open *bool, title string, content func()) ui.Element {
	return components.Accordion(c, open, title, content, CurrentTheme(c))
}

// SearchBar creates an M3 Search Bar (56dp height pill).
func SearchBar(c *ui.Context, query *string, placeholder string) *components.SearchBarBuilder {
	return components.NewSearchBar(c, query, placeholder, CurrentTheme(c))
}
