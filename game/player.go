package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	X, Y    float64 // Координаты игрока
	VX, VY  float64 // Скорости по осям X и Y
	Jumping bool    // Состояние прыжка
}

func (p *Player) Update(platforms []Platform) {
	// Применяем гравитацию
	const gravity = 0.5
	p.VY += gravity

	// Обновляем положение игрока
	p.X += p.VX
	p.Y += p.VY

	// Проверяем столкновение с платформами
	for _, platform := range platforms {
		if p.X+20 > platform.X && p.X < platform.X+platform.Width && // Проверка по оси X
			p.Y+20 > platform.Y && p.Y+20 <= platform.Y+platform.Height { // Проверка по оси Y (нижняя грань игрока касается платформы)
			p.Y = platform.Y - 20 // Устанавливаем игрока на платформу
			p.VY = 0              // Обнуляем вертикальную скорость
			p.Jumping = false     // Игрок больше не прыгает
			break
		}
	}
}

// Draw отрисовывает игрока на экране
func (p *Player) Draw(screen *ebiten.Image, cameraY float64) {
	// Смещаем позицию игрока с учетом камеры
	screenX := p.X
	screenY := p.Y - cameraY

	// Определяем цвет игрока (белый)
	playerColor := color.White

	// Отрисовка квадрата игрока
	const playerSize = 20
	ebitenutil.DrawRect(screen, screenX, screenY, playerSize, playerSize, playerColor)
}
