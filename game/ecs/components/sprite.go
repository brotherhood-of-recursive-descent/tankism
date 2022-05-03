package components

import (
	"time"

	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const SpriteType = "sprite"

type Step struct {
	Image    *ebiten.Image
	Duration time.Duration
}

type Sprite struct {
	Image       *ebiten.Image
	Steps       []Step
	CurrentStep Step
}

func (s Sprite) Type() ecs.ComponentType {
	return SpriteType
}
