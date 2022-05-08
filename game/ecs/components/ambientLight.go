package components

import (
	"image/color"

	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const AmbientLightType = "ambientLight"

type AmbientLight struct {
	CompositeMode ebiten.CompositeMode
	Color         color.Color
}

func (s AmbientLight) Type() ecs.ComponentType {
	return AmbientLightType
}
