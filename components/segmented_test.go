package components

import (
	"math"
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

// A segmented button's segments share the track's width equally, so the labels
// are evenly spaced and the group hugs its right edge. Content-sized segments
// space them by their own text instead, leaving a gap after the last one.
//
// Every option carries an icon and the chosen one shows a check of the same
// size, so each label sits the same distance into its own share; the spacing
// between labels is then the share itself.
func TestSegmentedButtonSegmentsShareTheWidth(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	const (
		window   = 360
		padding  = 10
		segments = 3
	)
	share := float32(window-2*padding) / segments

	chosen := 1
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Padding(padding).Children(func() {
			SegmentedButton(c, &chosen, []SegmentOption{
				{Label: "AAA", Icon: symbols.Favorite},
				{Label: "BBB", Icon: symbols.Favorite},
				{Label: "CCC", Icon: symbols.Favorite},
			}, th)
		})
	}, window, 80)

	centres := make([]float32, 0, segments)
	for _, label := range []string{"AAA", "BBB", "CCC"} {
		r, ok := tt.Find(label)
		if !ok {
			t.Fatalf("segment %q not found", label)
		}
		centres = append(centres, r.X+r.W/2)
	}

	for i := 1; i < segments; i++ {
		got := centres[i] - centres[i-1]
		if math.Abs(float64(got-share)) > 2.5 {
			t.Errorf("segments %d and %d are %v apart, want %v: they do not share the width",
				i-1, i, got, share)
		}
	}

	// The first label anchors the group to the track's left edge: it sits at
	// the centre of the first share plus half of the icon and gap before it.
	// Together with the even spacing this proves the segments fill the track.
	const iconBlock = 18 + 8 // the segment icon and the gap before its label
	wantFirst := padding + share*0.5 + iconBlock/2
	if math.Abs(float64(centres[0]-wantFirst)) > 2.5 {
		t.Errorf("the first segment is centred at %v, want %v: the group does not fill the track",
			centres[0], wantFirst)
	}
}
