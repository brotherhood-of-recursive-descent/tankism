package components

import (
	"image/color"

	"github.com/co0p/tankism/lib/ecs"
)

const ParticleEmitterType = "particleEmitter"

// Text holds all information needed to render text
type ParticleEmitter struct {
	Color        color.Color
	Position_min float64
	Position_max float64
	Velocity_min float64
	Velocity_max float64
	Lifetime_min int
	Lifetime_max int
	Spawnrate    int
}

func (t ParticleEmitter) Type() ecs.ComponentType {
	return ParticleEmitterType
}
