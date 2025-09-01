package ball

import (
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/objects/paddle"
)

var defaultBall = Ball{
	X:    config.ScreenW/2 - config.BallSize/2,
	Y:    config.ScreenH/2 - config.BallSize/2,
	VX:   config.BallSpeed,
	VY:   config.BallSpeedY,
	Size: config.BallSize,
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

func (b *Ball) GetPosition() (float32, float32) {
	return b.X, b.Y
}

func (b *Ball) SetPosition(x, y float32) {
	b.X = x
	b.Y = y
}

func (b *Ball) GetBounds() (float32, float32, float32, float32) {
	return b.X, b.Y, b.Size, b.Size
}

func (b *Ball) Bounce(p paddle.Paddle) {
	if p.Collides(b.X, b.Y, b.Size) {
		switch p.Pos {
		case paddle.Top:
			b.Y = p.Y + p.H
			b.VY = -b.VY
		case paddle.Right:
			b.X = p.X - b.Size
			b.VX = -b.VX
		case paddle.Bottom:
			b.Y = p.Y - b.Size
			b.VY = -b.VY
		case paddle.Left:
			b.X = p.X + p.W
			b.VX = -b.VX
		}
	}
}

func (b *Ball) CheckWalls(screenW, screenH, padding float32) bool {
	if b.X < padding || b.X+b.Size > screenW-padding ||
		b.Y < padding || b.Y+b.Size > screenH-padding {
		return true
	}
	return false
}
