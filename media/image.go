package media

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
)

func LoadImage(input []byte) (image.Image, error) {
	if img, _, err := image.Decode(bytes.NewReader(input)); err != nil {
		return img, err
	} else {
		return img, nil
	}
}
