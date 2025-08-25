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

// Horizontal paddle collision (top or bottom)
//
//	fromTop = true  → top paddle (ball placed just below paddle)
//	fromTop = false → bottom paddle (ball placed just above paddle)
func (b *Ball) CheckHorizontalPaddle(p Paddle, fromTop bool) {
	if p.Collides(b.X, b.Y, b.Size) {
		if fromTop {
			b.Y = p.Y + p.H
		} else {
			b.Y = p.Y - b.Size
		}
		b.VY = -b.VY
	}
}

// Vertical paddle collision (left or right)
//
//	fromLeft = true  → left paddle (ball placed just right of paddle)
//	fromLeft = false → right paddle (ball placed just left of paddle)
func (b *Ball) CheckVerticalPaddle(p Paddle, fromLeft bool) {
	if p.Collides(b.X, b.Y, b.Size) {
		if fromLeft {
			b.X = p.X + p.W
		} else {
			b.X = p.X - b.Size
		}
		b.VX = -b.VX
	}
}
