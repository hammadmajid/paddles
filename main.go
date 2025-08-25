package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenW = 640.0
	screenH = 480.0
	paddleW = 64
	paddleH = 12
)

type Game struct {
	top    Paddle
	bottom Paddle

	paddleSpeed float32 // movement speed of both paddles

	ball Ball
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.top.X += g.paddleSpeed
		g.bottom.X -= g.paddleSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.top.X -= g.paddleSpeed
		g.bottom.X += g.paddleSpeed
	}

	const padding = 16

	// clamp paddle movement
	g.top.Clamp(screenW, padding)
	g.bottom.Clamp(screenW, padding)

	// ball movment and bounce
	g.ball.Move()
	g.ball.CheckWalls(screenW, screenH, 16)
	g.ball.CheckPaddle(g.top, true)
	g.ball.CheckPaddle(g.bottom, false)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 0})

	vector.DrawFilledRect(screen, g.top.X, g.top.Y, g.top.W, g.top.H, color.White, true)
	vector.DrawFilledRect(screen, g.bottom.X, g.bottom.Y, g.bottom.W, g.bottom.H, color.White, true)

	vector.DrawFilledRect(screen, g.ball.X, g.ball.Y, g.ball.Size, g.ball.Size, color.RGBA{200, 0, 0, 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		top:    NewPaddle(screenW/2, 0.05*screenH),
		bottom: NewPaddle(screenW/2, screenH-paddleH-(0.05*screenH)),

		paddleSpeed: 3,

		ball: Ball{
			X:    screenW/2 - 4,
			Y:    screenH/2 - 4,
			VX:   1.5,
			VY:   2.5,
			Size: 8,
		},
	}); err != nil {
		log.Fatal(err)
	}
}
