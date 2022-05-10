package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/shaders"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type LightingSystem struct {
	ambientShader *ebiten.Shader
	entityManager *ecs.EntityManager
}

func NewLightingSystem(entityManager *ecs.EntityManager) *LightingSystem {

	ambientShader, err := ebiten.NewShader(shaders.AmbientLightShader)
	if err != nil {
		panic("failed to load ambient Shader: " + err.Error())
	}

	return &LightingSystem{
		entityManager: entityManager,
		ambientShader: ambientShader,
	}
}

func (s *LightingSystem) Update() error { return nil }
func (s *LightingSystem) Draw(screen *ebiten.Image) {

	/*
	* Apply ambient Light
	 */
	ambientLightEntity := s.entityManager.FindByComponents(components.AmbientLightType)
	if len(ambientLightEntity) > 0 {
		ambientLight := ambientLightEntity[0].GetComponent(components.AmbientLightType).(*components.AmbientLight)
		if ambientLight.Active {
			r, g, b, a := ambientLight.Color.RGBA()
			ambientColor := shaders.Vec4(r, g, b, a)
			w, h := screen.Size()
			op := &ebiten.DrawRectShaderOptions{}
			op.CompositeMode = ambientLight.CompositeMode
			op.Uniforms = map[string]interface{}{"AmbientColor": ambientColor}
			screen.DrawRectShader(w, h, s.ambientShader, op)
		}
	}

	/*
	* apply rest of lighting
	 */
	entities := s.entityManager.FindByComponents(components.LightType, components.TranslateType)

	for _, e := range entities {

		light := e.GetComponent(components.LightType).(*components.Light)
		translate := e.GetComponent(components.TranslateType).(*components.Translate)
		rect := light.Image.Bounds()

		x := translate.X
		y := translate.Y
		rotation := light.Rotation
		scale := light.Scale

		op := &ebiten.DrawImageOptions{}
		op.CompositeMode = light.CompositeMode
		op.GeoM.Translate(-float64(rect.Dx())/2, -float64(rect.Dy())/2)
		op.GeoM.Scale(scale, scale)
		op.GeoM.Rotate(rotation)
		op.GeoM.Translate(x, y)

		img := light.Image
		screen.DrawImage(img, op)
	}
}
