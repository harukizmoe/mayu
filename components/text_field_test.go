package components

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/harukizmoe/mayu/theme"
)

func TestTextField_Render(t *testing.T) {
	scheme := theme.SchemeFromSeed(ui.Hex("#6750A4"), false)
	th := theme.NewTheme(scheme, theme.DensityStandard)

	filledVal := "Alice"
	outlinedVal := ""

	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Gap(12).Children(func() {
			NewTextField(c, &filledVal, "Username", TextFieldFilled, th)
			NewTextField(c, &outlinedVal, "Email", TextFieldOutlined, th).
				Placeholder("user@example.com")
		})
	}, 400, 300)

	if !tt.HasText("Username") {
		t.Fatal("filled text field label 'Username' not found")
	}
	if !tt.HasText("Email") {
		t.Fatal("outlined text field label 'Email' not found")
	}
	if filledVal != "Alice" {
		t.Fatalf("expected filledVal == Alice, got %s", filledVal)
	}
}
