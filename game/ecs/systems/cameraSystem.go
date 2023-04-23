package systems

import (
	"errors"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/camera"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

// CameraSystem controls the camera by reading cameraComponent and setting values on the actual camera
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

	// find the entity with the camera attached and positioned in the world
	entities := s.EntityManager.FindByComponents(components.CameraType, components.TransformType)

	if len(entities) != 1 {
		return errors.New("expected exactly 1 entity with camera attached")
	}

	cameraComponent := entities[0].GetComponent(components.CameraType).(*components.Camera)
	transformComponent := entities[0].GetComponent(components.TransformType).(*components.Transform)

	switch cameraComponent.CameraMode {
	case camera.CameraModeDefault:
		break
	case camera.CameraModeCenter:
		// TODO(#39) continue here
		s.camera.Move(transformComponent.Point)
	}

	s.camera.SetZoom(cameraComponent.Zoom)

	return nil
}
