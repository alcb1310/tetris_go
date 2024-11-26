package main

import (
	"github.com/alcb1310/tetris_go/constants"
	"github.com/alcb1310/tetris_go/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(constants.FPS)
	g := game.NewGame()

	for !rl.WindowShouldClose() {
		g.Run()
	}
}
