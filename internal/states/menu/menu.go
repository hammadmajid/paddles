package menu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hammadmajid/paddle/assets"
	"github.com/hammadmajid/paddle/internal/config"
)

type Menu struct {
	Options []string
	Index   int
}

func (m *Menu) Draw(screen *ebiten.Image) {
	titleY := 60
	text.Draw(screen, "Paddles!", assets.Face, 100, titleY, config.ColorText)

	for i, opt := range m.Options {
		y := 120 + i*40
		if i == m.Index {
			text.Draw(screen, fmt.Sprintf("> %s <", opt), assets.Face, 100, y, config.ColorText)
		} else {
			text.Draw(screen, opt, assets.Face, 120, y, config.ColorText)
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
