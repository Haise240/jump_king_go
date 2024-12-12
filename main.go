package main

import (
	"Haise240/jump_king_go/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *game.Player
	World  *game.World
	Camera game.Camera
}

func NewGame() *Game {
	return &Game{
		Player: &game.Player{X: 150, Y: 480, VX: 0, VY: 0, Jumping: false}, // Позиция над платформой
		World:  game.NewWorld(),
		Camera: game.Camera{Y: 0},
	}
}

func (g *Game) Update() error {
	g.Player.Update(g.World.Platforms)
	g.Camera.Update(g.Player.Y)
	g.World.GeneratePlatforms(g.Player.Y)
	g.World.CleanupPlatforms(g.Camera.Y)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.World.Draw(screen, g.Camera.Y)
	g.Player.Draw(screen, g.Camera.Y)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
