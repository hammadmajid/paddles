package main

import (
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/objects/ball"
	"github.com/hammadmajid/paddle/internal/objects/paddle"
	"github.com/hammadmajid/paddle/internal/states/menu"
	"github.com/hammadmajid/paddle/internal/states/over"
	"github.com/hammadmajid/paddle/internal/states/play"
	"image/color"
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
	screen.Fill(color.RGBA{R: 20, G: 20, B: 20})
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
			Top: paddle.Paddle{
				X: config.ScreenW/2 - config.PaddleW/2, // center horizontally
				Y: 0.05 * config.ScreenH,               // near the top
				W: config.PaddleW,
				H: config.PaddleH,
			},
			Right: paddle.Paddle{
				X: config.ScreenW - 0.05*config.ScreenW - config.PaddleH, // near right edge
				Y: config.ScreenH/2 - config.PaddleW/2,                   // centered vertically
				W: config.PaddleH,
				H: config.PaddleW,
			},
			Bottom: paddle.Paddle{
				X: config.ScreenW/2 - config.PaddleW/2,                   // center horizontally
				Y: config.ScreenH - 0.05*config.ScreenH - config.PaddleH, // near the bottom
				W: config.PaddleW,
				H: config.PaddleH,
			},
			Left: paddle.Paddle{
				X: 0.05 * config.ScreenW,               // near left edge
				Y: config.ScreenH/2 - config.PaddleW/2, // centered vertically
				W: config.PaddleH,
				H: config.PaddleW,
			},

			Speed: 3,

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
