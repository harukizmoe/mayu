package components

import (
	"github.com/egoist/mygo/ui"
)

// StateLayer opacities defined by Material Design 3.
const (
	StateHoverAlpha   float32 = 0.08
	StateFocusAlpha   float32 = 0.10
	StatePressedAlpha float32 = 0.10
	StateDraggedAlpha float32 = 0.16

	DisabledContainerAlpha float32 = 0.12
	DisabledContentAlpha   float32 = 0.38
	DisabledOutlineAlpha   float32 = 0.12
)

// ApplyStateLayer blends the interaction state layer color over the base background color.
func ApplyStateLayer(baseColor, stateLayerColor ui.Color, hovered, focused, pressed bool) ui.Color {
	if pressed {
		if baseColor.A == 0 {
			return stateLayerColor.Alpha(StatePressedAlpha)
		}
		return stateLayerColor.Alpha(StatePressedAlpha).Over(baseColor)
	}
	if hovered {
		if baseColor.A == 0 {
			return stateLayerColor.Alpha(StateHoverAlpha)
		}
		return stateLayerColor.Alpha(StateHoverAlpha).Over(baseColor)
	}
	if focused {
		if baseColor.A == 0 {
			return stateLayerColor.Alpha(StateFocusAlpha)
		}
		return stateLayerColor.Alpha(StateFocusAlpha).Over(baseColor)
	}
	return baseColor
}
