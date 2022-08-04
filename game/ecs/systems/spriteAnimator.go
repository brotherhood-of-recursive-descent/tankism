package systems

import (
	"time"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteAnimator struct {
	EntityManager *ecs.EntityManager
}

func (s *SpriteAnimator) Draw(screen *ebiten.Image) {}

func (s *SpriteAnimator) Update() error {

	now := time.Now()
	entities := s.EntityManager.FindByComponents(components.SpriteAnimationType, components.SpriteType)

	for _, e := range entities {
		spriteAnimation := e.GetComponent(components.SpriteAnimationType).(*components.SpriteAnimation)
		sprite := e.GetComponent(components.SpriteType).(*components.Sprite)

		lastUpdate := spriteAnimation.LastUpdate
		duration := spriteAnimation.Duration
		idx := spriteAnimation.Idx

		// lazy init
		if lastUpdate.IsZero() {
			lastUpdate = now
			spriteAnimation.LastUpdate = now
		}

		// out of bounds
		if idx == len(spriteAnimation.Images)-1 {
			sprite.Image = nil
			continue
		}

		nextUpdate := lastUpdate.Add(duration)
		if now.After(nextUpdate) {
			sprite.Image = spriteAnimation.Images[idx]
			spriteAnimation.LastUpdate = nextUpdate
			spriteAnimation.Idx++
		}

	}

	return nil
}
