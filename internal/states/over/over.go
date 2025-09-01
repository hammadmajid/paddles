package over

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hammadmajid/paddle/assets"
	"github.com/hammadmajid/paddle/internal/config"
	"github.com/hammadmajid/paddle/internal/interfaces"
)

type Over struct{}

func (o Over) Update() interfaces.StateTransition {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return interfaces.StateTransition{
			ShouldTransition: true,
			NewState:         interfaces.StateMenu,
		}
	}
	return interfaces.StateTransition{}
}

func (o Over) Draw(screen *ebiten.Image) {
	y := 100
	text.Draw(screen, fmt.Sprintf("Game Over"), assets.Face, 100, y, config.ColorText)
	text.Draw(screen, fmt.Sprintf("> Enter <"), assets.Face, 100, y+40, config.ColorText)
}
