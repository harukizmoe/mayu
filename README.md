# Mayu (Material You for MyGo)

**Mayu** is an open-source Material Design 3 (Material You / M3) component library and design system written in pure Go for [MyGo](https://github.com/egoist/mygo)'s native GPU UI toolkit (`mygo/ui`).

## Features

- **Dynamic Color (HCT / CAM16)**: Seed-color-based tonal palette generation implementing Google's CAM16 color appearance model and perceptual $L^*$ Tone matching in pure Go.
- **Tonal Elevation**: Z-axis hierarchy represented through surface containers and subtle tonal variations rather than heavy blurred shadows.
- **M3 State Layers**: Interactive state feedback (Hover 8%, Focus 10%, Pressed 10%) mathematically composited over base containers.
- **Desktop Density Scale**: Three density tiers (`Standard`, `Compact`, `Comfortable`) for keyboard/mouse desktop ergonomics.
- **Fluent Builder API**: Idiomatic Go chaining for components with seamless MyGo `ui.Element` interoperability.

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

type appState struct{}

func (s *appState) view(c *ui.Context) {
	// 1. Initialize M3 Theme with seed color
	mayu.Use(c, mayu.ThemeOptions{
		SeedColor: ui.Hex("#6750A4"),
		Dark:      false,
		Density:   mayu.DensityStandard,
	})

	// 2. Build fluent M3 components
	ui.Column(c).Fill().Center().Gap(12).Children(func() {
		mayu.FilledButton(c, "Submit").
			Icon(symbols.Send).
			OnClick(func() {
				println("Submitted!")
			})

		mayu.OutlinedButton(c, "Cancel")
	})
}

func main() {
	s := &appState{}
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Mayu App",
			Width:   600,
			Height:  400,
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
├── mayu.go                    // Public facade: Use(), Context injection, button constructors
├── theme/                     // Design tokens & contracts
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
│   └── state_layer.go         // Hover (8%), Focus (10%), Pressed (10%) overlays
├── symbols/                   // Standard Material Symbols (SVG)
└── examples/
    └── gallery/               // Interactive showcase with live Seed Color picker
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
