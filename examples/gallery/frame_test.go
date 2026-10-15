package main

import (
	"image"
	"testing"

	"github.com/egoist/mygo/ui"
)

// pageFrame renders one page of the gallery in one mode and returns what it
// painted.
func pageFrame(t *testing.T, page int, dark bool) *image.RGBA {
	t.Helper()
	app := newGalleryApp()
	app.page = page
	app.dark = dark
	return ui.NewTester(app.view, 1000, 800).Image()
}

// distinct reports how many colors a frame holds, to tell a painted window from
// a blank one.
func distinct(img *image.RGBA) int {
	seen := map[uint32]struct{}{}
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			i := img.PixOffset(x, y)
			seen[uint32(img.Pix[i])<<16|uint32(img.Pix[i+1])<<8|uint32(img.Pix[i+2])] = struct{}{}
		}
	}
	return len(seen)
}

// differing reports the share of the frame's pixels the two frames do not
// share, from 0 to 1.
func differing(a, b *image.RGBA) float64 {
	var n, diff int
	bounds := a.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			i := a.PixOffset(x, y)
			j := b.PixOffset(x, y)
			n++
			if a.Pix[i] != b.Pix[j] || a.Pix[i+1] != b.Pix[j+1] || a.Pix[i+2] != b.Pix[j+2] {
				diff++
			}
		}
	}
	return float64(diff) / float64(n)
}

// The rail's destinations must be visible on the screen, not only in the text
// the tree holds: each page has to paint a different window, and the theme has
// to repaint all of it.
func TestGalleryPagesPaintTheirOwnPage(t *testing.T) {
	for _, dark := range []bool{false, true} {
		mode := "light"
		if dark {
			mode = "dark"
		}

		frames := make([]*image.RGBA, len(navPages))
		for page := range navPages {
			frames[page] = pageFrame(t, page, dark)

			// A blank window would hold a handful of colors: a painted one
			// holds the rail, the cards and the type.
			if got := distinct(frames[page]); got < 40 {
				t.Errorf("%s, page %q: the frame holds %d colors: it is not painted",
					mode, navPages[page].title, got)
			}
		}

		// Every page shows something the others do not.
		for i := range frames {
			for j := i + 1; j < len(frames); j++ {
				if share := differing(frames[i], frames[j]); share < 0.10 {
					t.Errorf("%s: the %q and %q pages share %.0f%% of their pixels: the rail is not switching them",
						mode, navPages[i].title, navPages[j].title, (1-share)*100)
				}
			}
		}
	}

	// And the theme repaints the window rather than a corner of it.
	for page := range navPages {
		light, dark := pageFrame(t, page, false), pageFrame(t, page, true)
		if share := differing(light, dark); share < 0.20 {
			t.Errorf("page %q paints only %.0f%% of its pixels differently in the two modes",
				navPages[page].title, share*100)
		}
	}
}
