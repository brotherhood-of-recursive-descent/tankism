package components

import (
	"time"

	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

/* WIP - no animation system yet present */

const AnimationType = "sprite"

type AnimationStep struct {
	Image    *ebiten.Image
	Duration time.Duration
}

type Animation struct {
	CurrentSprite *ebiten.Image
	ZIndex        int

	Steps       []AnimationStep
	Repeat      bool
	CurrentTime time.Time
}

func (s Animation) Type() ecs.ComponentType {
	return AnimationType
}
