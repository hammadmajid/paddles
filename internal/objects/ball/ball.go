package ball

import (
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/objects/paddle"
)

var defaultBall = Ball{
	X:    config.ScreenW/2 - 4,
	Y:    config.ScreenH/2 - 4,
	VX:   1.5,
	VY:   2.5,
	Size: 8,
}

type Ball struct {
	X, Y   float32
	VX, VY float32
	Size   float32
}

func NewBall() Ball {
	return defaultBall
}

// Reset sets the ball to its initial position and velocity
func (b *Ball) Reset() {
	b.X = defaultBall.X
	b.Y = defaultBall.Y
	b.VX = defaultBall.VX
	b.VY = defaultBall.VY
}

func (b *Ball) Move() {
	b.X += b.VX
	b.Y += b.VY
}

func (b *Ball) Bounce(p paddle.Paddle) {
	switch p.Pos {
	case paddle.Top:
		// bounce
		b.Y = p.Y + p.H
		// invert velocity
		b.VY = -b.VY
	case paddle.Right:
		// bounce
		b.X = p.X - b.Size
		// invert velocity
		b.VX = -b.VX
	case paddle.Bottom:
		// bounce
		b.Y = p.Y - b.Size
		// invert velocity
		b.VY = -b.VY
	case paddle.Left:
		// bounce
		b.X = p.X + p.W
		// invert velocity
		b.VX = -b.VX
	}
}

// CheckWalls return true and resets the ball position if it touches any wall
func (b *Ball) CheckWalls(screenW, screenH, padding float32) bool {
	// Left wall
	if b.X < padding {
		return true
	}
	// Right wall
	if b.X+b.Size > screenW-padding {
		return true
	}
	// Top wall
	if b.Y < padding {
		return true
	}
	// Bottom wall
	if b.Y+b.Size > screenH-padding {
		return true
	}

	return false
}

// abs returns the absolute value of a float32
func abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
