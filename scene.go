package tankism

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Init() error
	Draw(image *ebiten.Image)
	Update() error
	WindowDimension() (int, int)
	SetWindowDimension(int, int)
}
