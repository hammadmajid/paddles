package paddle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hammadmajid/paddle/internal/config"
)

type Position int

const (
	Top Position = iota
	Right
	Bottom
	Left
)

type Paddle struct {
	X, Y float32
	W, H float32
	Pos  Position

	speed float32
}

type Controls struct {
	NegX, PosX ebiten.Key // horizontal keys
	NegY, PosY ebiten.Key // vertical keys
}

func NewPaddle(pos Position) Paddle {
	var w, h float32
	var x, y float32

	switch pos {
	case Top:
		w = 64
		h = 12
		x = (float32(config.ScreenW) - w) / 2
		y = 10 // top padding
	case Bottom:
		w = 64
		h = 12
		x = (float32(config.ScreenW) - w) / 2
		y = float32(config.ScreenH) - h - 10 // bottom padding
	case Left:
		w = 12
		h = 64
		x = 10 // left padding
		y = (float32(config.ScreenH) - h) / 2
	case Right:
		w = 12
		h = 64
		x = float32(config.ScreenW) - w - 10 // right padding
		y = (float32(config.ScreenH) - h) / 2
	}

	return Paddle{
		X:     x,
		Y:     y,
		W:     w,
		H:     h,
		Pos:   pos,
		speed: 3.0,
	}
}

func (p *Paddle) Move(ctrl Controls) {
	if ebiten.IsKeyPressed(ctrl.NegX) {
		p.X -= p.speed
	}
	if ebiten.IsKeyPressed(ctrl.PosX) {
		p.X += p.speed
	}
	if ebiten.IsKeyPressed(ctrl.NegY) {
		p.Y -= p.speed
	}
	if ebiten.IsKeyPressed(ctrl.PosY) {
		p.Y += p.speed
	}
}

func (p *Paddle) AutoMove(ballX, ballY, ballVX, ballVY float32) {
	switch p.Pos {
	case Top, Bottom:
		if ballX < p.X {
			p.X -= ballVX
		} else if ballX > p.X {
			p.X += ballVX
		}
	case Left, Right:
		if ballY < p.Y {
			p.Y -= ballVY
		} else if ballY > p.Y {
			p.Y += ballVY
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

// ResetPosition resets the paddle position to its original position
func (p *Paddle) ResetPosition() {
	switch p.Pos {
	case Top:
		p.X = (float32(config.ScreenW) - p.W) / 2
		p.Y = 10 // top padding
	case Bottom:
		p.X = (float32(config.ScreenW) - p.W) / 2
		p.Y = float32(config.ScreenH) - p.H - 10 // bottom padding
	case Left:
		p.X = 10 // left padding
		p.Y = (float32(config.ScreenH) - p.H) / 2
	case Right:
		p.X = float32(config.ScreenW) - p.W - 10 // right padding
		p.Y = (float32(config.ScreenH) - p.H) / 2
	}
}
