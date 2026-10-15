package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

func TestSelectionControls(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	checked := false
	switched := false
	radioValue := "optionA"

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Gap(10).Children(func() {
			NewCheckbox(c, &checked, "Accept Terms", th)
			NewSwitch(c, &switched, "Airplane Mode", th)
			NewRadio(c, &radioValue, "optionA", "Option A", th)
			NewRadio(c, &radioValue, "optionB", "Option B", th)
		})
	}, 400, 300)

	// 1. Click Checkbox
	if err := tt.Click("Accept Terms"); err != nil {
		t.Fatalf("failed clicking Checkbox: %v", err)
	}
	if !checked {
		t.Error("expected checked == true after click")
	}

	// 2. Click Switch
	if err := tt.Click("Airplane Mode"); err != nil {
		t.Fatalf("failed clicking Switch: %v", err)
	}
	if !switched {
		t.Error("expected switched == true after click")
	}

	// 3. Click Radio Option B
	if err := tt.Click("Option B"); err != nil {
		t.Fatalf("failed clicking Radio Option B: %v", err)
	}
	if radioValue != "optionB" {
		t.Errorf("expected radioValue == optionB, got %s", radioValue)
	}
}
