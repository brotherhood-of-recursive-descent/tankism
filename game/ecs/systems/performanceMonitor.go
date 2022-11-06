package systems

import (
	"fmt"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type PerformanceMonitor struct {
	EntityManager *ecs.EntityManager
}

func (s *PerformanceMonitor) Draw(screen *ebiten.Image) {

	entities := s.EntityManager.FindByComponents(components.PerformanceType)
	for _, e := range entities {
		perf := e.GetComponent(components.PerformanceType).(*components.Performance)
		clrFps := lib.ColorGreen
		clrFps.A = 150

		clrGraph := lib.ColorRed
		clrGraph.A = 150
		graphHeight := 50.0
		if perf.ShowGraph {
			offsetX := 100.0
			offsetY := 40.0

			// draw values
			maxFps := max(&perf.PastFPS)
			for k, v := range perf.PastFPS {
				normalized := (v) / (maxFps) * graphHeight
				ebitenutil.DrawLine(
					screen,
					offsetX+float64(k),
					offsetY+graphHeight,
					offsetX+float64(k),
					normalized,
					clrFps)
			}

			// draw Graph
			ebitenutil.DrawLine(screen, offsetX, offsetY, offsetX, offsetY+graphHeight, clrGraph)
			ebitenutil.DrawLine(screen, offsetX, offsetY+graphHeight, offsetX+float64(perf.HistoryLength), offsetY+graphHeight, clrGraph)

		}
	}

}

func (s *PerformanceMonitor) Update() error {

	entityCount := s.EntityManager.Size()
	entities := s.EntityManager.FindByComponents(components.PerformanceType, components.TextType)

	for _, e := range entities {
		perf := e.GetComponent(components.PerformanceType).(*components.Performance)
		text := e.GetComponent(components.TextType).(*components.Text)

		finalText := ""
		if perf.ShowFPS {
			finalText = fmt.Sprintf("%sFps: %6.2f\n", finalText, ebiten.ActualFPS())
			addToHistory(&perf.PastFPS, ebiten.ActualFPS(), perf.HistoryLength)
		}
		if perf.ShowTPS {
			finalText = fmt.Sprintf("%sTps: %6.2f\n", finalText, ebiten.ActualFPS())
		}
		if perf.ShowEntityCount {
			finalText = fmt.Sprintf("%sEntities: %d\n", finalText, entityCount)
		}

		text.Value = finalText
	}

	return nil
}

func addToHistory(history *[]float64, value float64, maxLength int) {

	if len(*history) > maxLength {
		*history = (*history)[1:maxLength]
	}
	*history = append(*history, value)
}

func max(values *[]float64) float64 {
	var m float64
	for _, v := range *values {
		if v > m {
			m = v
		}
	}
	return m
}
