package main

import (
	_ "embed"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
)

//go:embed Array-Bold.otf
var fontBytes []byte

var Face font.Face

func init() {
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	Face, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
