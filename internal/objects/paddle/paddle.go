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
	NegX, PosX ebiten.Key
	NegY, PosY ebiten.Key
}

func NewPaddle(pos Position) Paddle {
	paddle := Paddle{
		Pos:   pos,
		speed: config.PaddleSpeed,
	}
	paddle.setDimensionsAndPosition()
	return paddle
}

func (p *Paddle) setDimensionsAndPosition() {
	switch p.Pos {
	case Top:
		p.W = config.PaddleLength
		p.H = config.PaddleThickness
		p.X = (config.ScreenW - p.W) / 2
		p.Y = config.PaddingEdge
	case Bottom:
		p.W = config.PaddleLength
		p.H = config.PaddleThickness
		p.X = (config.ScreenW - p.W) / 2
		p.Y = config.ScreenH - p.H - config.PaddingEdge
	case Left:
		p.W = config.PaddleThickness
		p.H = config.PaddleLength
		p.X = config.PaddingEdge
		p.Y = (config.ScreenH - p.H) / 2
	case Right:
		p.W = config.PaddleThickness
		p.H = config.PaddleLength
		p.X = config.ScreenW - p.W - config.PaddingEdge
		p.Y = (config.ScreenH - p.H) / 2
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

func (p *Paddle) Clamp(screenW, screenH float32, paddingX, paddingY float32) {
	if p.X < paddingX {
		p.X = paddingX
	}
	if p.X+p.W > screenW-paddingX {
		p.X = screenW - p.W - paddingX
	}

	if p.Y < paddingY {
		p.Y = paddingY
	}
	if p.Y+p.H > screenH-paddingY {
		p.Y = screenH - p.H - paddingY
	}
}

func (p *Paddle) Collides(ballX, ballY, ballSize float32) bool {
	return ballX+ballSize >= p.X &&
		ballX <= p.X+p.W &&
		ballY+ballSize >= p.Y &&
		ballY <= p.Y+p.H
}

func (p *Paddle) ResetPosition() {
	p.setDimensionsAndPosition()
}

func (p *Paddle) GetPosition() (float32, float32) {
	return p.X, p.Y
}

func (p *Paddle) SetPosition(x, y float32) {
	p.X = x
	p.Y = y
}

func (p *Paddle) GetBounds() (float32, float32, float32, float32) {
	return p.X, p.Y, p.W, p.H
}
