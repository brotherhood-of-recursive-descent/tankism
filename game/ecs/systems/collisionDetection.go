package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionDetection struct {
	EntityManager *ecs.EntityManager
}

func (s *CollisionDetection) Draw(screen *ebiten.Image) {}

func (s *CollisionDetection) Update() error {

	entities := s.EntityManager.FindByComponents(components.BoundingBoxType, components.TransformType)
	boundingBoxes := []boundingBox{}

	for _, entity := range entities {

		entityPos := entity.GetComponent(components.TransformType).(*components.Transform)
		entityDim := entity.GetComponent(components.BoundingBoxType).(*components.BoundingBox)
		entityBox := boundingBox{
			x:      entityPos.X,
			y:      entityPos.Y,
			width:  entityDim.Width * entityPos.Scale,
			height: entityDim.Height * entityPos.Scale,
			e:      entity,
		}

		boundingBoxes = append(boundingBoxes, entityBox)
	}

	for i := 0; i < len(boundingBoxes); i++ {

		rect1 := boundingBoxes[i]

		for j := i + 1; j < len(boundingBoxes); j++ {

			rect2 := boundingBoxes[j]

			if rect1.AABBCollision(rect2) {

				rect1.e.AddComponent(&components.Collision{Target: rect2.e})
				rect2.e.AddComponent(&components.Collision{Target: rect1.e})
			}
		}
	}
	return nil
}

type boundingBox struct {
	x      float64
	y      float64
	width  float64
	height float64
	e      *ecs.Entity
}

// TODO: replace with Separating Axis Theorem to accomodate for rotation of rectangles
// TODO: easier to grok: https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect and do this
// for each line of a rectacle vs all lines of the other
func (rect1 *boundingBox) AABBCollision(rect2 boundingBox) bool {
	return rect1.x < rect2.x+rect2.width &&
		rect1.x+rect1.width > rect2.x &&
		rect1.y < rect2.y+rect2.height &&
		rect1.y+rect1.height > rect2.y
}
