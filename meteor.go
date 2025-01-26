package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var meteors = mustLoadImages("assets/PNG/Meteors")

type Meteor struct {
	position Vector
	sprite   *ebiten.Image
}

func NewMeteor(x, y float64) *Meteor {
	return &Meteor{
		position: Vector{x: x, y: y},
		sprite:   meteors[rand.Intn(len(meteors))],
	}
}

func (m *Meteor) Update() {}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.x, m.position.y)
	screen.DrawImage(m.sprite, &op)
}
