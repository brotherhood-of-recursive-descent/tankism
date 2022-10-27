package systems

import (
	"image/color"

	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Particle struct {
	x, y     float64
	vx, vy   float64
	lifetime float64
	color    color.Color
}

type ParticleSystem struct {
	particles     []*Particle
	EntityManager *ecs.EntityManager
}

func NewParticleSystem(em *ecs.EntityManager, particleCount int) *ParticleSystem {

	s := &ParticleSystem{
		EntityManager: em,
	}

	s.particles = append(s.particles,
		&Particle{
			x:        100,
			y:        100,
			vx:       0.2,
			vy:       0.3,
			lifetime: 1000,
			color:    lib.ColorRed,
		},
		&Particle{
			x:        130,
			y:        130,
			vx:       0.2,
			vy:       -0.3,
			lifetime: 2000,
			color:    lib.ColorBlue,
		},
		&Particle{
			x:        120,
			y:        100,
			vx:       -0.2,
			vy:       0.3,
			lifetime: 100,
			color:    lib.ColorYellow,
		})

	return s
}

func (s *ParticleSystem) Update() error {
	//entities := s.EntityManager.FindByComponents(components.ParticleEmitterType, components.TranslateType)

	// position all emitters and let them emit particles
	/*for _, e := range entities {

		translate := e.GetComponent(components.TranslateType).(*components.Transform)
		emitter := e.GetComponent(components.ParticleEmitterType).(*components.ParticleEmitter)

		fmt.Printf("placed emitter at (%v,%v) with %v\n", translate.X, translate.Y, emitter)
	}*/

	// update position and lifetime of all particles
	for _, p := range s.particles {
		p.x = p.x + p.vx
		p.y = p.y + p.vy
		//particle.lifetime--
	}
	return nil
}

// Draw draws all particles
func (s *ParticleSystem) Draw(screen *ebiten.Image) {

	for _, particle := range s.particles {
		ebitenutil.DrawCircle(screen, particle.x, particle.x, 5.0, particle.color)
	}

}
