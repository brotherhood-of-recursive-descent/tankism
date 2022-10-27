package game

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/resources"
)

func FPSCounter(fps *ecs.Entity, width int) {
	fps.AddComponent(&components.Text{
		Value: "0",
		Font:  resources.FontMedium,
		Color: lib.ColorGreen,
	})
	fps.AddComponent(&components.Transform{
		X:        float64(width - 120),
		Y:        50.0,
		Scale:    1,
		Rotation: 0.0,
	})
	fps.AddComponent(&components.FPS{})
}
