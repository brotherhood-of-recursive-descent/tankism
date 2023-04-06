package systems

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type RotationSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *RotationSystem) Draw(screen *ebiten.Image) {}

func (s *RotationSystem) Update() error {

	// get all entities with
	// - translateComponent
	// - rotationComponent

	// calculate new x,y based on rotationComponent

	// save in new rotation and X,Y in translate

	return nil
}
