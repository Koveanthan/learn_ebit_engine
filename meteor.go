package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var meteors = mustLoadImages("assets/PNG/Meteors")

var target = Vector{
	x: ScreenWidth / 2,
	y: ScreenHeight / 2,
}

var radius = ScreenWidth / 2.0

type Meteor struct {
	position Vector
	sprite   *ebiten.Image
}

func NewMeteor() *Meteor {
	randAngle := rand.Float64() * math.Pi * 2
	return &Meteor{
		position: Vector{
			x: target.x + math.Cos(randAngle)*radius,
			y: target.y + math.Sin(randAngle)*radius,
		},
		sprite: meteors[rand.Intn(len(meteors))],
	}
}

func (m *Meteor) Update() {}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.x, m.position.y)
	screen.DrawImage(m.sprite, &op)
}
