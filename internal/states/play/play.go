package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/objects/ball"
	"github.com/hammadmajid/paddle/internal/objects/paddle"
)

type Play struct {
	Paddles []paddle.Paddle
	Ball    ball.Ball
}

func (p *Play) Update() bool {
	// this depends on how the Paddles array is initalized in
	// order matters in initalization
	const bottomPaddleIndex = 2

	bottomPaddleControls := paddle.Controls{NegX: ebiten.KeyArrowLeft, PosX: ebiten.KeyArrowRight}
	p.Paddles[bottomPaddleIndex].Move(bottomPaddleControls)

	const padding = 16

	p.Ball.Move()

	for i := range p.Paddles {
		if i != bottomPaddleIndex {
			p.Paddles[i].AutoMove(p.Ball.X, p.Ball.Y, p.Ball.VX, p.Ball.VY)
		}

		p.Paddles[i].Clamp(config.ScreenW, config.ScreenH, padding, padding)

		if p.Paddles[i].Collides(p.Ball.X, p.Ball.Y, p.Ball.Size) {
			p.Ball.Bounce(p.Paddles[i])
		}

		p.Paddles[i].UpdateColor()
	}

	collides := p.Ball.CheckWalls(config.ScreenW, config.ScreenH, padding)
	if collides {
		p.Ball.Reset()
		for i := range p.Paddles {
			p.Paddles[i].ResetPosition()
		}
		// return true to indicate game over
		return true
	}
	return false
}

func (p *Play) Draw(screen *ebiten.Image) {
	for i := range p.Paddles {
		vector.DrawFilledRect(screen, p.Paddles[i].X, p.Paddles[i].Y, p.Paddles[i].W, p.Paddles[i].H, p.Paddles[i].Color, true)
	}
	vector.DrawFilledRect(screen, p.Ball.X, p.Ball.Y, p.Ball.Size, p.Ball.Size, config.ColorRed, true)
}
