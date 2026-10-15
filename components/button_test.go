package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/symbols"
	"github.com/harukizmoe/mayu/theme"
)

func TestButton_VariantsRenderAndClick(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	variants := []struct {
		name    string
		variant ButtonVariant
	}{
		{"Filled", VariantFilled},
		{"Elevated", VariantElevated},
		{"FilledTonal", VariantFilledTonal},
		{"Outlined", VariantOutlined},
		{"Text", VariantText},
	}

	for _, tc := range variants {
		t.Run(tc.name, func(t *testing.T) {
			clicked := false
			tt := ui.NewTester(func(c *ui.Context) {
				NewButton(c, tc.name, tc.variant, th).
					Icon(symbols.Send).
					OnClick(func() {
						clicked = true
					})
			}, 300, 100)

			if !tt.HasText(tc.name) {
				t.Fatalf("element text %s not found in tester", tc.name)
			}

			if err := tt.Click(tc.name); err != nil {
				t.Fatalf("failed to click button %s: %v", tc.name, err)
			}

			if !clicked {
				t.Errorf("expected OnClick to be invoked for %s", tc.name)
			}
		})
	}
}

func TestButton_DisabledDoesNotClick(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	clicked := false
	tt := ui.NewTester(func(c *ui.Context) {
		NewButton(c, "DisabledBtn", VariantFilled, th).
			Disabled(true).
			OnClick(func() {
				clicked = true
			})
	}, 300, 100)

	_ = tt.Click("DisabledBtn")
	if clicked {
		t.Error("disabled button should not invoke OnClick")
	}
}
