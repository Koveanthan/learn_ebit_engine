package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenHeight = 800
	ScreenWidth  = 600
)

func main() {
	playerSpriteURL := "assets/PNG/playerShip1_blue.png"
	g := &Game{
		player: *NewPlayer(ScreenWidth/2, ScreenHeight/2, playerSpriteURL),
		meteor: *NewMeteor(ScreenWidth/2, ScreenHeight/2),
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type Game struct {
	player Player
	meteor Meteor
}

func (g *Game) Update() error {
	g.player.Update()
	g.meteor.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	g.meteor.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
