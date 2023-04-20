package lib

import "github.com/hajimehoshi/ebiten/v2"

func WidthHeight(img *ebiten.Image) (int, int) {
	return img.Bounds().Dx(), img.Bounds().Dy()
}
