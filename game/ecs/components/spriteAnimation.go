package components

import (
	"time"

	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const SpriteAnimationType = "spriteAnimation"

type SpriteAnimation struct {
	Idx        int
	Duration   time.Duration
	Images     []*ebiten.Image
	LastUpdate time.Time
}

func (s SpriteAnimation) Type() ecs.ComponentType {
	return SpriteAnimationType
}
