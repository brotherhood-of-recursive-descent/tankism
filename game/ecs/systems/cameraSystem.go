package systems

import (
	"errors"
	"fmt"
	"log"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CameraSystem struct {
	EntityManager *ecs.EntityManager
	surface       *ebiten.Image
}

func NewCameraSystem(em *ecs.EntityManager, w, h int) *CameraSystem {

	return &CameraSystem{
		EntityManager: em,
		surface:       ebiten.NewImage(w, h),
	}
}

func (s *CameraSystem) Draw(screen *ebiten.Image) {

	if s.surface == nil {
		panic("CameraSystem not initialized, use NewCameraSystem()")
	}

	entities := s.EntityManager.FindByComponents(components.CameraType, components.TransformType)

	if len(entities) != 1 {
		log.Fatalf("expected exactly 1 entity with camera attached")
		return
	}

	camera := entities[0].GetComponent(components.CameraType).(*components.Camera)

	_, h := lib.WidthHeight(screen)
	x, y := camera.Point.XY()

	s.surface.Clear()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	s.surface.DrawImage(screen, op)

	screen.DrawImage(s.surface, &ebiten.DrawImageOptions{})
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("x:%v, y:%v", x, y), 100, h-100)
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
	if s.surface == nil {
		return nil
	}

	camera.Point = transformTarget.Point

	return nil
}
