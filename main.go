package main

import (
	"Haise240/jump_king_go/game" // Пакет с логикой игры
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
		Player: &game.Player{X: 100, Y: 500, VX: 0, VY: 0, Jumping: false},
		World:  game.NewWorld(),
		Camera: game.Camera{Y: 0},
	}
}

func (g *Game) Update() error {
	g.Player.Update(g.World.Platforms)
	g.Camera.Update(g.Player.Y)

	// Генерация и очистка платформ
	g.World.GeneratePlatforms(g.Player.Y)
	g.World.CleanupPlatforms(g.Camera.Y)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.World.Draw(screen, g.Camera.Y)
	g.Player.Draw(screen, g.Camera.Y)
}

// Layout — задает размеры окна игры
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600 // Размеры окна: 800x600
}

// Точка входа в программу
func main() {
	game := NewGame() // Создаем объект игры

	// Запуск игрового цикла
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
