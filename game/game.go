package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/alcb1310/tetris_go/constants"
	gr "github.com/alcb1310/tetris_go/grid"
)

type Game interface {
	Run()
	draw()
}

type game struct {
	lastUpdateTime float64
	gameOver       bool
	grid           gr.Grid
}

func NewGame() Game {
	return &game{
		lastUpdateTime: 0,
		gameOver:       false,
		grid:           gr.NewGrid(),
	}
}

func (g *game) Run() {
	rl.BeginDrawing()
	rl.ClearBackground(constants.NEARLY_BLACK)

	g.draw()

	rl.EndDrawing()
}

func (g *game) draw() {
	g.grid.Draw()
}
