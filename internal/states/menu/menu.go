package menu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hammadmajid/paddle/assets"
	"image/color"
)

type Menu struct {
	Options []string
	Index   int
}

func (m *Menu) Draw(screen *ebiten.Image) {
	titleY := 60
	text.Draw(screen, "Paddles!", assets.Face, 100, titleY, color.White)

	for i, opt := range m.Options {
		y := 120 + i*40
		if i == m.Index {
			text.Draw(screen, fmt.Sprintf("> %s <", opt), assets.Face, 100, y, color.White)
		} else {
			text.Draw(screen, opt, assets.Face, 120, y, color.White)
		}
	}
}

func (m *Menu) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.Index++
		if m.Index >= len(m.Options) {
			m.Index = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.Index--
		if m.Index < 0 {
			m.Index = len(m.Options) - 1
		}
	}
}
