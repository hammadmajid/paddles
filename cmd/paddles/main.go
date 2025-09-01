package main

import (
	"log"

	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/interfaces"
	"github.com/hammadmajid/paddle/internal/objects/ball"
	"github.com/hammadmajid/paddle/internal/objects/paddle"
	"github.com/hammadmajid/paddle/internal/states/menu"
	"github.com/hammadmajid/paddle/internal/states/over"
	"github.com/hammadmajid/paddle/internal/states/play"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	currentState interfaces.StateType
	states       map[interfaces.StateType]interfaces.GameState
}

func NewGame() *Game {
	return &Game{
		currentState: interfaces.StateMenu,
		states: map[interfaces.StateType]interfaces.GameState{
			interfaces.StateMenu: &menu.Menu{
				Options: []string{"Play", "Quit"},
				Index:   0,
			},
			interfaces.StatePlay: &play.Play{
				Paddles: []paddle.Paddle{
					paddle.NewPaddle(paddle.Top),
					paddle.NewPaddle(paddle.Right),
					paddle.NewPaddle(paddle.Bottom),
					paddle.NewPaddle(paddle.Left),
				},
				Ball: ball.NewBall(),
			},
			interfaces.StateOver: &over.Over{},
		},
	}
}

func (g *Game) Update() error {
	state := g.states[g.currentState]
	transition := state.Update()

	if transition.ShouldExit {
		return ebiten.Termination
	}

	if transition.ShouldTransition {
		g.currentState = transition.NewState
		if g.currentState == interfaces.StatePlay {
			g.states[interfaces.StatePlay] = &play.Play{
				Paddles: []paddle.Paddle{
					paddle.NewPaddle(paddle.Top),
					paddle.NewPaddle(paddle.Right),
					paddle.NewPaddle(paddle.Bottom),
					paddle.NewPaddle(paddle.Left),
				},
				Ball: ball.NewBall(),
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(config.ColorBase)
	g.states[g.currentState].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(config.ScreenW), int(config.ScreenH)
}

func main() {
	ebiten.SetWindowSize(int(config.ScreenW), int(config.ScreenH))
	ebiten.SetWindowTitle("Paddles")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
