package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Platform struct {
	X, Y   float64     // Координаты верхнего левого угла платформы
	Width  float64     // Ширина платформы
	Height float64     // Высота платформы
	Color  color.Color // Цвет платформы
}

type World struct {
	Platforms []Platform // Список платформ
}

// Draw — метод отрисовки платформы
func (p *Platform) Draw(screen *ebiten.Image, cameraY float64) {
	x := p.X
	y := p.Y - cameraY
	width := p.Width
	height := p.Height

	//  нужный цвет
	platformColor := color.RGBA{R: 100, G: 100, B: 100, A: 255} // Серый цвет

	// Отрисовываем прямоугольник платформы
	ebitenutil.DrawRect(screen, x, y, width, height, platformColor)
}

func NewWorld() *World {
	return &World{
		Platforms: []Platform{
			{X: 100, Y: 500, Width: 200, Height: 20, Color: color.RGBA{R: 128, G: 128, B: 128, A: 255}}, // Базовая платформа
		},
	}
}

// Draw отрисовывает платформы мира
func (w *World) Draw(screen *ebiten.Image, cameraY float64) {
	for _, platform := range w.Platforms {
		// Смещаем платформу с учетом камеры
		screenX := platform.X
		screenY := platform.Y - cameraY
		ebitenutil.DrawRect(screen, screenX, screenY, platform.Width, platform.Height, platform.Color)
	}
}

// GeneratePlatforms — генерация новых платформ по мере продвижения игрока
func (w *World) GeneratePlatforms(playerY float64) {
	// Условие генерации (например, если игрок поднялся выше 150 пикселей относительно последней платформы)
	if len(w.Platforms) > 0 && playerY < w.Platforms[len(w.Platforms)-1].Y-150 {
		newPlatform := Platform{
			X:      50 + float64(len(w.Platforms)*100%300), // Псевдослучайное смещение по X
			Y:      w.Platforms[len(w.Platforms)-1].Y - 100,
			Width:  200,
			Height: 20,
		}
		w.Platforms = append(w.Platforms, newPlatform)
	}
}

// CleanupPlatforms — удаляет платформы ниже видимого экрана
func (w *World) CleanupPlatforms(cameraY float64) {
	newPlatforms := w.Platforms[:0]
	for _, platform := range w.Platforms {
		if platform.Y-cameraY > -50 { // Условие: платформа должна быть в видимом диапазоне
			newPlatforms = append(newPlatforms, platform)
		}
	}
	w.Platforms = newPlatforms
}
