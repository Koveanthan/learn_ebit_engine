package main

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

func mustLoadImage(path string) *ebiten.Image {
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

func mustLoadImages(path string) []*ebiten.Image {
	var result []*ebiten.Image
	folder, err := assets.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range folder {
		if file.IsDir() {
			panic("folder in path")
		}
		result = append(result, mustLoadImage(path+"/"+file.Name()))
	}

	return result
}
