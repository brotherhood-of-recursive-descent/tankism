package main

import (
	"image/color"
	"log"

	_ "embed"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

/* TODO: WIP make transform relational to parent. ... earth light should be relative to earth */

//go:embed earth.png
var assetEarth []byte

//go:embed moon.png
var assetMoon []byte

//go:embed sun.png
var assetSun []byte

//go:embed stars.png
var assetStars []byte

type PositionDemo struct {
	game.GameScene

	game *game.Game
}

func (s *PositionDemo) Init() error {
	s.entities()
	s.systems()
	return nil
}

func (s *PositionDemo) entities() {
	w, h := s.game.WindowSize()
	centerW := w / 2
	centerH := h / 2

	// background
	backgroundImg, _ := resources.LoadImage(assetStars)
	backgroundSprite := ebiten.NewImageFromImage(backgroundImg)

	background := s.EntityManager.NewEntity()
	background.AddComponents(
		&components.Transform{X: 0, Y: 0},
		&components.Sprite{Image: backgroundSprite, ZIndex: 1},
	)

	ambientLight := s.EntityManager.NewEntity()
	ambientLight.AddComponent(&components.AmbientLight{
		CompositeMode: ebiten.CompositeModeSourceOver,
		Color:         color.RGBA{R: 60, G: 76, B: 128, A: 10},
	})

	// sun
	sunImg, _ := resources.LoadImage(assetSun)
	sunSprite := ebiten.NewImageFromImage(sunImg)
	sunWidth, sunHeight := sunSprite.Size()

	sun := s.EntityManager.NewEntity()
	sun.AddComponents(
		&components.Transform{X: float64(centerW - sunWidth/2), Y: float64(centerH - sunHeight/2)},
		&components.Sprite{Image: sunSprite, ZIndex: 2},
		&components.Debug{Color: lib.ColorYellow},
	)

	sunLight := s.EntityManager.NewEntity()
	game.NewCircleLightWithColor(sunLight, float64(centerW-128), float64(centerH-128), lib.ColorYellow)
	sunLight.AddComponent(&components.Debug{Color: lib.ColorYellow})

	// earth
	earthImg, _ := resources.LoadImage(assetEarth)
	earthSprite := ebiten.NewImageFromImage(earthImg)
	earthWidth, earthHeight := sunSprite.Size()

	earth := s.EntityManager.NewEntity()
	earth.AddComponents(
		&components.Transform{X: float64(centerW - earthWidth/2 + 300), Y: float64(h/2 - earthHeight/2)},
		&components.Sprite{Image: earthSprite, ZIndex: 2},
		&components.Debug{Color: lib.ColorGreen},
	)
	earthLight := s.EntityManager.NewEntity()
	game.NewCircleLightWithColor(earthLight, float64(centerW+300-128), float64(centerH-128), lib.ColorGreen)
	earthLight.AddComponent(&components.Debug{Color: lib.ColorYellow})

	// moon
	moonImg, _ := resources.LoadImage(assetMoon)
	moonSprite := ebiten.NewImageFromImage(moonImg)

	moon := s.EntityManager.NewEntity()
	moon.AddComponents(
		&components.Transform{X: float64(centerW + 450), Y: float64(centerH) + 100},
		&components.Sprite{Image: moonSprite, ZIndex: 2},
		&components.Debug{Color: lib.ColorBlue},
	)
	moonLight := s.EntityManager.NewEntity()
	game.NewCircleLightWithColor(moonLight, float64(centerW+450), float64(centerH+100), lib.ColorBlue)
	moonLight.AddComponent(&components.Debug{Color: lib.ColorRed})
}

func (s *PositionDemo) systems() {

	s.Systems = append(s.Systems,
		&systems.SpriteRenderer{EntityManager: &s.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &s.EntityManager},
		&systems.TextRenderer{EntityManager: &s.EntityManager},
		&systems.MovementSystem{EntityManager: &s.EntityManager},
		systems.NewLightingSystem(&s.EntityManager),
		&systems.DebugRenderer{EntityManager: &s.EntityManager},
	)
}

func main() {
	game := game.NewGame()

	demo := PositionDemo{game: game}

	game.AddScene("PositionDemo", &demo)
	game.SetScene("PositionDemo")

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
