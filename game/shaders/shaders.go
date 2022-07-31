package shaders

import _ "embed"

//go:embed _lighting.go
var LightingShader []byte

func Vec4(r, g, b, a uint32) []float32 {
	return []float32{float32(float32(r) / 65535.0), float32(float32(g) / 65535.0), float32(float32(b) / 65535.0), float32(float32(a) / 65535.0)}
}
