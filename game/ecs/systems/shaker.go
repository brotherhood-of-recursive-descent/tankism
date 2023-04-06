package systems

import (
	"math/rand"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Shaker struct {
	EntityManager *ecs.EntityManager
}

func (s *Shaker) Draw(screen *ebiten.Image) {}

func (s *Shaker) Update() error {

	entities := s.EntityManager.FindByComponents(components.ShakingType, components.TransformType)

	for _, e := range entities {

		if rand.Intn(50) > 15 {
			continue
		}

		val := rand.Intn(2)
		xOrY := rand.Intn(2)
		minus := rand.Intn(2)
		dir := 0

		translate := e.GetComponent(components.TransformType).(*components.Transform)
		if minus == 0 {
			dir = 1
		} else {
			dir = -1
		}

		if xOrY == 0 {
			translate.Point.X = translate.Point.X + float64(val*dir)
		} else {
			translate.Point.Y = translate.Point.Y + float64(val*dir)
		}
	}

	return nil
}
