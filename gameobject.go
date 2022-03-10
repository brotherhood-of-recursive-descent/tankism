package tankism

import "github.com/hajimehoshi/ebiten/v2"

type GameObject interface {
	Draw(image *ebiten.Image)
	Update() error
}
