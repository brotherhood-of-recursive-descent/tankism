package objects

import (
	_ "embed"

	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ALPHA_STEP = 0.005
	ALPHA_MAX  = 1
)

// LoadingImage is the background image of the start screen
type LoadingImage struct {
	scene        lib.Scene
	image        *ebiten.Image
	currentAlpha float64
	count        int
}

func NewLoadingImage(scene lib.Scene) *LoadingImage {

	img, _ := resources.LoadImage(resources.BackgroundImage)
	sprite := ebiten.NewImageFromImage(img)

	return &LoadingImage{
		scene:        scene,
		image:        sprite,
		currentAlpha: 0,
		count:        0,
	}
}

func (l *LoadingImage) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	scaleX, scaleY := 0.75, 0.75
	imageWidth, imageHeight := l.image.Size()

	w, h := screen.Size()
	x := float64(w)*0.5 - (float64(imageWidth) * 0.5 * scaleX)
	y := float64(h)*0.5 - (float64(imageHeight) * 0.5 * scaleY)

	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(x, y)
	op.ColorM.ChangeHSV(1, 1, l.currentAlpha)

	screen.Fill(lib.ColorBlack)
	screen.DrawImage(l.image, op)
}

func (l *LoadingImage) Update() error {
	if l.currentAlpha < ALPHA_MAX {
		l.currentAlpha += ALPHA_STEP
	}
	return nil
}
