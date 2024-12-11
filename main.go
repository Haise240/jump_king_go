package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	// Логика игры (гравитация, коллизии)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Игрок
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(player.X, player.Y)
	screen.DrawImage(playerImage, op)

	// Платформы
	for _, platform := range platforms {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(platform.X, platform.Y)
		screen.DrawImage(platformImage, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600 // Размер окна
}

func main() {
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
