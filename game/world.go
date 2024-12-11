package game

type Platform struct {
	X, Y, Width, Height float64
}

var platforms = []Platform{
	{X: 100, Y: 500, Width: 200, Height: 20},
	{X: 300, Y: 400, Width: 200, Height: 20},
}

func (p *Player) CheckCollisions(platforms []Platform) {
	for _, platform := range platforms {
		if p.X+p.Width > platform.X && p.X < platform.X+platform.Width &&
			p.Y+p.Height > platform.Y && p.Y < platform.Y+platform.Height {
			p.Y = platform.Y - p.Height
			p.VY = 0
			p.Jumping = false
		}
	}
}
