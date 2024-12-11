package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Platform — структура для представления платформы
type Platform struct {
	X, Y, Width, Height float64
}

// Draw — метод отрисовки платформы
func (p *Platform) Draw(screen *ebiten.Image, cameraY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y-cameraY)
	ebitenutil.DrawRect(screen, p.X, p.Y-cameraY, p.Width, p.Height, ebiten.ColorM{})
}

// World — структура для управления состоянием мира
type World struct {
	Platforms []Platform
}

// NewWorld — конструктор для создания мира
func NewWorld() *World {
	return &World{
		Platforms: []Platform{
			{X: 50, Y: 550, Width: 200, Height: 20},
			{X: 300, Y: 450, Width: 200, Height: 20},
			{X: 100, Y: 350, Width: 200, Height: 20},
		},
	}
}

// Draw — метод отрисовки всех платформ мира
func (w *World) Draw(screen *ebiten.Image, cameraY float64) {
	for _, platform := range w.Platforms {
		platform.Draw(screen, cameraY)
	}
}

// GeneratePlatforms — генерация новых платформ по мере продвижения игрока
func (w *World) GeneratePlatforms(playerY float64) {
	// Условие генерации (например, если игрок поднялся выше 200 пикселей)
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
