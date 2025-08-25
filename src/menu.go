package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Menu struct {
	options []string
	index   int
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for i, opt := range m.options {
		y := 100 + i*40
		if i == m.index {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("> %s <", opt), 100, y)
		} else {
			ebitenutil.DebugPrintAt(screen, opt, 120, y)
		}
	}
}

func (m *Menu) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		m.index++
		if m.index >= len(m.options) {
			m.index = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		m.index--
		if m.index < 0 {
			m.index = len(m.options) - 1
		}
	}
}
