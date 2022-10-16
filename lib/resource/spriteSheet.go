package resource

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"image"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"

	_ "image/jpeg"
	_ "image/png"
)

type SpriteSheet []SpriteSheetEntry

type SpriteSheetEntry struct {
	ImageRef *ebiten.Image
	Name     string
}

func NewSpriteSheet(data []byte, strideX, strideY int) (SpriteSheet, error) {
	res := SpriteSheet{}

	// bytes is empty
	if len(data) == 0 {
		return res, errors.New("data must not be empty")
	}

	// not image
	ssImg, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return res, errors.New("failed to load image: " + err.Error())
	}

	bounds := ssImg.Bounds()
	ssW, ssH := bounds.Size().X, bounds.Size().Y
	// image width is less than width
	if ssW < strideX {
		return res, fmt.Errorf("spritesheet width: %d is less than width: %d", ssW, strideX)
	}

	// image width is not divisible by width
	if ssW%strideX != 0 {
		return res, fmt.Errorf("spritesheet width: %d is not divisible by width: %d", ssW, strideX)
	}

	// image height is less than height
	if ssH < strideY {
		return res, fmt.Errorf("spritesheet height: %d is less than height: %d", ssH, strideY)
	}

	// image height is not divisible by height
	if ssH%strideY != 0 {
		return res, fmt.Errorf("spritesheet height: %d is not divisible by height: %d", ssH, strideY)
	}

	// load (from left to right, top to bottom) and return
	texture := ebiten.NewImageFromImage(ssImg)
	idx := 0
	for y := 0; y < ssH; y += strideY {
		for x := 0; x < ssW; x += strideX {
			dim := image.Rectangle{image.Point{x, y}, image.Point{x + strideX, y + strideY}}
			res = append(res, SpriteSheetEntry{
				ImageRef: texture.SubImage(dim).(*ebiten.Image),
				Name:     strconv.Itoa(idx),
			})
			idx++
		}
	}

	return res, nil
}

type SpriteSheetConfig struct {
	XMLName xml.Name                 `xml:"TextureAtlas"`
	Entries []SpriteSheetConfigEntry `xml:"SubTexture"`
}

type SpriteSheetConfigEntry struct {
	PosX   int    `xml:"x,attr"`
	PosY   int    `xml:"y,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Name   string `xml:"name,attr"`
}

func NewSpriteSheetFromConfig(data []byte, configData []byte) (SpriteSheet, error) {
	res := SpriteSheet{}

	if len(data) == 0 {
		return res, errors.New("texture must not be empty")
	}
	if len(configData) == 0 {
		return res, errors.New("configData must not be empty")
	}

	// not image
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return res, errors.New("failed to load texture: " + err.Error())
	}

	var config SpriteSheetConfig

	// not parsable xml
	err = xml.Unmarshal(configData, &config)
	if err != nil {
		return res, errors.New("failed to load xml: " + err.Error())
	}

	texture := ebiten.NewImageFromImage(img)
	for _, v := range config.Entries {
		dim := image.Rectangle{
			image.Point{v.PosX, v.PosY},
			image.Point{v.PosX + v.Width, v.PosY + v.Height},
		}
		res = append(res, SpriteSheetEntry{
			ImageRef: texture.SubImage(dim).(*ebiten.Image),
			Name:     v.Name,
		})
	}

	return res, nil
}

func (s SpriteSheet) ByName(name string) *ebiten.Image {
	for _, v := range s {
		if v.Name == name {
			return v.ImageRef
		}
	}
	panic("could not find subimage with name:" + name)
}
