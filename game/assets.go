package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// AssetManager holds the assets used by components in the game.
//
// TODO:
// - add loadImagesFromMap() from spritemap file
// - add loadFont() from file
// - add loadSound() from file

type AssetManager struct {
	images map[string]*ebiten.Image
	// sounds map[string][]byte
	// fonts  map[string]font.Face
}

func NewAssetManager() *AssetManager {

	return &AssetManager{
		images: make(map[string]*ebiten.Image),
		// sounds: make(map[string][]byte),
		// fonts:  make(map[string]font.Face),
	}
}

func (a *AssetManager) LoadImage(key string, path string) error {
	// TODO - continue here
	return nil
}

func (a *AssetManager) AddImage(key string, image *ebiten.Image) {
	a.images[key] = image
}

func (a *AssetManager) GetImage(key string) (*ebiten.Image, bool) {
	v, ok := a.images[key]
	return v, ok
}

/*
func (a *AssetManager) AddSound(key string, sound []byte) {
	a.sounds[key] = sound
}

func (a *AssetManager) GetSound(key string) ([]byte, bool) {
	v, ok := a.sounds[key]
	return v, ok
}

func (a *AssetManager) AddFont(key string, font font.Face) {
	a.fonts[key] = font
}

func (a *AssetManager) GetFont(key string) (font.Face, bool) {
	v, ok := a.fonts[key]
	return v, ok
}

*/
