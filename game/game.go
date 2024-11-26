package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/alcb1310/tetris_go/blocks"
	"github.com/alcb1310/tetris_go/constants"
	gr "github.com/alcb1310/tetris_go/grid"
)

type Game interface {
	Run()
	draw()
	handleInput()
	moveBlockLeft()
	moveBlockRight()
	moveBlockDown()
}

type game struct {
	lastUpdateTime float64
	gameOver       bool
	grid           gr.Grid
	currentBlock   blocks.Block
}

func NewGame() Game {
	return &game{
		lastUpdateTime: 0,
		gameOver:       false,
		grid:           gr.NewGrid(),
		currentBlock:   blocks.NewZBlock(),
	}
}

func (g *game) Run() {
	g.handleInput()

	rl.BeginDrawing()
	rl.ClearBackground(constants.NEARLY_BLACK)

	g.draw()

	rl.EndDrawing()
}

func (g *game) draw() {
	g.grid.Draw()
	g.currentBlock.Draw()
}

func (g *game) handleInput() {
	key := rl.GetKeyPressed()

	switch key {
	case rl.KeyLeft:
		g.moveBlockLeft()
	case rl.KeyRight:
		g.moveBlockRight()
	case rl.KeyDown:
		g.moveBlockDown()
	}
}

func (g *game) moveBlockLeft() {
	g.currentBlock.Move(0, -1)
  if g.isBlockOutside() {
    g.currentBlock.Move(0, 1)
  }
}

func (g *game) moveBlockRight() {
	g.currentBlock.Move(0, 1)
  if g.isBlockOutside() {
    g.currentBlock.Move(0, -1)
  }
}

func (g *game) moveBlockDown() {
	g.currentBlock.Move(1, 0)
  if g.isBlockOutside() {
    g.currentBlock.Move(-1, 0)
  }
}

func (g *game) isBlockOutside() bool {
  tiles := g.currentBlock.GetCellPositions()

  for _, tile := range tiles {
    if g.grid.IsCellOutside(tile.Row, tile.Col) {
      return true
    }
  }

  return false
}
