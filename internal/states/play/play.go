package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/objects/ball"
	"github.com/hammadmajid/paddle/internal/objects/paddle"
	"image/color"
)

type Play struct {
	Top    paddle.Paddle
	Right  paddle.Paddle
	Bottom paddle.Paddle
	Left   paddle.Paddle

	Speed float32 // movement speed of both paddles

	Ball ball.Ball
}

func (p *Play) Update() bool {
	// Store previous positions to calculate velocity
	topOldX := p.Top.X
	rightOldY := p.Right.Y
	bottomOldX := p.Bottom.X
	leftOldY := p.Left.Y

	// assign keys to control each paddle
	//topControls := Controls{NegX: ebiten.KeyA, PosX: ebiten.KeyD}
	//rightControls := Controls{NegY: ebiten.KeyArrowUp, PosY: ebiten.KeyArrowDown}
	bottomControls := paddle.Controls{NegX: ebiten.KeyArrowLeft, PosX: ebiten.KeyArrowRight}
	//leftControls := Controls{NegY: ebiten.KeyW, PosY: ebiten.KeyS}

	//p.Top.Move(topControls)
	//p.Right.Move(rightControls)
	p.Bottom.Move(bottomControls)
	//p.Left.Move(leftControls)

	p.Top.AutoMove(p.Ball.X, p.Ball.Y)
	p.Left.AutoMove(p.Ball.X, p.Ball.Y)
	p.Right.AutoMove(p.Ball.X, p.Ball.Y)

	const padding = 16

	// clamp paddle movement
	p.Top.Clamp(config.ScreenW, config.ScreenH, padding, padding)
	p.Right.Clamp(config.ScreenW, config.ScreenH, padding, padding)
	p.Bottom.Clamp(config.ScreenW, config.ScreenH, padding, padding)
	p.Left.Clamp(config.ScreenW, config.ScreenH, padding, padding)

	// Calculate paddle velocities
	p.Top.VX = p.Top.X - topOldX
	p.Top.VY = 0
	p.Right.VX = 0
	p.Right.VY = p.Right.Y - rightOldY
	p.Bottom.VX = p.Bottom.X - bottomOldX
	p.Bottom.VY = 0
	p.Left.VX = 0
	p.Left.VY = p.Left.Y - leftOldY

	// Ball movement and bounce
	p.Ball.Move()
	p.Ball.CheckHorizontalPaddle(p.Top, true)
	p.Ball.CheckHorizontalPaddle(p.Bottom, false)
	p.Ball.CheckVerticalPaddle(p.Left, true)
	p.Ball.CheckVerticalPaddle(p.Right, false)

	return p.Ball.CheckWalls(config.ScreenW, config.ScreenH, padding)
}

func (p *Play) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.Top.X, p.Top.Y, p.Top.W, p.Top.H, color.White, true)
	vector.DrawFilledRect(screen, p.Right.X, p.Right.Y, p.Right.W, p.Right.H, color.White, true)
	vector.DrawFilledRect(screen, p.Bottom.X, p.Bottom.Y, p.Bottom.W, p.Bottom.H, color.White, true)
	vector.DrawFilledRect(screen, p.Left.X, p.Left.Y, p.Left.W, p.Left.H, color.White, true)

	vector.DrawFilledRect(screen, p.Ball.X, p.Ball.Y, p.Ball.Size, p.Ball.Size, color.RGBA{R: 200}, true)
}
