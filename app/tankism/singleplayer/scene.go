package singleplayer

import (
	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/systems"
)

type SinglePlayerScene struct {
	game.GameScene

	game *game.Game
}

func NewSinglePlayerScene(game *game.Game) *SinglePlayerScene {
	return &SinglePlayerScene{
		game: game,
	}
}

func (s *SinglePlayerScene) Init() error {

	s.Systems = append(s.Systems,
		&systems.SpriteRenderer{EntityManager: &s.EntityManager},
		systems.NewLightingSystem(&s.EntityManager),
		&systems.AISystem{EntityManager: &s.EntityManager},
		&systems.TextRenderer{EntityManager: &s.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &s.EntityManager},
		&systems.Controller{EntityManager: &s.EntityManager},
		&systems.AISystem{EntityManager: &s.EntityManager},
	)

	ambientLight := s.EntityManager.NewEntity()
	configureAmbientLight(ambientLight)

	tank := s.EntityManager.NewEntity()
	configureTank(tank)

	bigTank := s.EntityManager.NewEntity()
	configureAITank(bigTank)

	w, h := s.game.WindowSize()

	fpsCounter := s.EntityManager.NewEntity()
	game.FPSCounter(fpsCounter, w)

	tilemap := s.EntityManager.NewEntity()
	game.NewMap(tilemap, game.Tilemap{}, w, h)
	return nil
}
