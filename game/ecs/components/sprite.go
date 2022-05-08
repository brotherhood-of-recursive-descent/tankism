package components

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const SpriteType = "sprite"

type Sprite struct {
	Image  *ebiten.Image
	ZIndex int
}

func (s Sprite) Type() ecs.ComponentType {
	return SpriteType
}
