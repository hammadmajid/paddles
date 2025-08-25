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
	right  Paddle
	bottom Paddle
	left   Paddle

	paddleSpeed float32 // movement speed of both paddles

	ball Ball
}

func (g *Game) Update() error {
	// Store previous positions to calculate velocity
	topOldX := g.top.X
	rightOldY := g.right.Y
	bottomOldX := g.bottom.X
	leftOldY := g.left.Y

	// assign keys to control each paddle
	topControls := Controls{NegX: ebiten.KeyA, PosX: ebiten.KeyD}
	rightControls := Controls{NegY: ebiten.KeyArrowUp, PosY: ebiten.KeyArrowDown}
	bottomControls := Controls{NegX: ebiten.KeyArrowLeft, PosX: ebiten.KeyArrowRight}
	leftControls := Controls{NegY: ebiten.KeyW, PosY: ebiten.KeyS}

	g.top.Move(topControls)
	g.right.Move(rightControls)
	g.bottom.Move(bottomControls)
	g.left.Move(leftControls)

	const padding = 16

	// clamp paddle movement
	g.top.Clamp(screenW, screenH, padding, padding)
	g.right.Clamp(screenW, screenH, padding, padding)
	g.bottom.Clamp(screenW, screenH, padding, padding)
	g.left.Clamp(screenW, screenH, padding, padding)

	// Calculate paddle velocities
	g.top.VX = g.top.X - topOldX
	g.top.VY = 0
	g.right.VX = 0
	g.right.VY = g.right.Y - rightOldY
	g.bottom.VX = g.bottom.X - bottomOldX
	g.bottom.VY = 0
	g.left.VX = 0
	g.left.VY = g.left.Y - leftOldY

	// ball movement and bounce
	g.ball.Move()
	g.ball.CheckHorizontalPaddle(g.top, true)
	g.ball.CheckHorizontalPaddle(g.bottom, false)
	g.ball.CheckVerticalPaddle(g.left, true)
	g.ball.CheckVerticalPaddle(g.right, false)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 20, G: 20, B: 20})

	vector.DrawFilledRect(screen, g.top.X, g.top.Y, g.top.W, g.top.H, color.White, true)
	vector.DrawFilledRect(screen, g.right.X, g.right.Y, g.right.W, g.right.H, color.White, true)
	vector.DrawFilledRect(screen, g.bottom.X, g.bottom.Y, g.bottom.W, g.bottom.H, color.White, true)
	vector.DrawFilledRect(screen, g.left.X, g.left.Y, g.left.W, g.left.H, color.White, true)

	vector.DrawFilledRect(screen, g.ball.X, g.ball.Y, g.ball.Size, g.ball.Size, color.RGBA{R: 200}, true)
}

//goland:noinspection GoUnusedParameter
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Paddles")
	if err := ebiten.RunGame(&Game{
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
	}); err != nil {
		log.Fatal(err)
	}
}
