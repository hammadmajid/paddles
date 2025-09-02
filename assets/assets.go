package assets

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed fonts/array/Array-Bold.otf
var fontBytes []byte

var Face font.Face

//go:embed music/8_Bit_Surf_-_FesliyanStudios.com_-_David_Renda.mp3
var bgmBytes []byte

const sampleRate = 44100

var AudioContext *audio.Context
var BGMPlayer *audio.Player

func init() {
	// font
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

	// audio
	AudioContext = audio.NewContext(sampleRate)
	stream, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(bgmBytes))
	if err != nil {
		log.Fatal(err)
	}
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	BGMPlayer, err = AudioContext.NewPlayer(loop)
	if err != nil {
		log.Fatal(err)
	}

	BGMPlayer.Play()
}
