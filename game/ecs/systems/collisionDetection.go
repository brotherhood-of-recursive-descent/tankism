package systems

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionDetection struct {
	EntityManager *ecs.EntityManager
}

func (s *CollisionDetection) Draw(screen *ebiten.Image) {}

func (s *CollisionDetection) Update() error {

	//entities := s.EntityManager.FindByComponents(components.BoundingBoxType, components.TransformType)

	/* 	for _,_e := range entities {
		//translate := e.GetComponent(components.TransformType).(*components.Transform)
		//boundingBox := e.GetComponent(components.BoundingBoxType).(*components.BoundingBox)

	} */

	return nil
}
