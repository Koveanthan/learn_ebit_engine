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

var rotationSpeed = -0.2 + rand.Float64()*0.04

type Meteor struct {
	position Vector
	movement Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewMeteor() *Meteor {
	randAngle := rand.Float64() * math.Pi * 2
	velocity := 0.25 + rand.Float64()*1.5

	position := Vector{
		x: target.x + math.Cos(randAngle)*radius,
		y: target.y + math.Sin(randAngle)*radius,
	}
	direction := Vector{
		x: target.x - position.x,
		y: target.y - position.y,
	}

	normalizeDirection := direction.Normalize()

	movement := Vector{
		x: normalizeDirection.x * velocity,
		y: normalizeDirection.y * velocity,
	}
	return &Meteor{
		position: position,
		movement: movement,
		sprite:   meteors[rand.Intn(len(meteors))],
	}
}

func (m *Meteor) Update() {
	m.position.x += m.movement.x
	m.position.y += m.movement.y
	m.rotation += rotationSpeed
	println("Position.x: ", m.position.x, " Position.y: ", m.position.y, " Rotation: ", m.rotation)
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	op := ebiten.DrawImageOptions{}
	halfW := float64(bounds.Dx() / 2)
	halfH := float64(bounds.Dy() / 2)

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	// op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)
	op.GeoM.Translate(m.position.x, m.position.y)
	screen.DrawImage(m.sprite, &op)
}
