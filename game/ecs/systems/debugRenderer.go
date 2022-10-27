package systems

import (
	"image/color"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var LightColor = color.RGBA{100, 100, 200, 160}
var SpriteColor = color.RGBA{200, 100, 100, 160}

// DebugRenderer renders transparent overlays for debug purposes
type DebugRenderer struct {
	EntityManager *ecs.EntityManager
	active        bool
}

func (s *DebugRenderer) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		s.active = !s.active
	}
	return nil
}

func (s *DebugRenderer) Draw(screen *ebiten.Image) {

	if !s.active {
		return
	}

	entities := s.EntityManager.FindByComponents(components.DebugType, components.SpriteType, components.TranslateType)
	for _, e := range entities {

		sprite := e.GetComponent(components.SpriteType).(*components.Sprite)
		originalImg := sprite.Image
		if originalImg == nil {
			continue
		}
		w, h := originalImg.Size()
		img := ebiten.NewImage(w, h)
		img.Fill(SpriteColor)

		transform := e.GetComponent(components.TranslateType).(*components.Transform)
		x := transform.X
		y := transform.Y
		rotation := transform.Rotation

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(rotation)
		op.GeoM.Translate(float64(w)/2, float64(h)/2)
		op.GeoM.Translate(x, y)

		screen.DrawImage(img, op)
	}

	lightEntities := s.EntityManager.FindByComponents(components.DebugType, components.LightType, components.TranslateType)
	for _, e := range lightEntities {

		light := e.GetComponent(components.LightType).(*components.Light)
		originalImg := light.Image
		if originalImg == nil {
			continue
		}
		w, h := originalImg.Size()
		img := ebiten.NewImage(w, h)
		img.Fill(LightColor)

		transform := e.GetComponent(components.TranslateType).(*components.Transform)
		x := transform.X
		y := transform.Y
		rotation := transform.Rotation

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(rotation)
		op.GeoM.Translate(float64(w)/2, float64(h)/2)
		op.GeoM.Translate(x, y)

		screen.DrawImage(img, op)
	}
}
