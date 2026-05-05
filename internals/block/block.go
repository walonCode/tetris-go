package block

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/walonCode/tetris-go/internals/colors"
)

type Position struct{}

type Block struct {
	cellSize int 
	rotationState int 
	rowOffset int 
	columnOffset int
	colors []rl.Color

	//public 
	Id int
	Cells map[int][]Position
}

func (b *Block)Draw(offsetX, offsetY int){
	titles := []int{}
	for _, items := range  titles {
		rl.DrawRectangle(item.column * b.cellSize + offsetX, item.row * b.cellSize + offsetY, int32(b.cellSize) - 1, int32(b.cellSize - 1), b.colors[b.Id])
	}
}

func (b *Block)Move(rows, columns int){
	b.rowOffset += rows
	b.columnOffset += columns
}

func (b *Block)GetCellPostion()[]Position{
	titles := b.Cells[b.rotationState]
	var movedTitles []Position

	for _, title := range titles {
		newPos := Position{}
		movedTitles.PushBack(newPos)
	}

	return movedTitles
}

func (b *Block)Rotate(){
	b.rotationState++
	if b.rotationState == int(len(b.Cells)) {
		b.rotationState = 0
	}
}

func (b *Block) UndoRotation(){
	b.rotationState--
	if b.rotationState == -1 {
		b.rotationState = len(b.Cells) - 1
	}
}

func New()*Block{
	return &Block{
		cellSize: 30,
		rotationState: 0,
		colors: colors.GetColors(),
		rowOffset: 0,
		columnOffset: 0,
	}
}
