package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/test"
	"github.com/hajimehoshi/ebiten/v2"
)

type LoadingScene struct {
	lib.Scene
	entityManager *ecs.EntityManager
	systems       []ecs.System
}

func (s *LoadingScene) Init(sm *lib.SceneManager) error {

	state, _ := game.NewState(test.GameState_Valid)
	s.entityManager = ecs.NewEntityManager(state.Entities)

	s.systems = append(s.systems,
		&systems.SpriteRenderer{EntityManager: s.entityManager},
	)
	return nil
}

func main() {
	emptyScene := LoadingScene{}

	client := game.NewGame()
	client.AddScene("Loading", &emptyScene)

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(client); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
