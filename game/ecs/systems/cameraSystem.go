package systems

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type CameraSystem struct {
	EntityManager *ecs.EntityManager
	image         *ebiten.Image
}

func (s *CameraSystem) Draw(screen *ebiten.Image) {
	entities := s.EntityManager.FindByComponents(components.TransformType, components.CameraType)

	if len(entities) == 0 {
		return
	}

	if len(entities) > 1 {
		log.Fatalf("expected max of 1 entity with camera attached")
		return
	}

	camera := entities[0].GetComponent(components.CameraType).(*components.Camera)

	// lazy load
	if s.image == nil {
		s.image = ebiten.NewImageFromImage(screen)
	}

	// TODO - convert world coords to camera coords
	w, h := s.image.Size()
	window := image.Rectangle{
		Min: image.Point{int(camera.X), int(camera.Y)},
		Max: image.Point{int(camera.X) + w, int(camera.Y) + h},
	}
	fmt.Printf("window: %v\n", window)

	view := ebiten.NewImageFromImage(screen).SubImage(window).(*ebiten.Image)
	view.Fill(color.Black)
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
	transform := entities[0].GetComponent(components.TransformType).(*components.Transform)

	camera.X = transform.X
	camera.Y = transform.Y

	return nil
}
