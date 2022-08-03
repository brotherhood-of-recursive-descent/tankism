package test

import _ "embed"

//go:embed emptyFile.png
var EmptyFile []byte

//go:embed notImage.txt
var NotImageFile []byte

//go:embed 64x64.png
var ValidImage64x64 []byte
