package main

var AmbientColor vec4

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	// ( color of lightingmap + ambient color ) * destination color
	return (imageSrc0UnsafeAt(texCoord) + AmbientColor) * (color)
}
