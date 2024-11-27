package blocks

import "github.com/alcb1310/tetris_go/position"

type Block interface {
	Draw()
	GetCellPositions() []position.Position
	Move(rows int, cols int)
	Rotate()
	UndoRotate()
}
