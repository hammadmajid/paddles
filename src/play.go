package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Play struct {
	top    Paddle
	right  Paddle
	bottom Paddle
	left   Paddle

	paddleSpeed float32 // movement speed of both paddles

	ball Ball
}

func (p *Play) Update() {
	// Store previous positions to calculate velocity
	topOldX := p.top.X
	rightOldY := p.right.Y
	bottomOldX := p.bottom.X
	leftOldY := p.left.Y

	// assign keys to control each paddle
	topControls := Controls{NegX: ebiten.KeyA, PosX: ebiten.KeyD}
	rightControls := Controls{NegY: ebiten.KeyArrowUp, PosY: ebiten.KeyArrowDown}
	bottomControls := Controls{NegX: ebiten.KeyArrowLeft, PosX: ebiten.KeyArrowRight}
	leftControls := Controls{NegY: ebiten.KeyW, PosY: ebiten.KeyS}

	p.top.Move(topControls)
	p.right.Move(rightControls)
	p.bottom.Move(bottomControls)
	p.left.Move(leftControls)

	const padding = 16

	// clamp paddle movement
	p.top.Clamp(screenW, screenH, padding, padding)
	p.right.Clamp(screenW, screenH, padding, padding)
	p.bottom.Clamp(screenW, screenH, padding, padding)
	p.left.Clamp(screenW, screenH, padding, padding)

	// Calculate paddle velocities
	p.top.VX = p.top.X - topOldX
	p.top.VY = 0
	p.right.VX = 0
	p.right.VY = p.right.Y - rightOldY
	p.bottom.VX = p.bottom.X - bottomOldX
	p.bottom.VY = 0
	p.left.VX = 0
	p.left.VY = p.left.Y - leftOldY

	// ball movement and bounce
	p.ball.Move()
	p.ball.CheckHorizontalPaddle(p.top, true)
	p.ball.CheckHorizontalPaddle(p.bottom, false)
	p.ball.CheckVerticalPaddle(p.left, true)
	p.ball.CheckVerticalPaddle(p.right, false)
}

func (p *Play) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.top.X, p.top.Y, p.top.W, p.top.H, color.White, true)
	vector.DrawFilledRect(screen, p.right.X, p.right.Y, p.right.W, p.right.H, color.White, true)
	vector.DrawFilledRect(screen, p.bottom.X, p.bottom.Y, p.bottom.W, p.bottom.H, color.White, true)
	vector.DrawFilledRect(screen, p.left.X, p.left.Y, p.left.W, p.left.H, color.White, true)

	vector.DrawFilledRect(screen, p.ball.X, p.ball.Y, p.ball.Size, p.ball.Size, color.RGBA{R: 200}, true)
}
