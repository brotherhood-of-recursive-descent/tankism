package systems

import (
	"image/color"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/shaders"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type LightingSystem struct {
	ambientShader *ebiten.Shader
	entityManager *ecs.EntityManager

	lightingTexture *ebiten.Image
}

func NewLightingSystem(entityManager *ecs.EntityManager) *LightingSystem {

	ambientShader, err := ebiten.NewShader(shaders.LightingShader)
	if err != nil {
		panic("failed to load ambient Shader: " + err.Error())
	}

	return &LightingSystem{
		entityManager:   entityManager,
		ambientShader:   ambientShader,
		lightingTexture: nil,
	}
}

func (s *LightingSystem) Update() error { return nil }
func (s *LightingSystem) Draw(screen *ebiten.Image) {
	if s.lightingTexture == nil {
		s.lightingTexture = ebiten.NewImageFromImage(screen)
	}
	s.lightingTexture.Fill(color.Black)

	ar, ag, ab, aa := color.White.RGBA()

	// ambient light
	aEntity := s.entityManager.FindByComponents(components.AmbientLightType)
	if len(aEntity) == 1 {
		ambientLight := aEntity[0].GetComponent(components.AmbientLightType).(*components.AmbientLight)
		ar, ag, ab, aa = ambientLight.Color.RGBA()
	}

	// draw each light
	lights := s.entityManager.FindByComponents(components.LightType, components.TransformType)
	for _, l := range lights {
		light := l.GetComponent(components.LightType).(*components.Light)
		translate := l.GetComponent(components.TransformType).(*components.Transform)
		img := light.Image
		rect := img.Bounds()

		x := translate.X
		y := translate.Y
		scale := translate.Scale

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(rect.Dx())/2, -float64(rect.Dy())/2)
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(float64(rect.Dx())/2, float64(rect.Dy())/2)
		op.GeoM.Translate(x, y)

		r, g, b, a := lib.GetRGBA64(light.Color)
		op.ColorM.Scale(r, g, b, a)
		s.lightingTexture.DrawImage(img, op)
	}

	// send to shader
	// see shader: https://github.com/hajimehoshi/ebiten/blob/21207f827c062deb451088f2a8c4e32b041a5793/examples/shader/lighting.go
	w, h := screen.Size()
	op := &ebiten.DrawRectShaderOptions{}
	op.CompositeMode = ebiten.CompositeModeMultiply
	op.Uniforms = map[string]interface{}{
		"AmbientColor": shaders.Vec4(ar, ag, ab, aa),
	}
	op.Images[0] = s.lightingTexture

	screen.DrawRectShader(w, h, s.ambientShader, op)

}
