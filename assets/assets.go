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

//go:embed music/8-bit-laser-151672.mp3
var hitEffectBytes []byte

//go:embed music/gamestart-272829.mp3
var gameStartBytes []byte

const sampleRate = 44100

var AudioContext *audio.Context
var BGMPlayer *audio.Player
var HitEffectPlayer *audio.Player
var GameStartPlayer *audio.Player

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

	bgmStream, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(bgmBytes))
	if err != nil {
		log.Fatal(err)
	}

	bgmLoop := audio.NewInfiniteLoop(bgmStream, bgmStream.Length())

	BGMPlayer, err = AudioContext.NewPlayer(bgmLoop)
	if err != nil {
		log.Fatal(err)
	}

	hitEffectStream, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(hitEffectBytes))
	if err != nil {
		log.Fatal(err)
	}
	HitEffectPlayer, err = AudioContext.NewPlayer(hitEffectStream)
	if err != nil {
		log.Fatal(err)
	}

	gameStartStream, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(gameStartBytes))
	if err != nil {
		log.Fatal(err)
	}
	GameStartPlayer, err = AudioContext.NewPlayer(gameStartStream)
	if err != nil {
		log.Fatal(err)
	}

	BGMPlayer.SetVolume(0.6)
	BGMPlayer.Play()
}
