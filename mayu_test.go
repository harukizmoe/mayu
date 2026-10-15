package mayu_test

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu"
	"github.com/harukizmoe/mayu/symbols"
)

func TestMayu_UseAndButtonAPI(t *testing.T) {
	submitClicked := false
	cancelClicked := false

	tt := ui.NewTester(func(c *ui.Context) {
		mayu.Use(c, mayu.ThemeOptions{
			SeedColor: ui.Hex("#6750A4"),
			Dark:      false,
			Density:   mayu.DensityStandard,
		})

		ui.Row(c).Gap(12).Children(func() {
			mayu.FilledButton(c, "Submit").
				Icon(symbols.Send).
				OnClick(func() {
					submitClicked = true
				})

			mayu.OutlinedButton(c, "Cancel").
				OnClick(func() {
					cancelClicked = true
				})
		})
	}, 400, 200)

	if !tt.HasText("Submit") || !tt.HasText("Cancel") {
		t.Fatal("buttons not found in tester view")
	}

	if err := tt.Click("Submit"); err != nil {
		t.Fatalf("failed clicking Submit: %v", err)
	}
	if !submitClicked {
		t.Error("expected Submit button to invoke OnClick")
	}

	if err := tt.Click("Cancel"); err != nil {
		t.Fatalf("failed clicking Cancel: %v", err)
	}
	if !cancelClicked {
		t.Error("expected Cancel button to invoke OnClick")
	}
}
