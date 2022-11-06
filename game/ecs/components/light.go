package components

import (
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const LightType = "light"

type Light struct {
	Image *ebiten.Image
	Color lib.Color
}

func (s Light) Type() ecs.ComponentType {
	return LightType
}
