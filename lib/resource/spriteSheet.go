package resource

import (
	"bytes"
	"errors"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	_ "image/jpeg"
	_ "image/png"
)

// helper https://stackoverflow.com/questions/16072910/trouble-getting-a-subimage-of-an-image-in-go
type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

type SpriteSheet struct{}

type SpriteSheetConfig struct {
	Width, Height int
	Entries       []SpriteSheetEntry
}

type SpriteSheetEntry struct {
	PosX, PosY, Width, Height int
}

func NewSpriteSheet(data []byte, strideX, strideY int) ([]*ebiten.Image, error) {
	res := []*ebiten.Image{}

	// bytes is empty
	if len(data) == 0 {
		return res, errors.New("data must not be empty")
	}

	// not image
	ssImg, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return res, errors.New("failed to load spritesheet: " + err.Error())
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
	for y := 0; y < ssH; y += strideY {
		for x := 0; x < ssW; x += strideX {

			subImg := ssImg.(SubImager).SubImage(image.Rect(x, y, x+strideX, y+strideY))
			fmt.Printf("subImg: %v", subImg.Bounds())

			res = append(res, ebiten.NewImageFromImage(subImg))
		}
	}

	return res, nil
}
