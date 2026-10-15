package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

func TestCard_VariantsRenderAndClick(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	variants := []struct {
		name    string
		variant CardVariant
	}{
		{"ElevatedCard", CardElevated},
		{"FilledCard", CardFilled},
		{"OutlinedCard", CardOutlined},
	}

	for _, tc := range variants {
		t.Run(tc.name, func(t *testing.T) {
			clicked := false
			tt := ui.NewTester(func(c *ui.Context) {
				NewCard(c, tc.variant, th).
					Children(func() {
						ui.Text(c, tc.name)
					}).
					OnClick(func() {
						clicked = true
					})
			}, 300, 200)

			if !tt.HasText(tc.name) {
				t.Fatalf("card content text %s not found", tc.name)
			}

			if err := tt.Click(tc.name); err != nil {
				t.Fatalf("failed clicking card %s: %v", tc.name, err)
			}

			if !clicked {
				t.Errorf("expected card %s OnClick to fire", tc.name)
			}
		})
	}
}
