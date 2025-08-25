package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenW  = 640.0
	screenH  = 480.0
	hPaddleW = 64
	hPaddleH = 12
	vPaddleW = 12
	vPaddleH = 64
)

type Game struct {
	top    Paddle
	right  Paddle
	bottom Paddle
	left   Paddle

	paddleSpeed float32 // movement speed of both paddles

	ball Ball
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.top.X += g.paddleSpeed
		g.right.Y += g.paddleSpeed

		g.bottom.X -= g.paddleSpeed
		g.left.Y -= g.paddleSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.top.X -= g.paddleSpeed
		g.right.Y -= g.paddleSpeed

		g.bottom.X += g.paddleSpeed
		g.left.Y += g.paddleSpeed
	}

	const padding = 16

	// clamp paddle movement
	g.top.Clamp(screenW, screenH, padding, padding)
	g.right.Clamp(screenW, screenH, padding, padding)
	g.bottom.Clamp(screenW, screenH, padding, padding)
	g.left.Clamp(screenW, screenH, padding, padding)

	// ball movment and bounce
	g.ball.Move()
	g.ball.CheckPaddle(g.top, true)
	g.ball.CheckPaddle(g.right, true)
	g.ball.CheckPaddle(g.bottom, false)
	g.ball.CheckPaddle(g.left, false)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 0})

	vector.DrawFilledRect(screen, g.top.X, g.top.Y, g.top.W, g.top.H, color.White, true)
	vector.DrawFilledRect(screen, g.right.X, g.right.Y, g.right.W, g.right.H, color.White, true)
	vector.DrawFilledRect(screen, g.bottom.X, g.bottom.Y, g.bottom.W, g.bottom.H, color.White, true)
	vector.DrawFilledRect(screen, g.left.X, g.left.Y, g.left.W, g.left.H, color.White, true)

	vector.DrawFilledRect(screen, g.ball.X, g.ball.Y, g.ball.Size, g.ball.Size, color.RGBA{200, 0, 0, 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		top: Paddle{
			X: screenW/2 - hPaddleW/2, // center horizontally
			Y: 0.05 * screenH,         // near the top
			W: hPaddleW,
			H: hPaddleH,
		},
		right: Paddle{
			X: screenW - 0.05*screenW - hPaddleH, // near right edge
			Y: screenH/2 - hPaddleW/2,            // centered vertically
			W: hPaddleH,
			H: hPaddleW,
		},
		bottom: Paddle{
			X: screenW/2 - hPaddleW/2,            // center horizontally
			Y: screenH - 0.05*screenH - hPaddleH, // near the bottom
			W: hPaddleW,
			H: hPaddleH,
		},
		left: Paddle{
			X: 0.05 * screenW,         // near left edge
			Y: screenH/2 - hPaddleW/2, // centered vertically
			W: hPaddleH,
			H: hPaddleW,
		},

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
