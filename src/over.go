package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Over struct{}

func (o Over) Update() bool {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		return true
	}
	return false
}

func (o Over) Draw(screen *ebiten.Image) {
	y := 100
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Game Over"), 100, y)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("> Enter <"), 100, y+40)
}
