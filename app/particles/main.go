package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleDemo struct {
	game.GameScene

	game *game.Game
}

func (p *ParticleDemo) Init() error {

	fps := p.EntityManager.NewEntity()
	game.FPSCounter(fps, p.game.ScreenWidth-200)

	// add systems
	p.Systems = append(p.Systems,
		systems.NewParticleSystem(&p.EntityManager, 100),
		&systems.TextRenderer{EntityManager: &p.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &p.EntityManager},
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
