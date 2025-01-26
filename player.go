package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	x float64
	y float64
}

type Player struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewPlayer(x, y float64, imagePath string) *Player {
	p := Player{
		position: Vector{x: x, y: y},
		sprite:   mustLoadImage(imagePath),
	}

	bounds := p.sprite.Bounds()
	p.position.x -= float64(bounds.Dx() / 2)
	p.position.y -= float64(bounds.Dy() / 2)

	return &p
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx() / 2)
	halfH := float64(bounds.Dy() / 2)

	op := ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.x, p.position.y)
	screen.DrawImage(p.sprite, &op)
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
}
