package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/objects/ball"
	"github.com/hammadmajid/paddle/internal/objects/paddle"
)

type Play struct {
	Top    paddle.Paddle
	Right  paddle.Paddle
	Bottom paddle.Paddle
	Left   paddle.Paddle

	Ball ball.Ball
}

func (p *Play) Update() bool {
	// assign keys to control each paddle
	//topControls := Controls{NegX: ebiten.KeyA, PosX: ebiten.KeyD}
	//rightControls := Controls{NegY: ebiten.KeyArrowUp, PosY: ebiten.KeyArrowDown}
	bottomControls := paddle.Controls{NegX: ebiten.KeyArrowLeft, PosX: ebiten.KeyArrowRight}
	//leftControls := Controls{NegY: ebiten.KeyW, PosY: ebiten.KeyS}

	//p.Top.Move(topControls)
	//p.Right.Move(rightControls)
	p.Bottom.Move(bottomControls)
	//p.Left.Move(leftControls)

	p.Top.AutoMove(p.Ball.X, p.Ball.Y, p.Ball.VX, p.Ball.VY)
	p.Left.AutoMove(p.Ball.X, p.Ball.Y, p.Ball.VX, p.Ball.VY)
	p.Right.AutoMove(p.Ball.X, p.Ball.Y, p.Ball.VX, p.Ball.VY)

	const padding = 16

	// clamp paddle movement
	p.Top.Clamp(config.ScreenW, config.ScreenH, padding, padding)
	p.Right.Clamp(config.ScreenW, config.ScreenH, padding, padding)
	p.Bottom.Clamp(config.ScreenW, config.ScreenH, padding, padding)
	p.Left.Clamp(config.ScreenW, config.ScreenH, padding, padding)

	// Ball movement and bounce
	p.Ball.Move()
	p.Ball.Bounce(p.Top)
	p.Ball.Bounce(p.Bottom)
	p.Ball.Bounce(p.Left)
	p.Ball.Bounce(p.Right)

	collides := p.Ball.CheckWalls(config.ScreenW, config.ScreenH, padding)
	if collides {
		p.Ball.Reset()
		p.Top.ResetPosition()
		p.Right.ResetPosition()
		p.Bottom.ResetPosition()
		p.Left.ResetPosition()
		// return true to indicate game over
		return true
	}
	return false
}

func (p *Play) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.Top.X, p.Top.Y, p.Top.W, p.Top.H, config.ColorLavender, true)
	vector.DrawFilledRect(screen, p.Right.X, p.Right.Y, p.Right.W, p.Right.H, config.ColorLavender, true)
	vector.DrawFilledRect(screen, p.Bottom.X, p.Bottom.Y, p.Bottom.W, p.Bottom.H, config.ColorLavender, true)
	vector.DrawFilledRect(screen, p.Left.X, p.Left.Y, p.Left.W, p.Left.H, config.ColorLavender, true)

	vector.DrawFilledRect(screen, p.Ball.X, p.Ball.Y, p.Ball.Size, p.Ball.Size, config.ColorRed, true)
}
