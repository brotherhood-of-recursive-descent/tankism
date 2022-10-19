package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/resource"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type AnimationDemo struct {
	game.Scene

	explosionSprites *resource.SpriteSheet
	lights           *resource.SpriteSheet
}

func (s *AnimationDemo) Init(sm *lib.SceneManager) error {

	s.Systems = append(s.Systems,
		&systems.CleanupSystem{EntityManager: &s.EntityManager},
		&systems.SpriteRenderer{EntityManager: &s.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &s.EntityManager},
		&systems.TextRenderer{EntityManager: &s.EntityManager},
		&systems.SpriteAnimator{EntityManager: &s.EntityManager},
		&systems.TriggerSystem{EntityManager: &s.EntityManager},
		systems.NewLightingSystem(&s.EntityManager),
	)

	fps := s.EntityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	// ambient light
	ambient := s.EntityManager.NewEntity()
	game.NewAmbientLight(ambient)

	// background map
	tilemap := s.EntityManager.NewEntity()
	game.NewMap(tilemap, game.Tilemap{}, 1024, 1024)

	// lights
	lights := game.NewLightSpritesheet()
	s.lights = &lights

	// BOOM
	explosionSprites, err := resource.NewSpriteSheetFromConfig(resources.AllSprites, resources.AllSpritesConfig)
	if err != nil {
		panic("failed to load sprite sheet and or config")
	}
	s.explosionSprites = &explosionSprites

	return nil
}

func (s *AnimationDemo) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		e := s.EntityManager.NewEntity()
		x, y := ebiten.CursorPosition()
		game.NewExplosion(e, *s.explosionSprites, *s.lights, x-60, y-60)
	}
	return s.Scene.Update()
}

func main() {

	client := game.NewGame()
	demo := AnimationDemo{}
	client.AddScene("AnimationDemo", &demo)

	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(client); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
