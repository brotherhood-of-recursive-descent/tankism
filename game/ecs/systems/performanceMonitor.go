package systems

import (
	"fmt"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type PerformanceMonitor struct {
	EntityManager *ecs.EntityManager
}

func (s *PerformanceMonitor) Draw(screen *ebiten.Image) {}

func (s *PerformanceMonitor) Update() error {

	entityCount := s.EntityManager.Size()
	entities := s.EntityManager.FindByComponents(components.FPSType, components.TextType)

	for _, e := range entities {
		text := e.GetComponent(components.TextType).(*components.Text)
		text.Value = fmt.Sprintf("fps: %6.2f (%d)", ebiten.CurrentFPS(), entityCount)
	}

	return nil
}
