package main

type Paddle struct {
	X, Y float32
	W, H float32
}

// Clamp the movement of paddle to screen width and height
func (p *Paddle) Clamp(screenW, screenH float32, paddingX, paddingY float32) {
	// Horizontal clamp
	if p.X < paddingX {
		p.X = paddingX
	}
	if p.X+p.W > screenW-paddingX {
		p.X = screenW - p.W - paddingX
	}

	// Vertical clamp
	if p.Y < paddingY {
		p.Y = paddingY
	}
	if p.Y+p.H > screenH-paddingY {
		p.Y = screenH - p.H - paddingY
	}
}

// Collides checks if paddle collides with ball
func (p *Paddle) Collides(ballX, ballY, ballSize float32) bool {
	return ballX+ballSize >= p.X &&
		ballX <= p.X+p.W &&
		ballY+ballSize >= p.Y &&
		ballY <= p.Y+p.H
}
