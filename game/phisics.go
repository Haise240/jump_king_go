package game

const Gravity = 0.5 // Сила гравитации

func ApplyGravity(vy *float64) {
	*vy += Gravity
}

func CheckCollision(playerX, playerY, playerWidth, playerHeight float64, platforms []Platform) (onPlatform bool, platformY float64) {
	for _, platform := range platforms {
		if playerX+playerWidth > platform.X && playerX < platform.X+platform.Width &&
			playerY+playerHeight > platform.Y && playerY+playerHeight <= platform.Y+5 { // "5" для небольшого допущения
			return true, platform.Y
		}
	}
	return false, 0
}

const JumpForce = -10 // Сила прыжка

func Jump(vy *float64, isJumping *bool) {
	if !*isJumping {
		*vy = JumpForce
		*isJumping = true
	}
}

func UpdatePlayerPosition(x, y, vx, vy *float64, platforms []Platform) {
	// Обновление позиции
	*x += *vx
	*y += *vy

	// Гравитация
	ApplyGravity(vy)

	// Коллизии
	onPlatform, platformY := CheckCollision(*x, *y, 50, 50, platforms) // 50x50 — размер игрока
	if onPlatform {
		*y = platformY - 50 // Устанавливаем позицию на платформу
		*vy = 0
	}
}
