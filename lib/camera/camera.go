package camera

import (
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type CameraMode int

const (
	CameraModeDefault = iota
	CameraModeCenter
)

// Camera implements a basic camera following the idea from https://github.com/MelonFunction/ebiten-camera
type Camera struct {
	position vector.Vec2d
	scale    float64
	mode     CameraMode

	width, height int
	Surface       *ebiten.Image
}

func NewCamera(w, h int) *Camera {
	return &Camera{
		width:   w,
		height:  h,
		mode:    CameraModeDefault,
		scale:   1,
		Surface: ebiten.NewImage(w, h),
	}
}

func (c *Camera) SetMode(mode CameraMode) {
	c.mode = mode
}

func (c *Camera) Move(dest vector.Vec2d) {
	c.position = dest
}

// SetZoom sets the zoom
func (c *Camera) SetZoom(zoom float64) *Camera {
	c.scale = zoom
	if c.scale <= 0.01 {
		c.scale = 0.01
	}
	c.resize(c.width, c.height)
	return c
}

// resize resizes the camera Surface
func (c *Camera) resize(w, h int) *Camera {
	c.width = w
	c.height = h
	newW := int(float64(w) * 1.0 / c.scale)
	newH := int(float64(h) * 1.0 / c.scale)
	if newW <= 16384 && newH <= 16384 {
		c.Surface.Dispose()
		c.Surface = ebiten.NewImage(newW, newH)
	}
	return c
}

func (c *Camera) Draw(target *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	w, h := lib.WidthHeight(c.Surface)
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0

	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Scale(c.scale, c.scale)
	// op.GeoM.Rotate(c.Rot)
	op.GeoM.Translate(cx*c.scale, cy*c.scale)

	op.GeoM.Translate(c.position.X, c.position.Y)

	target.DrawImage(c.Surface, op)

}
