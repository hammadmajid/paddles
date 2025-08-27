package paddle

import "github.com/hajimehoshi/ebiten/v2"

type Position int

const (
	Top Position = iota
	Right
	Bottom
	Left
)

type Paddle struct {
	X, Y   float32
	W, H   float32
	VX, VY float32 // velocity for collision influence
	Pos    Position
}

type Controls struct {
	NegX, PosX ebiten.Key // horizontal keys
	NegY, PosY ebiten.Key // vertical keys
}

func (p *Paddle) Move(ctrl Controls) {
	if ebiten.IsKeyPressed(ctrl.NegX) {
		p.X -= 3
	}
	if ebiten.IsKeyPressed(ctrl.PosX) {
		p.X += 3
	}
	if ebiten.IsKeyPressed(ctrl.NegY) {
		p.Y -= 3
	}
	if ebiten.IsKeyPressed(ctrl.PosY) {
		p.Y += 3
	}
}

func (p *Paddle) AutoMove(ballX, ballY float32) {
	switch p.Pos {
	case Top | Bottom:
		if ballX < p.X {
			p.X -= 3
		} else if ballX > p.X {
			p.X += 3
		}
	case Left | Right:
		if ballY < p.Y {
			p.Y -= 3
		} else if ballY > p.Y {
			p.Y += 3
		}
	}
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
