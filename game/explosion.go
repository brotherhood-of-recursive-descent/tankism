package game

import (
	"time"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/resource"
	"github.com/co0p/tankism/lib/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewExplosion(e *ecs.Entity, s resource.SpriteSheet, l resource.SpriteSheet, x, y int) *ecs.Entity {
	created := time.Now()
	duration := time.Millisecond * 67
	removeAction := func(e *ecs.Entity, em *ecs.EntityManager) {
		if time.Now().After(created.Add(duration * 4)) {
			e.AddComponent(&components.Cleanup{})
		}
	}
	trigger := components.Trigger{Action: removeAction}

	spriteAnimation := components.SpriteAnimation{
		Images: []*ebiten.Image{
			s.ByName("explosion1.png"),
			s.ByName("explosion2.png"),
			s.ByName("explosion3.png"),
			s.ByName("explosion4.png"),
			s.ByName("explosion5.png"),
		},
		Idx:      0,
		Duration: duration,
	}

	light := components.Light{
		Image: l.ByName("0"),
		Color: lib.ColorRed,
	}

	sprite := components.Sprite{
		Image:  nil,
		ZIndex: 100,
	}

	translate := components.Transform{
		Point: vector.Vec2d{
			X: float64(x),
			Y: float64(y),
		},
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(&translate, &sprite, &spriteAnimation, &trigger, &light)
	return e
}
