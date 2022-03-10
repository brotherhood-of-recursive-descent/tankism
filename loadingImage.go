package tankism

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	ALPHA_STEP = 0.005
	ALPHA_MAX  = 1
)

var colorBlack = color.Black

// LoadingImage is the background image of the start screen
type LoadingImage struct {
	scene        Scene
	image        *ebiten.Image
	currentAlpha float64
	count        int
}

func NewLoadingImage(scene Scene) *LoadingImage {

	img, _ := LoadImage(BackgroundImage)
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

	w, h := l.scene.WindowDimension()
	x := float64(w)*0.5 - (float64(imageWidth) * 0.5 * scaleX)
	y := float64(h)*0.5 - (float64(imageHeight) * 0.5 * scaleY)

	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(x, y)
	op.ColorM.ChangeHSV(1, 1, l.currentAlpha)

	screen.Fill(colorBlack)
	screen.DrawImage(l.image, op)
}

func (l *LoadingImage) Update() error {
	if l.currentAlpha < ALPHA_MAX {
		l.currentAlpha += ALPHA_STEP
	}
	return nil
}
