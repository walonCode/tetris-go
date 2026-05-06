package block

import "github.com/walonCode/tetris-go/internals/position"

// L Block
func NewLBlock() *Block {
	b := New()
	b.Id = 1
	b.Cells = [][]position.Position{
			{
				{Row: 0, Column: 2}, {Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 1, Column: 2},
			},
			{
				{Row: 0, Column: 1}, {Row: 1, Column: 1}, {Row: 2, Column: 1}, {Row: 2, Column: 2},
			},
			{
				{Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 2, Column: 0},
			},
			{
				{Row: 0, Column: 0}, {Row: 0, Column: 1}, {Row: 1, Column: 1}, {Row: 2, Column: 1},
			},
	}

	b.Move(0, 3)
	return b
}

// J Block
func NewJBlock() *Block {
	b := New()
	b.Id = 2
	b.Cells = [][]position.Position{
			{
				{Row: 0, Column: 0}, {Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 1, Column: 2},
			},
			{
				{Row: 0, Column: 1}, {Row: 0, Column: 2}, {Row: 1, Column: 1}, {Row: 2, Column: 1},
			},
			{
				{Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 2, Column: 2},
			},
			{
				{Row: 0, Column: 1}, {Row: 1, Column: 1}, {Row: 2, Column: 0}, {Row: 2, Column: 1},
			},
	}

	b.Move(0, 3)
	return b
}

// I Block
func NewIBlock() *Block {
	b := New()
	b.Id = 3
	b.Cells = [][]position.Position{
			{
				{Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 1, Column: 3},
			},
			{
				{Row: 0, Column: 2}, {Row: 1, Column: 2}, {Row: 2, Column: 2}, {Row: 3, Column: 2},
			},
			{
				{Row: 2, Column: 0}, {Row: 2, Column: 1}, {Row: 2, Column: 2}, {Row: 2, Column: 3},
			},
			{
				{Row: 0, Column: 1}, {Row: 1, Column: 1}, {Row: 2, Column: 1}, {Row: 3, Column: 1},
			},
	}

	b.Move(-1, 3)
	return b
}

// O Block
func NewOBlock() *Block {
	b := New()
	b.Id = 4
	b.Cells = [][]position.Position{
			{
				{Row: 0, Column: 0}, {Row: 0, Column: 1}, {Row: 1, Column: 0}, {Row: 1, Column: 1},
			},
	}

	b.Move(0, 4)
	return b
}

// S Block
func NewSBlock() *Block {
	b := New()
	b.Id = 5
	b.Cells = [][]position.Position{
			{
				{Row: 0, Column: 1}, {Row: 0, Column: 2}, {Row: 1, Column: 0}, {Row: 1, Column: 1},
			},
			{
				{Row: 0, Column: 1}, {Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 2, Column: 2},
			},
			{
				{Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 2, Column: 0}, {Row: 2, Column: 1},
			},
			{
				{Row: 0, Column: 0}, {Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 2, Column: 1},
			},
	}

	b.Move(0, 3)
	return b
}

// T Block
func NewTBlock() *Block {
	b := New()
	b.Id = 6
	b.Cells = [][]position.Position{
			{
				{Row: 0, Column: 1}, {Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 1, Column: 2},
			},
			{
				{Row: 0, Column: 1}, {Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 2, Column: 1},
			},
			{
				{Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 2, Column: 1},
			},
			{
				{Row: 0, Column: 1}, {Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 2, Column: 1},
			},
	}

	b.Move(0, 3)
	return b
}

// Z Block
func NewZBlock() *Block {
	b := New()
	b.Id = 7
	b.Cells = [][]position.Position{
			{
				{Row: 0, Column: 0}, {Row: 0, Column: 1}, {Row: 1, Column: 1}, {Row: 1, Column: 2},
			},
			{
				{Row: 0, Column: 2}, {Row: 1, Column: 1}, {Row: 1, Column: 2}, {Row: 2, Column: 1},
			},
			{
				{Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 2, Column: 1}, {Row: 2, Column: 2},
			},
			{
				{Row: 0, Column: 1}, {Row: 1, Column: 0}, {Row: 1, Column: 1}, {Row: 2, Column: 0},
			},
	}

	b.Move(0, 3)
	return b
}
