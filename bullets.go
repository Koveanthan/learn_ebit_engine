package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const bulletVelocityPerSecond = 350.0

var bulletImage = mustLoadImage("assets/PNG/Lasers/laserRed01.png")

type Bullet struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

// var bounds = bulletImage.Bounds()
//
// var width = float64(bounds.Dx())
//
// var height = float64(bounds.Dy())
//
// var halfWidth = width / 2
//
// var halfHeight = height / 2

func (b *Bullet) Update() {
	// Need to caculate position with relation to rotated angle of the player

	bulletVelocity := bulletVelocityPerSecond / float64(ebiten.TPS())
	b.position.x = b.position.x + math.Sin(b.rotation)*bulletVelocity
	b.position.y = b.position.y + math.Cos(b.rotation)*-bulletVelocity

	// Need to add direction(can be taken from rotation?)
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	bounds := bulletImage.Bounds()
	width := float64(bounds.Dx())
	height := float64(bounds.Dy())
	halfW := width / 2
	halfH := height / 2

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(halfW, halfH)
	op.GeoM.Translate(b.position.x, b.position.y)
	screen.DrawImage(b.sprite, &op)
}

func NewBullet(x, y float64, rotation float64) *Bullet {
	return &Bullet{
		Vector{x, y},
		rotation,
		bulletImage,
	}
}
