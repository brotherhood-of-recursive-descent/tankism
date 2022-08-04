package test

import _ "embed"

//go:embed emptyFile.png
var EmptyFile []byte

//go:embed notImage.txt
var NotImageFile []byte

//go:embed 64x64.png
var ValidImage64x64 []byte

//go:embed texture64x64.xml
var Texture64x64 []byte

//go:embed texture4x32.xml
var Texture4x32 []byte
