package systems

import (
	"math"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type AISystem struct {
	EntityManager *ecs.EntityManager
}

func (s *AISystem) Draw(screen *ebiten.Image) {}

func (s *AISystem) Update() error {

	aiEntities := s.EntityManager.FindByComponents(components.AIType, components.TranslateType)
	targetEntities := s.EntityManager.FindByComponents(components.TargetType, components.TranslateType)

	for _, aie := range aiEntities {
		ai := aie.GetComponent(components.AIType).(*components.AI)
		aiTranslate := aie.GetComponent(components.TranslateType).(*components.Transform)

		for _, e := range targetEntities {
			target := e.GetComponent(components.TargetType).(*components.Target)
			targetTranslate := e.GetComponent(components.TranslateType).(*components.Transform)

			if ai.TargetGroup == target.GroupId {

				dx := aiTranslate.X - targetTranslate.X
				dy := aiTranslate.Y - targetTranslate.Y

				// -dy because y increases 'downwards'
				targetRad := math.Atan2(dx, -dy)
				aiTranslate.Rotation = targetRad
			}
		}

	}

	return nil
}
