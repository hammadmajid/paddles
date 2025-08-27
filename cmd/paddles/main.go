package main

import (
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/objects/ball"
	"github.com/hammadmajid/paddle/internal/objects/paddle"
	"github.com/hammadmajid/paddle/internal/states/menu"
	"github.com/hammadmajid/paddle/internal/states/over"
	"github.com/hammadmajid/paddle/internal/states/play"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlay
	StateOver
)

type Game struct {
	state GameState

	menu *menu.Menu
	play *play.Play
	over *over.Over
}

func (g *Game) Update() error {
	switch g.state {
	case StateMenu:
		g.menu.Update()
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			if g.menu.Options[g.menu.Index] == "Play" {
				g.state = StatePlay
			}
			if g.menu.Options[g.menu.Index] == "Quit" {
				return ebiten.Termination
			}
		}
	case StatePlay:
		stateChanged := g.play.Update()
		if stateChanged {
			g.state = StateOver
		}
	case StateOver:
		stateChanged := g.over.Update()
		if stateChanged {
			g.state = StateMenu
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(config.ColorBase)

	switch g.state {
	case StateMenu:
		g.menu.Draw(screen)
	case StatePlay:
		g.play.Draw(screen)
	case StateOver:
		g.over.Draw(screen)
	}
}

//goland:noinspection GoUnusedParameter
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenW, config.ScreenH
}

func main() {
	ebiten.SetWindowSize(config.ScreenW, config.ScreenH)
	ebiten.SetWindowTitle("Paddles")
	if err := ebiten.RunGame(&Game{
		menu: &menu.Menu{
			Options: []string{"Play", "Quit"},
			Index:   0,
		},
		play: &play.Play{
			Top:    paddle.NewPaddle(paddle.Top),
			Right:  paddle.NewPaddle(paddle.Right),
			Bottom: paddle.NewPaddle(paddle.Bottom),
			Left:   paddle.NewPaddle(paddle.Left),

			Ball: ball.Ball{
				X:    config.ScreenW/2 - 4,
				Y:    config.ScreenH/2 - 4,
				VX:   1.5,
				VY:   2.5,
				Size: 8,
			},
		},
		over: &over.Over{},
	}); err != nil {
		log.Fatal(err)
	}
}
