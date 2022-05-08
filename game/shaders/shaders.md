Shaders
==========

Shaders are special programs that run on the gpu. 

Ebiten uses a byte slice as an input for the shader. 
The shader itself is written in KAGE, a language
similar to go. 

To have the shader source ignored by the go tooling,
add a '_' as a file name prefix.

Read more about shaders at https://ebiten.org/documents/shader.html