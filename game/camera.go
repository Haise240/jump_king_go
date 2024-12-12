package game

// Camera управляет положением камеры в игре
type Camera struct {
	Y float64 // Текущее положение камеры по оси Y
}

// Update обновляет положение камеры на основе координаты игрока
func (c *Camera) Update(playerY float64) {
	const offset = 200 // Смещение камеры от игрока
	c.Y = playerY - offset
	if c.Y < 0 {
		c.Y = 0 // Камера не опускается ниже нуля
	}
}
