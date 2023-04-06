package systems

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/co0p/tankism/game/ecs/components"
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
	particles     []Particle
	MaxParticles  int
	EntityManager *ecs.EntityManager
}

func (s *ParticleSystem) Update() error {
	entities := s.EntityManager.FindByComponents(components.ParticleEmitterType, components.TransformType)

	// position all emitters and let them emit particles
	for _, e := range entities {

		transform := e.GetComponent(components.TransformType).(*components.Transform)
		emitter := e.GetComponent(components.ParticleEmitterType).(*components.ParticleEmitter)

		now := time.Now()
		// now.Sub(emitter.Last_emitted) > emitter.Spawn_interval {

		if now.Add(-emitter.Spawn_interval).After(emitter.Last_emitted) {

			// maybe change to a ringbuffer
			if len(s.particles) < s.MaxParticles {

				vx := float64(emitter.Direction_min+rand.Intn(emitter.Direction_max-emitter.Direction_min)) * math.Pi / 180
				vy := float64(emitter.Direction_min+rand.Intn(emitter.Direction_max-emitter.Direction_min)) * math.Pi / 180
				velX := math.Cos(vx)
				velY := math.Sin(vy)

				lifetime := rand.Float64()*(float64(emitter.Lifetime_max)-float64(emitter.Lifetime_min)) + float64(emitter.Lifetime_min)

				p := Particle{
					color:    emitter.Color,
					x:        transform.Point.X,
					y:        transform.Point.Y,
					vx:       velX,
					vy:       velY,
					lifetime: lifetime,
				}

				s.particles = append(s.particles, p)
				emitter.Last_emitted = now
			}
		}
	}

	// keep 'active' particles
	ps := s.particles[:0]
	for _, p := range s.particles {

		p.x = p.x + p.vx
		p.y = p.y + p.vy
		p.lifetime--

		if p.lifetime > 0 {
			ps = append(ps, p)
		}
	}
	s.particles = ps

	return nil
}

// Draw draws all particles
func (s *ParticleSystem) Draw(screen *ebiten.Image) {

	for _, particle := range s.particles {
		ebitenutil.DrawCircle(screen, particle.x, particle.y, 5.0, particle.color)
	}
}
