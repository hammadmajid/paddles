package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type Menu struct {
	options []string
	index   int
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for i, opt := range m.options {
		y := 100 + i*40
		if i == m.index {
			text.Draw(screen, fmt.Sprintf("> %s <", opt), Face, 100, y, color.White)
		} else {
			text.Draw(screen, opt, Face, 120, y, color.White)
		}
	}
}

func (m *Menu) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.index++
		if m.index >= len(m.options) {
			m.index = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.index--
		if m.index < 0 {
			m.index = len(m.options) - 1
		}
	}
}
