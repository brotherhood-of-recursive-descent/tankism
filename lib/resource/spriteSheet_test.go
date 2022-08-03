package resource_test

import (
	_ "embed"
	"testing"

	"github.com/co0p/tankism/lib/resource"
	"github.com/co0p/tankism/test"
)

func Test_NewSpriteSheet_OneSpriteSheetOneImage(t *testing.T) {

	// given
	data := test.ValidImage64x64

	// when
	ss, err := resource.NewSpriteSheet(data, 64, 64)

	// then
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	if len(ss) != 1 {
		t.Errorf("expected spritesheet to have length 1, got %d", len(ss))
	}
}

func Test_NewSpriteSheet_OneSpriteSheet4Images(t *testing.T) {

	// given
	data := test.ValidImage64x64

	// when
	ss, err := resource.NewSpriteSheet(data, 32, 32)

	// then
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	if len(ss) != 4 {
		t.Errorf("expected spritesheet to have length 4, got %d", len(ss))
	}
}

func Test_NewSpriteSheet_EmptyFileProvided(t *testing.T) {

	// given
	data := test.EmptyFile

	// when
	ss, err := resource.NewSpriteSheet(data, 0, 0)

	// then
	if err == nil {
		t.Errorf("expected err not to be nil\n")
	}

	if len(ss) > 0 {
		t.Errorf("expected spritesheet to be empty")
	}
}

func Test_NewSpriteSheet_NotImageProvided(t *testing.T) {

	// given
	data := test.NotImageFile

	// when
	ss, err := resource.NewSpriteSheet(data, 0, 0)

	// then
	if err == nil {
		t.Errorf("expected err not to be nil\n")
	}

	if len(ss) > 0 {
		t.Errorf("expected spritesheet to be empty")
	}
}

func Test_NewSpriteSheet_OutOfBoundsWidthAndHeight(t *testing.T) {

	// given
	data := test.ValidImage64x64

	// when
	invalidWith := 1024
	validHeight := 8
	ss, err := resource.NewSpriteSheet(data, invalidWith, validHeight)

	// then
	if err == nil {
		t.Errorf("expected err not to be nil\n")
	}

	if len(ss) > 0 {
		t.Errorf("expected spritesheet to be empty")
	}

	// when
	validWidth := 8
	invalidHeight := 1024
	ss, err = resource.NewSpriteSheet(data, validWidth, invalidHeight)

	// then
	if err == nil {
		t.Errorf("expected err not to be nil\n")
	}

	if len(ss) > 0 {
		t.Errorf("expected spritesheet to be empty")
	}
}

func Test_NewSpriteSheet_InvalidWidthAndHeight(t *testing.T) {

	// given
	data := test.ValidImage64x64

	// when
	invalidWith := 13
	validHeight := 8
	ss, err := resource.NewSpriteSheet(data, invalidWith, validHeight)

	// then
	if err == nil {
		t.Errorf("expected err not to be nil\n")
	}

	if len(ss) > 0 {
		t.Errorf("expected spritesheet to be empty")
	}

	// when
	validWidth := 8
	invalidHeight := 13
	ss, err = resource.NewSpriteSheet(data, validWidth, invalidHeight)

	// then
	if err == nil {
		t.Errorf("expected err not to be nil\n")
	}

	if len(ss) > 0 {
		t.Errorf("expected spritesheet to be empty")
	}
}
