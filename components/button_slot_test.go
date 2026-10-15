package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

// A button's leading slot draws content of its own, not only an icon, and
// Variant restyles the one element rather than needing a second beside it.
func TestButtonLeadingSlotAndVariant(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	swatch := ui.Hex("#B3261E")

	// A Row sizes its children by their content, so each button hugs its own
	// slots and its edges are known from its rect.
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Padding(10).Gap(10).Children(func() {
			ui.Row(c).Children(func() {
				NewButton(c, "Outlined", VariantOutlined, th).
					Leading(func() {
						ui.Box(c).Size(14, 14).Radius(7).Shrink(0).Background(swatch)
					})
			})
			ui.Row(c).Children(func() {
				NewButton(c, "Filled", VariantOutlined, th).
					Leading(func() {
						ui.Box(c).Size(14, 14).Radius(7).Shrink(0).Background(swatch)
					}).
					Variant(VariantFilled)
			})
		})
	}, 260, 120)

	img := tt.Image()
	at := func(x, y int) ui.Color {
		i := img.PixOffset(x, y)
		return ui.RGB(img.Pix[i], img.Pix[i+1], img.Pix[i+2])
	}

	for _, label := range []string{"Outlined", "Filled"} {
		r, ok := tt.Find(label)
		if !ok {
			t.Fatalf("button %q not found", label)
		}

		// Find returns the labeled element, the button itself. Its leading
		// padding is 16 DIPs and the swatch is 14 wide, so the disc's centre
		// sits 23 in from its left edge: at the start, before the label.
		x, y := int(r.X)+23, int(r.Y+r.H/2)
		if got := at(x, y); got != swatch {
			t.Errorf("%s: the leading slot shows %v at (%d,%d), want the swatch %v",
				label, got, x, y, swatch)
		}
		if x-x >= int(r.W) {
			t.Errorf("%s: the swatch is not inside the button", label)
		}

		// The fill follows the variant, inside the 24 DIPs of trailing
		// padding, as a colour the window behind it never shows.
		fill := at(int(r.X+r.W)-12, int(r.Y+r.H/2))
		switch label {
		case "Filled":
			if fill != scheme.Primary {
				t.Errorf("Filled: the container is %v, want the primary %v", fill, scheme.Primary)
			}
		case "Outlined":
			if fill == scheme.Primary {
				t.Errorf("Outlined: the container is the primary, want none of its own fill")
			}
		}
	}
}
