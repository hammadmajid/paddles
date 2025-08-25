package main

type Ball struct {
	X, Y   float32
	VX, VY float32
	Size   float32
}

func (b *Ball) Move() {
	b.X += b.VX
	b.Y += b.VY
}

func (b *Ball) BounceX() { b.VX = -b.VX }
func (b *Ball) BounceY() { b.VY = -b.VY }

func (b *Ball) CheckWalls(screenW, screenH, padding float32) {
	if b.X < padding {
		b.X = padding
		b.BounceX()
	}
	if b.X+b.Size > screenW-padding {
		b.X = screenW - padding - b.Size
		b.BounceX()
	}
}

func (b *Ball) CheckPaddle(p Paddle, fromTop bool) {
	if p.Collides(b.X, b.Y, b.Size) {
		if fromTop {
			b.Y = p.Y + p.H
		} else {
			b.Y = p.Y - b.Size
		}
		b.BounceY()
	}
}
