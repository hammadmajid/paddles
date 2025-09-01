package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/interfaces"
	"github.com/hammadmajid/paddle/internal/objects/ball"
	"github.com/hammadmajid/paddle/internal/objects/paddle"
)

const PlayerPaddleIndex = 2

type Play struct {
	Paddles []paddle.Paddle
	Ball    ball.Ball
}

func (p *Play) Update() interfaces.StateTransition {
	bottomPaddleControls := paddle.Controls{
		NegX: ebiten.KeyArrowLeft,
		PosX: ebiten.KeyArrowRight,
	}
	p.Paddles[PlayerPaddleIndex].Move(bottomPaddleControls)

	p.Ball.Move()

	for i := range p.Paddles {
		if i != PlayerPaddleIndex {
			p.Paddles[i].AutoMove(p.Ball.X, p.Ball.Y, p.Ball.VX, p.Ball.VY)
		}
		p.Paddles[i].Clamp(config.ScreenW, config.ScreenH, config.PaddingDefault, config.PaddingDefault)
		p.Ball.Bounce(p.Paddles[i])
	}

	if p.Ball.CheckWalls(config.ScreenW, config.ScreenH, config.PaddingDefault) {
		p.Ball.Reset()
		for i := range p.Paddles {
			p.Paddles[i].ResetPosition()
		}
		return interfaces.StateTransition{
			ShouldTransition: true,
			NewState:         interfaces.StateOver,
		}
	}
	return interfaces.StateTransition{}
}

func (p *Play) Draw(screen *ebiten.Image) {
	for _, v := range p.Paddles {
		vector.DrawFilledRect(screen, v.X, v.Y, v.W, v.H, config.ColorLavender, true)
	}
	vector.DrawFilledRect(screen, p.Ball.X, p.Ball.Y, p.Ball.Size, p.Ball.Size, config.ColorRed, true)
}
