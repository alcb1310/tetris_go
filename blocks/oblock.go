package blocks

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/alcb1310/tetris_go/colors"
	"github.com/alcb1310/tetris_go/position"
)

type OBlock struct {
	Id            int
	cellSize      int
	rotationState int
	colors        []rl.Color
	cells         map[int][]position.Position
	rowOffset     int
	colOffset     int
}

func NewOBlock() *OBlock {
	c := colors.NewColor()
	cell := make(map[int][]position.Position)

	cell[0] = []position.Position{
		{Row: 0, Col: 0},
		{Row: 0, Col: 1},
		{Row: 1, Col: 0},
		{Row: 1, Col: 1},
	}

	b := &OBlock{
		cellSize:      30,
		Id:            4,
		colors:        c.GetAllColors(),
		cells:         cell,
		rotationState: 0,
		rowOffset:     0,
		colOffset:     0,
	}

	b.Move(0, 4)

	return b
}

func (b *OBlock) Draw() {
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

func (b *OBlock) GetCellPositions() []position.Position {
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

func (b *OBlock) Move(rows int, cols int) {
	b.rowOffset += rows
	b.colOffset += cols
}

func (b *OBlock) Rotate() {
  b.rotationState = (b.rotationState + 1) % len(b.cells)
}

func (b *OBlock) UndoRotate() {
	b.rotationState--

	if b.rotationState < 0 {
		b.rotationState = len(b.cells) - 1
	}
}
