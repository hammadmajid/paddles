package assets

import (
	_ "embed"
	"fmt"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed fonts/array/Array-Bold.otf
var fontBytes []byte

var Face font.Face

func init() {
	var err error
	Face, err = loadFont()
	if err != nil {
		panic(fmt.Sprintf("Failed to load font: %v", err))
	}
}

func loadFont() (font.Face, error) {
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create font face: %w", err)
	}

	return face, nil
}
