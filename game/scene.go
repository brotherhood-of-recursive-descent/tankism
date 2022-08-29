package game

import (
	"github.com/co0p/tankism/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type EmptyScene struct {
}

func (s EmptyScene) Init(sm *lib.SceneManager) error {
	return nil
}

func (s EmptyScene) Update() error {
	return nil
}

func (s EmptyScene) Draw(screen *ebiten.Image) {
}
