package colors

import rl "github.com/gen2brain/raylib-go/raylib"

type Color interface {
	GetAllColors() []rl.Color
}

type color struct {
	DarkGrey rl.Color
	Green    rl.Color
	Red      rl.Color
	Orange   rl.Color
	Yellow   rl.Color
	Purple   rl.Color
	Cyan     rl.Color
	Blue     rl.Color
}

func NewColor() Color {
	return &color{
		DarkGrey: rl.NewColor(80, 80, 80, 255),
		Green:    rl.NewColor(0, 255, 0, 255),
		Red:      rl.NewColor(255, 0, 0, 255),
		Orange:   rl.NewColor(255, 165, 0, 255),
		Yellow:   rl.NewColor(255, 255, 0, 255),
		Purple:   rl.NewColor(128, 0, 128, 255),
		Cyan:     rl.NewColor(0, 255, 255, 255),
		Blue:     rl.NewColor(0, 0, 255, 255),
	}
}

func (c *color) GetAllColors() []rl.Color {
	return []rl.Color{
		c.DarkGrey,
		c.Green,
		c.Red,
		c.Orange,
		c.Yellow,
		c.Purple,
		c.Cyan,
		c.Blue,
	}
}
