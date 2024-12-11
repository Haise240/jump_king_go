package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y    float64 // Позиция
	VX, VY  float64 // Скорость
	Jumping bool    // В прыжке или нет
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.VX = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.VX = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !p.Jumping {
		p.VY = -10
		p.Jumping = true
	}
	// Гравитация
	p.VY += 0.5
	p.X += p.VX
	p.Y += p.VY
	// Ограничения
	if p.Y > 500 {
		p.Y = 500
		p.VY = 0
		p.Jumping = false
	}
}
