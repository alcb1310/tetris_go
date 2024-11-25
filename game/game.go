package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/alcb1310/tetris_go/constants"
)
type Game interface {
	Run()
}

type game struct {
	lastUpdateTime float64
	gameOver       bool
}

func NewGame() Game {
	return &game{
		lastUpdateTime: 0,
		gameOver:       false,
	}
}

func (g *game) Run() {
	rl.BeginDrawing()
	rl.ClearBackground(constants.NEARLY_BLACK)


	rl.EndDrawing()
}
}
