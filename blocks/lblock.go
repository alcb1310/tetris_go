package blocks

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/alcb1310/tetris_go/colors"
	"github.com/alcb1310/tetris_go/position"
)

type LBlock struct {
	Id            int
	cellSize      int
	rotationState int
	colors        []rl.Color
	cells         map[int][]position.Position
	rowOffset     int
	colOffset     int
}

func NewLBlock() *LBlock {
	c := colors.NewColor()
	cell := make(map[int][]position.Position)

	cell[0] = []position.Position{
		{Row: 0, Col: 2},
		{Row: 1, Col: 0},
		{Row: 1, Col: 1},
		{Row: 1, Col: 2},
	}
	cell[1] = []position.Position{
		{Row: 0, Col: 1},
		{Row: 1, Col: 1},
		{Row: 2, Col: 1},
		{Row: 2, Col: 2},
	}
	cell[2] = []position.Position{
		{Row: 1, Col: 0},
		{Row: 1, Col: 1},
		{Row: 1, Col: 2},
		{Row: 2, Col: 0},
	}
	cell[3] = []position.Position{
		{Row: 0, Col: 0},
		{Row: 2, Col: 1},
		{Row: 1, Col: 1},
		{Row: 2, Col: 1},
	}

	b := &LBlock{
		cellSize:      30,
		Id:            1,
		colors:        c.GetAllColors(),
		cells:         cell,
		rotationState: 0,
		rowOffset:     0,
		colOffset:     0,
	}

	b.Move(0, 3)

	return b
}

func (b *LBlock) Draw() {
	tiles := b.GetCellPositions()

	for _, tile := range tiles {
		rl.DrawRectangle(
			int32(tile.Col*b.cellSize)+1,
			int32(tile.Row*b.cellSize)+1,
			int32(b.cellSize)-1,
			int32(b.cellSize)-1,
			b.colors[b.Id],
		)
	}
}

func (b *LBlock) GetCellPositions() []position.Position {
	tiles := b.cells[b.rotationState]
	movedTiles := make([]position.Position, 0)

	for _, pos := range tiles {
		newPosition := position.Position{
			Row: pos.Row + b.rowOffset,
			Col: pos.Col + b.colOffset,
		}

		movedTiles = append(movedTiles, newPosition)
	}

	return movedTiles
}

func (b *LBlock) Move(rows int, cols int) {
	b.rowOffset += rows
	b.colOffset += cols
}

func (b *LBlock) Rotate() {
  b.rotationState = (b.rotationState + 1) % len(b.cells)
}

func (b *LBlock) UndoRotate() {
	b.rotationState--

	if b.rotationState < 0 {
		b.rotationState = len(b.cells) - 1
	}
}
