package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenHeight = 800
	ScreenWidth  = 600
)

func main() {
	g := &Game{
		player:  NewPlayer(),
		timer:   NewTimer(5 * time.Second),
		meteors: []*Meteor{},
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type Game struct {
	player  *Player
	timer   *MeteorTimer
	meteors []*Meteor
}

func (g *Game) Update() error {
	g.player.Update()

	g.timer.update()

	if g.timer.isReady() {
		g.timer.reset()

		meteor := NewMeteor()
		g.meteors = append(g.meteors, meteor)
	}

	for _, meteor := range g.meteors {
		meteor.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	for _, meteor := range g.meteors {
		meteor.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
