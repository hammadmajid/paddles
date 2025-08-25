package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenW = 640.0
	screenH = 480.0
	paddleW = 64
	paddleH = 12
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlay
)

type Game struct {
	state GameState

	menu *Menu
	play *Play
}

func (g *Game) Update() error {

	switch g.state {
	case StateMenu:
		g.menu.Update()
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			if g.menu.options[g.menu.index] == "Play" {
				g.state = StatePlay
			}
			if g.menu.options[g.menu.index] == "Quit" {
				return ebiten.Termination
			}
		}
	case StatePlay:
		g.play.Update()
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
	}
}

//goland:noinspection GoUnusedParameter
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Paddles")
	if err := ebiten.RunGame(&Game{
		menu: &Menu{
			options: []string{"Play", "Quit"},
			index:   0,
		},
		play: &Play{
			top: Paddle{
				X: screenW/2 - paddleW/2, // center horizontally
				Y: 0.05 * screenH,        // near the top
				W: paddleW,
				H: paddleH,
			},
			right: Paddle{
				X: screenW - 0.05*screenW - paddleH, // near right edge
				Y: screenH/2 - paddleW/2,            // centered vertically
				W: paddleH,
				H: paddleW,
			},
			bottom: Paddle{
				X: screenW/2 - paddleW/2,            // center horizontally
				Y: screenH - 0.05*screenH - paddleH, // near the bottom
				W: paddleW,
				H: paddleH,
			},
			left: Paddle{
				X: 0.05 * screenW,        // near left edge
				Y: screenH/2 - paddleW/2, // centered vertically
				W: paddleH,
				H: paddleW,
			},

			paddleSpeed: 3,

			ball: Ball{
				X:    screenW/2 - 4,
				Y:    screenH/2 - 4,
				VX:   1.5,
				VY:   2.5,
				Size: 8,
			},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
