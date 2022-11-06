package components

import (
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const AmbientLightType = "ambientLight"

type AmbientLight struct {
	CompositeMode ebiten.CompositeMode
	Color         lib.Color
	Active        bool
}

func (s AmbientLight) Type() ecs.ComponentType {
	return AmbientLightType
}
