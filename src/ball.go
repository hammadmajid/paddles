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

// CheckHorizontalPaddle Horizontal paddle collision (top or bottom)
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

		// Add influence based on paddle velocity
		velocityInfluence := p.VX * 0.3 // scale factor for influence
		b.VY = -b.VY
		b.VX += velocityInfluence // impart horizontal velocity based on paddle movement

		// Clamp ball velocity to prevent excessive speeds
		if b.VX > 3 {
			b.VX = 3
		} else if b.VX < -4.0 {
			b.VX = -4.0
		}
	}
}

// CheckVerticalPaddle Vertical paddle collision (left or right)
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

		// Add influence based on paddle velocity
		velocityInfluence := p.VY * 0.3 // scale factor for influence
		b.VX = -b.VX
		b.VY += velocityInfluence // impart vertical velocity based on paddle movement

		// Clamp ball velocity to prevent excessive speeds
		if b.VY > 3 {
			b.VY = 3
		} else if b.VY < -4.0 {
			b.VY = -4.0
		}
	}
}

// CheckWalls return true and resets the ball position if it touches any wall
func (b *Ball) CheckWalls(screenW, screenH, padding float32) bool {
	wall := false
	// Left wall
	if b.X < padding {
		wall = true
	}
	// Right wall
	if b.X+b.Size > screenW-padding {
		wall = true
	}
	// Top wall
	if b.Y < padding {
		wall = true
	}
	// Bottom wall
	if b.Y+b.Size > screenH-padding {
		wall = true
	}

	// Reset position if wall is reached for next game
	if wall {
		b.X = screenW / 2
		b.Y = screenH / 2
	}

	return wall
}
