package main

import (
	"log"
	"time"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleDemo struct {
	game.GameScene

	game *game.Game
}

func (p *ParticleDemo) Init() error {

	// perf monitoring
	fps := p.EntityManager.NewEntity()
	game.FPSCounter(fps, p.game.ScreenWidth-200)

	// our emitter
	redEmitter := p.EntityManager.NewEntity()
	redEmitter.AddComponent(&components.ParticleEmitter{
		Color:        lib.ColorRed,
		Velocity_min: 1.0,
		Velocity_max: 20.0,
		Lifetime_min: 10,
		Lifetime_max: 100,

		Velocity:      1,
		Direction_min: 0,
		Direction_max: 90,
	})
	redEmitter.AddComponent(&components.Transform{X: 300, Y: 300})
	redEmitter.AddComponent(&components.Debug{})

	blueEmitter := p.EntityManager.NewEntity()
	blueEmitter.AddComponent(&components.ParticleEmitter{
		Color:          lib.ColorBlue,
		Velocity_min:   0.1,
		Velocity_max:   1.1,
		Lifetime_min:   10,
		Lifetime_max:   100,
		Spawn_interval: time.Duration(1 * time.Second),
		Velocity:       1,
		Direction_min:  0,
		Direction_max:  360,
	})

	blueEmitter.AddComponent(&components.Transform{X: 400, Y: 500})
	blueEmitter.AddComponent(&components.Debug{})

	greenEmitter := p.EntityManager.NewEntity()
	greenEmitter.AddComponent(&components.ParticleEmitter{
		Color:          lib.ColorGreen,
		Velocity_min:   0.2,
		Velocity_max:   2.1,
		Lifetime_min:   10,
		Lifetime_max:   100,
		Spawn_interval: time.Duration(100 * time.Nanosecond),
		Velocity:       1,
		Direction_min:  80,
		Direction_max:  90,
	})
	greenEmitter.AddComponent(&components.Transform{X: 600, Y: 500})
	greenEmitter.AddComponent(&components.Debug{})

	yellowEmitter := p.EntityManager.NewEntity()
	yellowEmitter.AddComponent(&components.ParticleEmitter{
		Color:          lib.ColorYellow,
		Velocity_min:   0.1,
		Velocity_max:   1.1,
		Lifetime_min:   10,
		Lifetime_max:   100,
		Spawn_interval: time.Duration(500 * time.Nanosecond),
		Last_emitted:   time.Now(),
		Velocity:       1,
		Direction_min:  110,
		Direction_max:  180,
	})
	yellowEmitter.AddComponent(&components.Transform{X: 600, Y: 800})
	yellowEmitter.AddComponent(&components.Debug{})

	// add systems
	p.Systems = append(p.Systems,
		&systems.ParticleSystem{
			EntityManager: &p.EntityManager,
			MaxParticles:  100,
		},
		&systems.TextRenderer{EntityManager: &p.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &p.EntityManager},
		&systems.DebugRenderer{EntityManager: &p.EntityManager},
	)

	return nil
}

func main() {
	game := game.NewGame()

	demo := ParticleDemo{game: game}
	game.AddScene("ParticleDemo", &demo)
	game.SetScene("ParticleDemo")

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
