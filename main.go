package main

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &Game{
		playerPosition: Vector{250, 250},
	}
	if err := ebiten.RunGame(g); err != nil {
		println("invoking panic", err)
		panic(err)
	}
}

//go:embed assets/*
var assets embed.FS

var playerSprite = mustLoadImage("assets/PNG/playerShip1_blue.png")

type Vector struct {
	x float64
	y float64
}

type Game struct {
	playerPosition Vector
}

func (g *Game) Update() error {
	g.playerPosition.x += 0.5
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.x, g.playerPosition.y)
	screen.DrawImage(playerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

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
