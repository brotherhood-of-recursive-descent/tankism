package main

var AmbientColor vec4

// Fragment is applying the uniform AmbientColor to each pixel based on th
// supplied composite mode
func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	return vec4(AmbientColor)
}
