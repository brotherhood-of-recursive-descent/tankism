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
	image        *ebiten.Image
	currentAlpha float64
	count        int
}

func NewLoadingImage() *LoadingImage {

	img, _ := resources.LoadImage(resources.BackgroundImage)
	sprite := ebiten.NewImageFromImage(img)

	return &LoadingImage{
		image:        sprite,
		currentAlpha: 0,
		count:        0,
	}
}

func (l *LoadingImage) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	scaleX, scaleY := 0.75, 0.75
	imageWidth, imageHeight := lib.WidthHeight(l.image)

	w, h := lib.WidthHeight(screen)
	x := float64(w)*0.5 - (float64(imageWidth) * 0.5 * scaleX)
	y := float64(h)*0.5 - (float64(imageHeight) * 0.5 * scaleY)

	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleAlpha(float32(l.currentAlpha))

	screen.Fill(lib.ColorBlack)
	screen.DrawImage(l.image, op)
}

func (l *LoadingImage) Update() error {
	if l.currentAlpha < ALPHA_MAX {
		l.currentAlpha += ALPHA_STEP
	}
	return nil
}
