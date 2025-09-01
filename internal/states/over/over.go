package over

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hammadmajid/paddle/assets"
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/score"
)

type Over struct{}

func (o Over) Update() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return true
	}
	return false
}

func (o Over) Draw(screen *ebiten.Image) {
	y := 100
	text.Draw(screen, fmt.Sprintf("Game Over"), assets.Face, 100, y, config.ColorText)
	text.Draw(screen, fmt.Sprintf("Score >> %d", score.Get()), assets.Face, 100, y+60, config.ColorText)
	text.Draw(screen, fmt.Sprintf("> Enter <"), assets.Face, 100, y+100, config.ColorText)
}
