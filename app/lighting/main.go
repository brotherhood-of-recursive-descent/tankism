package main

import (
	"image/color"
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type LightingDemo struct {
	WindowWidth, WindowHeight int

	entityManager *ecs.EntityManager
	systems       []ecs.System
	ambientColors []color.Color
	colorIdx      int
}

func NewLightingDemo() *LightingDemo {

	scene := LightingDemo{
		systems:       []ecs.System{},
		entityManager: &ecs.EntityManager{},
		colorIdx:      0,
	}
	scene.ambientColors = append(scene.ambientColors,
		color.RGBA{255, 255, 255, 255},
		color.RGBA{100, 100, 100, 100},
		color.RGBA{150, 100, 100, 100},
		color.RGBA{150, 150, 100, 100},
		color.RGBA{150, 150, 150, 150},
		color.RGBA{200, 200, 200, 200},
	)

	scene.systems = append(scene.systems,
		&systems.SpriteRenderer{EntityManager: scene.entityManager},
		&systems.PerformanceMonitor{EntityManager: scene.entityManager},
		&systems.TextRenderer{EntityManager: scene.entityManager},
		&systems.Controller{EntityManager: scene.entityManager},
		systems.NewLightingSystem(scene.entityManager),
	)

	return &scene
}

func (s *LightingDemo) Init() error {
	fps := s.entityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	// background map
	tilemap := s.entityManager.NewEntity()
	game.NewMap(tilemap, game.Tilemap{}, 1024, 1024)

	// add some items
	barrel := s.entityManager.NewEntity()
	game.NewDrum(barrel, 300, 300)
	game.NewPointLight(s.entityManager.NewEntity(), 325, 325)

	crate := s.entityManager.NewEntity()
	game.NewCrate(crate, 100, 300)
	game.NewPointLight(s.entityManager.NewEntity(), 125, 330)

	// add different PointLight component to tanks
	circleLight := s.entityManager.NewEntity()
	game.NewCircleLight(circleLight, 500, 500)

	tank := s.entityManager.NewEntity()
	game.NewTank(tank)

	// add ambient light entity
	ambientLight := s.entityManager.NewEntity()
	game.NewAmbientLight(ambientLight)

	return nil
}

func (s *LightingDemo) Draw(screen *ebiten.Image) {
	for _, v := range s.systems {
		v.Draw(screen)
	}
}

func (s *LightingDemo) Update() error {
	var err error
	for _, v := range s.systems {
		err = v.Update()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyPageDown) {
		s.colorIdx++
		s.colorIdx %= len(s.ambientColors)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyPageUp) {
		s.colorIdx += len(s.ambientColors) - 1
		s.colorIdx %= len(s.ambientColors)
	}

	ambientLightEntity := s.entityManager.FindByComponents(components.AmbientLightType)
	if len(ambientLightEntity) == 1 {
		ambientLight := ambientLightEntity[0].GetComponent(components.AmbientLightType).(*components.AmbientLight)
		ambientLight.Color = s.ambientColors[s.colorIdx]
	}

	return err
}

func (s *LightingDemo) Layout(outsideWidth, outsideHeight int) (int, int) {
	s.WindowWidth = outsideWidth
	s.WindowHeight = outsideHeight
	return outsideWidth, outsideHeight
}

func main() {

	ebiten.SetFullscreen(true)
	client := NewLightingDemo()
	client.Init()

	if err := ebiten.RunGame(client); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
