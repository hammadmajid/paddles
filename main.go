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
	topPaddlePosX, topPaddlePosY    float32 // top paddle position
	topPaddleWidth, topPaddleHeight float32 // top paddle size
	btmPaddlePosX, btmPaddlePosY    float32 // bottom paddle position
	btmPaddleWidth, btmPaddleHeight float32 // bottom paddle size
	paddleSpeed                     float32 // movement speed of both paddles
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.topPaddlePosX += g.paddleSpeed
		g.btmPaddlePosX += -g.paddleSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.topPaddlePosX += -g.paddleSpeed
		g.btmPaddlePosX += g.paddleSpeed
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 0})

	vector.DrawFilledRect(screen, g.topPaddlePosX, g.topPaddlePosY, g.topPaddleWidth, g.topPaddleHeight, color.White, true)
	vector.DrawFilledRect(screen, g.btmPaddlePosX, g.btmPaddlePosY, g.btmPaddleWidth, g.btmPaddleHeight, color.White, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		topPaddlePosX:   screenW / 2,
		topPaddlePosY:   0.05 * float32(screenH),
		topPaddleWidth:  paddleW,
		topPaddleHeight: paddleH,
		btmPaddlePosX:   screenW / 2,
		btmPaddlePosY:   float32(screenH) - paddleH - (0.05 * float32(screenH)),
		btmPaddleWidth:  paddleW,
		btmPaddleHeight: paddleH,
		paddleSpeed:     3,
	}); err != nil {
		log.Fatal(err)
	}
}
