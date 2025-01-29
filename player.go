package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var playerSprite = mustLoadImage("assets/PNG/playerShip1_blue.png")

type Vector struct {
	x float64
	y float64
}

func (v Vector) Normalize() Vector {
	magnitude := math.Sqrt(v.x*v.x + v.y*v.y)
	return Vector{v.x / magnitude, v.y / magnitude}
}

type Player struct {
	position Vector
	rotation float64
	bullets  []*Bullet
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	p := Player{
		position: Vector{x: ScreenWidth / 2, y: ScreenHeight / 2},
		bullets:  []*Bullet{},
		sprite:   playerSprite,
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

	for _, bullet := range p.bullets {
		bullet.Draw(screen)
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.bullets = append(p.bullets, NewBullet(p.position.x, p.position.y, p.rotation))
	}

	for _, bullet := range p.bullets {
		bullet.Update()
	}
}
