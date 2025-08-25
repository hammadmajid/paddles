package main

type Paddle struct {
	X, Y float32
	W, H float32
}

func NewPaddle(x, y float32) Paddle {
	return Paddle{
		X: x - paddleW/2, // center horizontally on given x
		Y: y,
		W: paddleW,
		H: paddleH,
	}
}

func (p *Paddle) Clamp(screenW float32, padding float32) {
	if p.X < padding {
		p.X = padding
	}
	if p.X+p.W > screenW-padding {
		p.X = screenW - p.W - padding
	}
}

func (p *Paddle) Collides(ballX, ballY, ballSize float32) bool {
	return ballX+ballSize >= p.X &&
		ballX <= p.X+p.W &&
		ballY+ballSize >= p.Y &&
		ballY <= p.Y+p.H
}
