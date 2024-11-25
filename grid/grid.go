package grid

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	cl "github.com/alcb1310/tetris_go/colors"
)

type Grid interface {
	Initialize()
	Print()
	Draw()
}

type grid struct {
	Grid [20][10]int

	numRows     int
	numCols     int
	cellSize    int
	colorValues []rl.Color
}

func NewGrid() Grid {
	c := cl.NewColor()

	g := &grid{
		numRows:     20,
		numCols:     10,
		cellSize:    30,
		colorValues: c.GetAllColors(),
	}

	g.Initialize()

	return g
}

func (g *grid) Initialize() {
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			g.Grid[i][j] = 0
		}
	}
}

func (g *grid) Print() {
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			fmt.Printf("%d ", g.Grid[i][j])
		}
		fmt.Println()
	}
}

func (g *grid) Draw() {
	for row := 0; row < g.numRows; row++ {
		for col := 0; col < g.numCols; col++ {
			cellColor := g.Grid[row][col]
			rl.DrawRectangle(
				int32(col*g.cellSize)+1,
				int32(row*g.cellSize)+1,
				int32(g.cellSize)-1,
				int32(g.cellSize)-1,
				g.colorValues[cellColor],
			)
		}
	}
}
