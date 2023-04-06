package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/collision"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionDetection struct {
	EntityManager *ecs.EntityManager
}

func (s *CollisionDetection) Draw(screen *ebiten.Image) {}

func (s *CollisionDetection) Update() error {

	entities := s.EntityManager.FindByComponents(components.BoundingBoxType, components.TransformType)
	boundingBoxes := []collision.BoundingBox{}

	for _, entity := range entities {

		translate := entity.GetComponent(components.TransformType).(*components.Transform)
		bb := entity.GetComponent(components.BoundingBoxType).(*components.BoundingBox)

		entityBox := collision.BoundingBox{
			X:      translate.Point.X,
			Y:      translate.Point.Y,
			Width:  bb.Width * translate.Scale,
			Height: bb.Height * translate.Scale,
			E:      entity,
		}

		boundingBoxes = append(boundingBoxes, entityBox)
	}

	for i := 0; i < len(boundingBoxes); i++ {

		rect1 := boundingBoxes[i]

		for j := i + 1; j < len(boundingBoxes); j++ {

			rect2 := boundingBoxes[j]

			if collision.AABBCollision(rect1, rect2) {

				rect1.E.AddComponent(&components.Collision{Target: rect2.E})
				rect2.E.AddComponent(&components.Collision{Target: rect1.E})
			}
		}
	}
	return nil
}
