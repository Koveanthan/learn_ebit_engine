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
	}
	if err := ebiten.RunGame(g); err != nil {
		println("invoking panic", err)
		panic(err)
	}
}

type Game struct {
	player Player
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
