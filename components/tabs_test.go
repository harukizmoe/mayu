package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

func TestTabs(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	chosen := 0
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Padding(16).Children(func() {
			Tabs(c, &chosen, []string{"Flights", "Trains", "Buses"}, th)
		})
	}, 420, 120)

	for _, name := range []string{"Flights", "Trains", "Buses"} {
		if !tt.HasText(name) {
			t.Fatalf("tab %q is missing", name)
		}
	}

	if err := tt.Click("Buses"); err != nil {
		t.Fatalf("failed clicking a tab: %v", err)
	}
	if chosen != 2 {
		t.Errorf("expected the third tab to be chosen, got %d", chosen)
	}

	// Three tabs share the row (16..398), so the third spans about 270..398 and
	// the first about 16..142. The indicator line is the row's last three DIPs:
	// the chosen tab's segment is 3 DIPs of primary and the others 1 DIP of the
	// outline color, so y=62 tells the two apart and y=63 finds the hairline.
	img := tt.Image()
	at := func(x, y int) ui.Color {
		i := img.PixOffset(x, y)
		return ui.RGB(img.Pix[i], img.Pix[i+1], img.Pix[i+2])
	}

	if got := at(330, 62); got != scheme.Primary {
		t.Errorf("the chosen tab's indicator is %v, want the primary %v: it is not three DIPs thick",
			got, scheme.Primary)
	}
	if got := at(70, 62); got == scheme.Primary {
		t.Errorf("an unchosen tab's indicator is the primary color, want the outline")
	}
	if got := at(70, 63); got != scheme.OutlineVariant {
		t.Errorf("an unchosen tab's hairline is %v, want the outline %v", got, scheme.OutlineVariant)
	}
}

func TestCircularProgress(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Row(c).Gap(20).Padding(20).Children(func() {
			CircularProgress(c, 0.25, CircularProgressOptions{}, th)
			CircularProgress(c, 0.75, CircularProgressOptions{ShowValue: true}, th)
			CircularProgress(c, -1, CircularProgressOptions{}, th)
		})
	}, 260, 100)

	if !tt.HasText("75%") {
		t.Fatalf("the value is not drawn, texts: %q", tt.Texts())
	}

	img := tt.Image()
	at := func(x, y int) ui.Color {
		i := img.PixOffset(x, y)
		return ui.RGB(img.Pix[i], img.Pix[i+1], img.Pix[i+2])
	}

	// A ring of 48 DIPs at (20, 20) has its center at (44, 44) and a stroke
	// centerline 22 DIPs out, so a quarter turn covers twelve o'clock to three.
	// Sampled at 45 degrees, which is inside that quarter, and at 135, which is
	// past it: (44±22/√2, 44∓22/√2) is about (60, 28) and (60, 60).
	cx, cy := 44, 44
	if got := at(cx, cy-22); got != scheme.Primary {
		t.Errorf("at twelve o'clock the ring is %v, want the primary %v: a determinate ring starts there",
			got, scheme.Primary)
	}
	if got := at(60, 28); got != scheme.Primary {
		t.Errorf("45 degrees round the ring is %v, want the primary %v: a quarter is filled",
			got, scheme.Primary)
	}
	if got := at(60, 60); got != scheme.SecondaryContainer {
		t.Errorf("135 degrees round the ring is %v, want the track %v", got, scheme.SecondaryContainer)
	}
	if got := at(cx, cy+22); got != scheme.SecondaryContainer {
		t.Errorf("at six o'clock the ring is %v, want the track %v", got, scheme.SecondaryContainer)
	}

	// A track is drawn all the way round, so even an empty ring shows.
	empty := ui.NewTester(func(c *ui.Context) {
		ui.Box(c).Padding(20).Children(func() {
			CircularProgress(c, 0, CircularProgressOptions{}, th)
		})
	}, 100, 100)
	eimg := empty.Image()
	i := eimg.PixOffset(44, 22)
	if got := ui.RGB(eimg.Pix[i], eimg.Pix[i+1], eimg.Pix[i+2]); got != scheme.SecondaryContainer {
		t.Errorf("an empty ring is %v at twelve o'clock, want its track %v", got, scheme.SecondaryContainer)
	}
}
