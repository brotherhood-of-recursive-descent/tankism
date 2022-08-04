package game

import (
	"time"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/resource"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewExplosion(e *ecs.Entity, s resource.SpriteSheet, x, y int) *ecs.Entity {

	translate := components.Translate{
		X:        float64(x),
		Y:        float64(y),
		Scale:    1,
		Rotation: 0,
	}

	spriteAnimation := components.SpriteAnimation{
		Images: []*ebiten.Image{
			s.ByName("explosion1.png"),
			s.ByName("explosion2.png"),
			s.ByName("explosion3.png"),
			s.ByName("explosion4.png"),
			s.ByName("explosion5.png"),
		},
		Idx:      0,
		Duration: time.Millisecond * 76,
	}

	sprite := components.Sprite{
		Image:  nil,
		ZIndex: 100,
	}

	e.AddComponents(&translate, &sprite, &spriteAnimation)
	return e
}
