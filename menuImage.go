package tankism

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
)

// MenuImage is the background image of the menu screen
type MenuImage struct {
	scene Scene
	image *ebiten.Image
}

func NewMenuImage(scene Scene) *MenuImage {

	img, _ := LoadImage(BackgroundImage)
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

	screen.Fill(colorBlack)
	screen.DrawImage(l.image, op)
}

func (l *MenuImage) Update() error {
	return nil
}
