package internal

import (
	"log"
	"path"

	sdlImage "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type AssetManager struct {
	assetDir string
	sprites  map[string]*sdl.Surface
}

func NewAssetManager(assetDirectory string) (*AssetManager, error) {

	m := AssetManager{
		assetDir: assetDirectory,
		sprites:  make(map[string]*sdl.Surface),
	}

	if err := m.loadSprites(); err != nil {
		return nil, err
	}
	log.Printf("asset manager: loaded %d assets \n", len(m.sprites))

	return &m, nil
}

func (am *AssetManager) loadSprites() error {

	tankName := "tank_dark.png"
	img, err := sdlImage.Load(path.Join(am.assetDir, "sprites", tankName))
	am.sprites[tankName] = img
	return err
}

func (am *AssetManager) getSprite(name string) *sdl.Surface {
	sprite, ok := am.sprites[name]
	if !ok {
		log.Panic("failed to load sprite: " + name)
	}
	return sprite
}

func (am *AssetManager) GetTexture(r *sdl.Renderer, name string) *sdl.Texture {
	tex, err := r.CreateTextureFromSurface(am.getSprite(name))
	if err != nil {
		log.Panic("failed to get texture: " + name)
	}
	return tex
}
