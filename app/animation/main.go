package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/resource"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type AnimationDemo struct {
	WindowWidth, WindowHeight int

	entityManager *ecs.EntityManager
	systems       []ecs.System
	spriteSheet   *resource.SpriteSheet
}

func NewAnimationDemo() *AnimationDemo {

	scene := AnimationDemo{
		systems:       []ecs.System{},
		entityManager: &ecs.EntityManager{},
	}

	scene.systems = append(scene.systems,
		&systems.SpriteRenderer{EntityManager: scene.entityManager},
		&systems.PerformanceMonitor{EntityManager: scene.entityManager},
		&systems.TextRenderer{EntityManager: scene.entityManager},
		&systems.SpriteAnimator{EntityManager: scene.entityManager},
		systems.NewLightingSystem(scene.entityManager),
	)

	return &scene
}

func (s *AnimationDemo) Init() error {
	fps := s.entityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	// background map
	tilemap := s.entityManager.NewEntity()
	game.NewMap(tilemap, game.Tilemap{}, 1024, 1024)

	// BOOM
	spriteSheet, err := resource.NewSpriteSheetFromConfig(media.AllSprites, media.AllSpritesConfig)
	if err != nil {
		panic("failed to load sprite sheet and or config")
	}
	s.spriteSheet = &spriteSheet

	return nil
}

func (s *AnimationDemo) Draw(screen *ebiten.Image) {
	for _, v := range s.systems {
		v.Draw(screen)
	}
}

func (s *AnimationDemo) Update() error {
	var err error
	for _, v := range s.systems {
		err = v.Update()
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		e := s.entityManager.NewEntity()
		x, y := ebiten.CursorPosition()
		game.NewExplosion(e, *s.spriteSheet, x-60, y-60)
	}

	return err
}

func (s *AnimationDemo) Layout(outsideWidth, outsideHeight int) (int, int) {
	s.WindowWidth = outsideWidth
	s.WindowHeight = outsideHeight
	return outsideWidth, outsideHeight
}

func main() {

	ebiten.SetFullscreen(true)
	client := NewAnimationDemo()
	client.Init()

	if err := ebiten.RunGame(client); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
