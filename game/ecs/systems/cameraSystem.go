package systems

import (
	"errors"
	"image"
	"log"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type CameraSystem struct {
	EntityManager *ecs.EntityManager
	view          *ebiten.Image
}

func (s *CameraSystem) Draw(screen *ebiten.Image) {
	entities := s.EntityManager.FindByComponents(components.CameraType, components.TransformType)

	if len(entities) != 1 {
		log.Fatalf("expected exactly 1 entity with camera attached")
		return
	}

	camera := entities[0].GetComponent(components.CameraType).(*components.Camera)

	// lazy load
	if s.view == nil {
		s.view = ebiten.NewImageFromImage(screen)
	}

	// TODO - convert world coords to camera coords
	w, h := lib.WidthHeight(s.view)
	window := image.Rectangle{
		Min: image.Point{int(camera.X), int(camera.Y)},
		Max: image.Point{int(camera.X) + w, int(camera.Y) + h},
	}

	view := ebiten.NewImageFromImage(screen).SubImage(window).(*ebiten.Image)
	screen.DrawImage(view, &ebiten.DrawImageOptions{})
}

func (s *CameraSystem) Update() error {

	entities := s.EntityManager.FindByComponents(components.TransformType, components.CameraType)

	if len(entities) == 0 {
		return nil
	}

	if len(entities) > 1 {
		return errors.New("expected max of 1 entity with camera attached")
	}

	camera := entities[0].GetComponent(components.CameraType).(*components.Camera)

	// world coordinates
	transformTarget := entities[0].GetComponent(components.TransformType).(*components.Transform)

	// bounce checks for map, ignore for now
	// target x < view width / 2
	// do not adjust camera x

	// target x + view width / 2 > view width
	// do not adjust camera x

	// target y < view height / 2
	// do not adjust camera y

	// target y + view height / 2 > view height
	// do not adjust camera y

	// avoid race condition when view has not been initialized by first update call
	if s.view == nil {
		return nil
	}

	width, height := lib.WidthHeight(s.view)
	camera.X = transformTarget.Point.X - float64(width/2)
	camera.Y = transformTarget.Point.Y - float64(height/2)

	return nil
}
