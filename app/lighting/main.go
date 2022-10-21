package main

import (
	"image/color"
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type LightingDemo struct {
	game.GameScene

	ambientColors []color.Color
	colorIdx      int
}

func (demo *LightingDemo) Init() error {

	demo.ambientColors = append(demo.ambientColors,
		color.RGBA{60, 76, 128, 10},
		color.RGBA{60, 76, 128, 50},
		color.RGBA{60, 76, 128, 10},
		color.RGBA{60, 76, 128, 150},
		color.RGBA{60, 76, 128, 200},
		color.RGBA{255, 255, 255, 255},
		color.RGBA{100, 100, 100, 100},
		color.RGBA{150, 100, 100, 100},
		color.RGBA{150, 150, 100, 100},
		color.RGBA{150, 150, 150, 150},
		color.RGBA{200, 200, 200, 200},
	)

	demo.Systems = append(demo.Systems,
		&systems.SpriteRenderer{EntityManager: &demo.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &demo.EntityManager},
		&systems.TextRenderer{EntityManager: &demo.EntityManager},
		&systems.Controller{EntityManager: &demo.EntityManager},
		systems.NewLightingSystem(&demo.EntityManager),
	)

	fps := demo.EntityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	// background map
	tilemap := demo.EntityManager.NewEntity()
	game.NewMap(tilemap, game.Tilemap{}, 1024, 1024)

	// add some items
	barrel := demo.EntityManager.NewEntity()
	game.NewDrum(barrel, 300, 300)
	game.NewPointLight(demo.EntityManager.NewEntity(), 325, 325)

	crate := demo.EntityManager.NewEntity()
	game.NewCrate(crate, 100, 300)
	game.NewPointLight(demo.EntityManager.NewEntity(), 125, 330)

	// add different PointLight component to tanks
	circleLight := demo.EntityManager.NewEntity()
	game.NewCircleLight(circleLight, 500, 500)

	circleLightGreen := demo.EntityManager.NewEntity()
	game.NewCircleLightWithColor(circleLightGreen, 750, 750, lib.ColorGreen)

	circleLightRed := demo.EntityManager.NewEntity()
	game.NewCircleLightWithColor(circleLightRed, 600, 700, lib.ColorRed)

	circleLightBlue := demo.EntityManager.NewEntity()
	game.NewCircleLightWithColor(circleLightBlue, 650, 650, lib.ColorBlue)

	tank := demo.EntityManager.NewEntity()
	game.NewTank(tank)

	// add ambient light entity
	ambientLight := demo.EntityManager.NewEntity()
	game.NewAmbientLight(ambientLight)

	return nil
}

func (s *LightingDemo) Update() error {

	ambientLightEntity := s.EntityManager.FindByComponents(components.AmbientLightType)
	if len(ambientLightEntity) == 1 {
		ambientLight := ambientLightEntity[0].GetComponent(components.AmbientLightType).(*components.AmbientLight)
		ambientLight.Color = s.ambientColors[s.colorIdx]
	}

	return s.GameScene.Update()
}

func (s *LightingDemo) HandleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyPageDown) {
		s.colorIdx++
		s.colorIdx %= len(s.ambientColors)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyPageUp) {
		s.colorIdx += len(s.ambientColors) - 1
		s.colorIdx %= len(s.ambientColors)
	}
}

func main() {

	demo := LightingDemo{}
	game := game.NewGame()
	game.AddScene("LightDemo", &demo)
	game.SetScene("LightDemo")

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
