package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
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
	text.Draw(screen, fmt.Sprintf("Game Over"), Face, 100, y, color.White)
	text.Draw(screen, fmt.Sprintf("> Enter <"), Face, 100, y+40, color.White)
}
