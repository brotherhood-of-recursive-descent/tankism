package systems

import (
	"errors"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	camera "github.com/melonfunction/ebiten-camera"
)

type CameraSystem struct {
	EntityManager *ecs.EntityManager
	camera        *camera.Camera
}

func NewCameraSystem(em *ecs.EntityManager, cam *camera.Camera) *CameraSystem {

	return &CameraSystem{
		EntityManager: em,
		camera:        cam,
	}
}

func (s *CameraSystem) Draw(screen *ebiten.Image) {}

func (s *CameraSystem) Update() error {

	entities := s.EntityManager.FindByComponents(components.TransformType, components.CameraType)

	if len(entities) != 1 {
		return errors.New("expected exactly 1 entity with camera attached")
	}

	cameraComponent := entities[0].GetComponent(components.CameraType).(*components.Camera)
	transformComponent := entities[0].GetComponent(components.TransformType).(*components.Transform)

	x, y := transformComponent.Point.XY()

	s.camera.SetZoom(cameraComponent.Zoom)

	switch cameraComponent.CameraMode {
	case components.CameraModeDefault:
	case components.CameraModeCenter:
	}
	s.camera.SetPosition(x, y)

	return nil
}
