package objects

import (
	_ "embed"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
)

// MenuImage is the background image of the menu screen
type MenuImage struct {
	scene lib.Scene
	image *ebiten.Image
}

func NewMenuImage(scene lib.Scene) *MenuImage {

	img, _ := media.LoadImage(media.BackgroundImage)
	sprite := ebiten.NewImageFromImage(img)

	return &MenuImage{
		scene: scene,
		image: sprite,
	}
}

func (l *MenuImage) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	scaleX, scaleY := 0.75, 0.75
	imageWidth, imageHeight := l.image.Size()

	w, h := l.scene.WindowDimension()
	x := float64(w)*0.5 - (float64(imageWidth) * 0.5 * scaleX)
	y := float64(h)*0.5 - (float64(imageHeight) * 0.5 * scaleY)

	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(x, y)

	screen.Fill(lib.ColorBlack)
	screen.DrawImage(l.image, op)
}

func (l *MenuImage) Update() error {
	return nil
}
