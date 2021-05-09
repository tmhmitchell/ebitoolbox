// Based on Ebiten's Camera example:
// https://github.com/hajimehoshi/ebiten/blob/29eade9b4a79f23637597af88cdee9b8e2e44eea/examples/camera/main.go

package camera

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmhmitchell/ebitoolbox/datastructures/vector"
)

// Drawable is an interface that allows a struct to be drawn to the user's screen.
type Drawable interface {
	// To draw an entity, we need to know it's position in the world.
	X() float64
	Y() float64

	// We also need to know what the entity looks like.
	Sprite() *ebiten.Image

	// Finally, an entity can provide us a GeoM to adjust how it's sprite is drawn.
	// Typically this is used to flip the sprite in the vertical axis around some point.
	// If this is not required, it's acceptable to retun an identity matrix.
	FlipGeoM() ebiten.GeoM
}

// Camera ...
type Camera struct {
	// viewport represents the width and height of the what the camera displays
	viewport vector.Vec2

	// target represents what the camera is centered on, in world space
	target vector.Vec2

	// zoom represents the amount drawn sprites are scaled by
	zoom float64

	// baseSpriteSize is the minimum length of the X and Y axis of your game's sprites
	// If your world space is measured in pixels, set this to 1!
	baseSpriteSize float64
}

// NewCamera creates a new camera instance.
func NewCamera(w, h, baseSpriteSize float64) *Camera {
	return &Camera{
		viewport:       vector.NewVec2(w, h),
		target:         vector.NewVec2(0, 0),
		zoom:           1,
		baseSpriteSize: baseSpriteSize,
	}
}

// Target returns the camera's target attribute
func (c Camera) Target() vector.Vec2 { return c.target }

// SetTarget sets the x and y components of the camera's target attribute
func (c *Camera) SetTarget(x, y float64) {
	c.target.SetX(x)
	c.target.SetY(y)
}

// Zoom returns the camera's zoom attribute
func (c Camera) Zoom() float64 { return c.zoom }

// SetZoom sets the camera's zoom attribute
func (c *Camera) SetZoom(z float64) { c.zoom = z }

// WorldToScreenGeoM returns a transformation matrix for changing world-space
// coordinates to screen-space coordinates.
func (c *Camera) worldToScreenGeoM(wx, wy float64) ebiten.GeoM {
	g := ebiten.GeoM{}

	g.Translate(
		((wx - c.target.X()) * c.baseSpriteSize),
		((wy - c.target.Y()) * c.baseSpriteSize),
	)

	g.Scale(c.zoom, c.zoom)

	g.Translate(c.viewport.X()*0.5, c.viewport.Y()*0.5)

	return g
}

// Draw draws a Drawable to a destination image.
func (c *Camera) Draw(dst *ebiten.Image, d Drawable) {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Concat(d.FlipGeoM())
	ops.GeoM.Concat(c.worldToScreenGeoM(d.X(), d.Y()))
	dst.DrawImage(d.Sprite(), ops)
}
