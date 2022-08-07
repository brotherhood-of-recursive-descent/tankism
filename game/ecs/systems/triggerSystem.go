package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type TriggerSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *TriggerSystem) Draw(screen *ebiten.Image) {}

func (s *TriggerSystem) Update() error {

	entities := s.EntityManager.FindByComponents(components.TriggerType)

	for _, e := range entities {
		trigger := e.GetComponent(components.TriggerType).(*components.Trigger)
		trigger.Action(e, s.EntityManager)
	}

	return nil
}
