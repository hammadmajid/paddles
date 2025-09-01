package config

import "image/color"

const (
	ScreenW         = 640.0
	ScreenH         = 480.0
	PaddingDefault  = 16.0
	PaddingEdge     = 10.0
	PaddleSpeed     = 3.0
	BallSpeed       = 1.5
	BallSpeedY      = 2.5
	BallSize        = 8.0
	PaddleThickness = 12.0
	PaddleLength    = 64.0
)

// Cattppuccin Mocha Palette
// See: https://catppuccin.com/palette/
var (
	ColorBase     = color.RGBA{R: 30, G: 30, B: 46}
	ColorText     = color.RGBA{R: 205, G: 214, B: 244}
	ColorLavender = color.RGBA{R: 180, G: 190, B: 254}
	ColorRed      = color.RGBA{R: 243, G: 139, B: 168}
)
