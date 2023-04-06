package main

import (
	"log"

	_ "embed"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/vector"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed earth.png
var assetEarth []byte

//go:embed moon.png
var assetMoon []byte

//go:embed sun.png
var assetSun []byte

//go:embed stars.png
var assetStars []byte

// PositionDemo demonstrates the relative positioning of elements using planets
// sun isParentOf( earth isParentOf( moon ) )
// once #36 is implemented, then the moon will have a light attached as well :-)
//
// press 'd' to display debug overlay
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
	centerW := float64(w / 2)
	centerH := float64(h / 2)

	lightImage, _ := resources.LoadImage(resources.LightCircle)
	lightSprite := ebiten.NewImageFromImage(lightImage)

	// fps
	performance := s.EntityManager.NewEntity()
	game.FPSCounter(performance, 1024)

	// background
	backgroundImg, _ := resources.LoadImage(assetStars)
	backgroundSprite := ebiten.NewImageFromImage(backgroundImg)

	background := s.EntityManager.NewEntity()
	background.AddComponents(
		&components.Transform{},
		&components.Sprite{Image: backgroundSprite, ZIndex: 1},
	)

	ambientLight := s.EntityManager.NewEntity()
	ambientLight.AddComponent(&components.AmbientLight{
		CompositeMode: ebiten.CompositeModeSourceOver,
		Color:         lib.Color{R: 30, G: 30, B: 30, A: 100},
	})

	// sun
	sunImg, _ := resources.LoadImage(assetSun)
	sunSprite := ebiten.NewImageFromImage(sunImg)
	sunWidth, sunHeight := sunSprite.Size()

	sun := s.EntityManager.NewEntity()
	sunTransform := components.Transform{
		Point: vector.Vec2d{
			X: centerW - float64(sunWidth/2),
			Y: centerH - float64(sunHeight/2),
		},
	}

	sun.AddComponents(
		&sunTransform,
		&components.Sprite{Image: sunSprite, ZIndex: 2},
		&components.Controller{},
		&components.Velocity{},
		&components.Debug{},
	)

	sunLight := s.EntityManager.NewEntity()
	sunLightSprite := &components.Light{Image: lightSprite, Color: lib.ColorYellow}
	sunLightTransform := &components.Transform{Scale: 2, OffsetX: float64(sunWidth/2) - 128, OffsetY: float64(sunHeight/2) - 128}
	sunLightTransform.AddParent(&sunTransform)
	sunLight.AddComponents(sunLightSprite, sunLightTransform, &components.Debug{})

	// earth
	earthImg, _ := resources.LoadImage(assetEarth)
	earthSprite := ebiten.NewImageFromImage(earthImg)
	earthWidth, earthHeight := sunSprite.Size()

	earth := s.EntityManager.NewEntity()
	earthTransform := components.Transform{OffsetX: 300, OffsetY: 0}
	earthTransform.AddParent(&sunTransform)
	earth.AddComponents(
		&earthTransform,
		&components.Sprite{Image: earthSprite, ZIndex: 2},
		&components.Debug{},
	)

	earthLight := s.EntityManager.NewEntity()
	earthLightSprite := &components.Light{Image: lightSprite, Color: lib.ColorBlue}
	earthLightTransform := &components.Transform{Scale: 2, OffsetX: float64(earthWidth/2) - 128, OffsetY: float64(earthHeight/2) - 128}
	earthLightTransform.AddParent(&earthTransform)
	earthLight.AddComponents(earthLightSprite, earthLightTransform, &components.Debug{})

	// moon
	moonImg, _ := resources.LoadImage(assetMoon)
	moonSprite := ebiten.NewImageFromImage(moonImg)
	moonWidth, moonHeight := moonSprite.Size()
	moon := s.EntityManager.NewEntity()
	moonTransform := components.Transform{OffsetX: 150, OffsetY: 150}
	moonTransform.AddParent(&earthTransform)
	moon.AddComponents(
		&moonTransform,
		&components.Sprite{Image: moonSprite, ZIndex: 2},
		&components.Debug{},
	)
	moonLight := s.EntityManager.NewEntity()
	moonLightSprite := &components.Light{Image: lightSprite, Color: lib.ColorWhite}
	moonLightTransform := &components.Transform{Scale: 2, OffsetX: float64(moonWidth/2) - 128, OffsetY: float64(moonHeight/2) - 128}
	moonLightTransform.AddParent(&moonTransform)
	moonLight.AddComponents(moonLightSprite, moonLightTransform, &components.Debug{})
}

func (s *PositionDemo) systems() {

	s.Systems = append(s.Systems,
		&systems.SpriteRenderer{EntityManager: &s.EntityManager},
		&systems.MovementSystem{EntityManager: &s.EntityManager},
		&systems.PositioningSystem{EntityManager: &s.EntityManager},
		systems.NewLightingSystem(&s.EntityManager),
		&systems.Controller{EntityManager: &s.EntityManager},
		&systems.DebugRenderer{EntityManager: &s.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &s.EntityManager},
		&systems.TextRenderer{EntityManager: &s.EntityManager},
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
