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

	paddleSpeed float32 // movement speed of both paddles

	targetX, targetY   float32 // target position
	targetVX, targetVY float32 // target velocity
	targetSize         float32 // target size
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.topPaddlePosX += g.paddleSpeed
		g.btmPaddlePosX += -g.paddleSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.topPaddlePosX += -g.paddleSpeed
		g.btmPaddlePosX += g.paddleSpeed
	}

	const padding = 16

	// clamp top paddle
	if g.topPaddlePosX < padding {
		g.topPaddlePosX = padding
	}
	if g.topPaddlePosX+g.topPaddleWidth > float32(screenW-padding) {
		g.topPaddlePosX = float32(screenW) - g.topPaddleWidth - padding
	}

	// clamp bottom paddle
	if g.btmPaddlePosX < padding {
		g.btmPaddlePosX = padding
	}
	if g.btmPaddlePosX+g.btmPaddleWidth > float32(screenW-padding) {
		g.btmPaddlePosX = float32(screenW) - g.btmPaddleWidth - padding
	}

	g.targetX += g.targetVX
	g.targetY += g.targetVY

	// bounce left/right walls
	if g.targetX < padding {
		g.targetX = padding
		g.targetVX = -g.targetVX
	}
	if g.targetX+g.targetSize > float32(screenW-padding) {
		g.targetX = float32(screenW-padding) - g.targetSize
		g.targetVX = -g.targetVX
	}

	// bounce top paddle
	if g.targetY <= g.topPaddlePosY+g.topPaddleHeight &&
		g.targetX+g.targetSize >= g.topPaddlePosX &&
		g.targetX <= g.topPaddlePosX+g.topPaddleWidth {
		g.targetY = g.topPaddlePosY + g.topPaddleHeight
		g.targetVY = -g.targetVY
	}

	// bounce bottom paddle
	if g.targetY+g.targetSize >= g.btmPaddlePosY &&
		g.targetX+g.targetSize >= g.btmPaddlePosX &&
		g.targetX <= g.btmPaddlePosX+g.btmPaddleWidth {
		g.targetY = g.btmPaddlePosY - g.targetSize
		g.targetVY = -g.targetVY
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 0})

	vector.DrawFilledRect(screen, g.topPaddlePosX, g.topPaddlePosY, g.topPaddleWidth, g.topPaddleHeight, color.White, true)
	vector.DrawFilledRect(screen, g.btmPaddlePosX, g.btmPaddlePosY, g.btmPaddleWidth, g.btmPaddleHeight, color.White, true)

	vector.DrawFilledRect(screen, g.targetX, g.targetY, g.targetSize, g.targetSize, color.RGBA{200, 0, 0, 255}, true)
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

		paddleSpeed: 3,

		targetX:    float32(screenW/2 - 4),
		targetY:    float32(screenH/2 - 4),
		targetVX:   1.5,
		targetVY:   2.5,
		targetSize: 8,
	}); err != nil {
		log.Fatal(err)
	}
}
