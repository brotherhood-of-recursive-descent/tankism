package components

import (
	"image/color"

	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const LightType = "light"

type Light struct {
	Image *ebiten.Image
	Color color.Color
}

func (s Light) Type() ecs.ComponentType {
	return LightType
}
