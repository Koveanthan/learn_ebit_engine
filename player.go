package main

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var playerSprite = mustLoadImage("./assets/png/playerShip1_blue.png")

func mustLoadImage(path string) image.Image {
	f, err := assets.Open(path)
	if err != nil {
		panic(err)
	}

	image, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(image)
}
