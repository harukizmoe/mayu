# Mayu (Material You for MyGo)

**Mayu** is an open-source Material Design 3 (Material You / M3) component library and design system written in pure Go for [MyGo](https://github.com/egoist/mygo)'s native GPU UI toolkit (`mygo/ui`).

## Features

- **Dynamic Color (HCT / CAM16)**: Seed-color-based tonal palette generation implementing Google's CAM16 color appearance model and perceptual $L^*$ Tone matching in pure Go.
- **Tonal Elevation**: Z-axis hierarchy represented through surface containers and subtle tonal variations rather than heavy blurred shadows.
- **M3 State Layers**: Interactive state feedback (Hover 8%, Focus 10%, Pressed 10%) mathematically composited over base containers.
- **Desktop Density Scale**: Three density tiers (`Standard`, `Compact`, `Comfortable`) for keyboard/mouse desktop ergonomics.
- **Full Common Component Suite**: Comprehensive implementations of Buttons, Cards, Selection Controls, TextFields, Chips, Sliders, Dialogs, Side Sheets, Dividers, and Progress Indicators.
- **Fluent Builder API**: Idiomatic Go chaining for components with seamless MyGo `ui.Element` interoperability.

## Component Catalog

| Category | Components | Description |
|---|---|---|
| **Buttons** | `FilledButton`, `ElevatedButton`, `FilledTonalButton`, `OutlinedButton`, `TextButton` | 5 M3 emphasis variants with continuous pill radius and leading/trailing icons |
| **Floating Actions** | `FAB`, `ExtendedFAB` | 56x56 floating action buttons with elevation level 3 |
| **Icon Buttons** | `IconButton`, `FilledIconButton`, `FilledTonalIconButton`, `OutlinedIconButton` | 40x40 circular icon action buttons |
| **Cards** | `ElevatedCard`, `FilledCard`, `OutlinedCard` | 12dp rounded cards with tonal surface containers and clickable states |
| **Selection** | `Checkbox`, `Switch`, `Radio` | M3 toggle controls with responsive state layers and accompanying labels |
| **Text Inputs** | `FilledTextField`, `OutlinedTextField` | Single-line text inputs with floating labels, placeholders, and error/helper text |
| **Chips** | `Chip`, `FilterChip` | Compact 32dp action tags and selected filter chips |
| **Sliders** | `Slider` | Continuous track sliders with active/inactive colored ranges and round thumb |
| **Dividers** | `HorizontalDivider`, `VerticalDivider` | 1dp hairline separators in `OutlineVariant` |
| **Indicators & Badges** | `LinearProgress`, `Badge` | Determinate progress bars and notification count pills/dots |
| **Dialogs** | `AlertDialog` | Modal dialogs with Scrim backdrop, `SurfaceContainerHigh` panel, title, body, and action buttons |
| **Side Sheets** | `SideSheet` | Modal side sheet anchored to the right edge: full-height `SurfaceContainerLow` panel with rounded left corners, header, scrolling content, and pinned footer actions. Dims and blocks the content behind it, takes focus, closes on Escape or scrim click |
| **Capsule Sliders** | `VerticalCapsuleSlider`, `HorizontalCapsuleSlider`, `CapsuleSlider` | Control-Center style wide capsule sliders (72x160 / 220x60) with auto-contrasting interior icon and label |
| **Labeled Slider** | `LabeledSlider` | Card-style slider: header of icon, label and a live value badge above a full-width track with a ringed round thumb |
| **Navigation** | `NavigationRail`, `SegmentedButton` | 80dp desktop navigation rail with active pill indicator + connected M3 segmented button group |
| **Lists** | `ListItem` | 3-slot list rows (leading icon/avatar, headline + supporting text, trailing switch/badge) |
| **Overlays** | `Tooltip`, `Snackbar`, `Menu` | Inverse-surface tooltips, floating feedback pills with action, anchored dropdown menus with shortcuts |
| **Notifications** | `ToastHost`, `ShowToast`, `CloseToast` | Material 3 toasts hugging their content, auto-dismissing, deduped by title. Two outlines (`ToastCapsule` pill / `ToastCorner` rounded rectangle) and two palettes (`ToastSurface` in the app's own colors / `ToastInverse` the classic dark snackbar), info/success/error kind icons, an action and a close button |
| **Media & Structure** | `Avatar`, `InitialsAvatar`, `Accordion`, `SearchBar` | Circular avatars (icon or hash-initials), collapsible panels, and 56dp search pills |

## Installation

```sh
go get github.com/harukizmoe/mayu
```

## Quick Start

```go
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
	"github.com/harukizmoe/mayu/symbols"
)

type appState struct {
	checked bool
	volume  float64
	text    string
}

func (s *appState) view(c *ui.Context) {
	// 1. Initialize M3 Theme with seed color
	mayu.Use(c, mayu.ThemeOptions{
		SeedColor: ui.Hex("#6750A4"),
		Dark:      false,
		Density:   mayu.DensityStandard,
	})

	// 2. Build fluent M3 components
	ui.Column(c).Fill().Padding(24).Gap(16).Children(func() {
		// Buttons
		ui.Row(c).Gap(12).Children(func() {
			mayu.FilledButton(c, "Submit").
				Icon(symbols.Send).
				OnClick(func() { println("Submitted!") })

			mayu.OutlinedButton(c, "Cancel")
		})

		// Selection Controls
		mayu.Checkbox(c, &s.checked, "Enable Notifications")

		// Sliders
		mayu.Slider(c, &s.volume, 0, 100)

		// TextFields
		mayu.OutlinedTextField(c, &s.text, "Username").
			Placeholder("Enter username")
	})
}

func main() {
	s := &appState{volume: 50.0}
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Mayu App",
			Width:   640,
			Height:  480,
			Content: ui.View(s.view),
		})
	})
	mygo.App.Run()
}
```

## Repository Structure

```text
mayui/
├── go.mod                     // github.com/harukizmoe/mayu
├── README.md
├── mayu.go                    // Public facade: Use(), Context injection, constructors
├── theme/                     // Design tokens & styling contracts
│   ├── scheme.go              // ColorScheme (Primary, SurfaceContainer...)
│   ├── typography.go          // Display, Headline, Title, Body, Label tokens
│   ├── shape.go               // Corner radius scale (None to Full)
│   └── density.go             // Compact / Standard / Comfortable metrics
├── color/                     // Algorithmic foundation (HCT & Tonal Palettes)
│   ├── cam16.go               // CAM16 color appearance model calculations
│   ├── hct.go                 // HCT (Hue, Chroma, Tone) representation & conversions
│   ├── palette.go             // TonalPalette (0~100 tone steps)
│   └── dynamic.go             // Seed Color -> 6 Tonal Palettes generator
├── components/                // M3 Widget implementations
│   ├── button.go              // Filled, Elevated, Tonal, Outlined, Text
│   ├── fab.go                 // FAB, Extended FAB
│   ├── icon_button.go         // Circular Icon Buttons
│   ├── card.go                // Elevated, Filled, Outlined Cards
│   ├── selection.go           // Checkbox, Switch, Radio
│   ├── text_field.go          // Filled, Outlined TextField
│   ├── chip.go                // Action Chip, Filter Chip
│   ├── slider.go              // Continuous Slider
│   ├── divider.go             // Horizontal & Vertical Dividers
│   ├── indicator.go           // LinearProgress & Badge
│   ├── dialog.go              // AlertDialog
│   ├── side_sheet.go          // SideSheet
│   └── state_layer.go         // Hover (8%), Focus (10%), Pressed (10%) overlays
├── symbols/                   // Standard Material Symbols (SVG)
└── examples/
    └── gallery/               // Comprehensive interactive showcase
        └── main.go
```

## Running the Gallery

```sh
cd examples/gallery
go run main.go
```

## Testing

Run unit tests and race detection across all packages:

```sh
go test -race ./...
```
