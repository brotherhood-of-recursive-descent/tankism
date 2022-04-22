package lib

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Init(*SceneManager) error
	Draw(image *ebiten.Image)
	Update() error
}
