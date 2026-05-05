package block

import "github.com/walonCode/tetris-go/internals/position"

// L Block
func NewLBlock() *Block {
	b := &Block{
		Id: 1,
		Cells: [][]position.Position{
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
		},
	}

	b.Move(0, 3)
	return b
}

// J Block
func NewJBlock() *Block {
	b := &Block{
		Id: 2,
		Cells: [][]position.Position{
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
		},
	}

	b.Move(0, 3)
	return b
}

// I Block
func NewIBlock() *Block {
	b := &Block{
		Id: 3,
		Cells: [][]position.Position{
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
		},
	}

	b.Move(-1, 3)
	return b
}

// O Block
func NewOBlock() *Block {
	b := &Block{
		Id: 4,
		Cells: [][]position.Position{
			{
				{Row: 0, Column: 0}, {Row: 0, Column: 1}, {Row: 1, Column: 0}, {Row: 1, Column: 1},
			},
		},
	}

	b.Move(0, 4)
	return b
}

// S Block
func NewSBlock() *Block {
	b := &Block{
		Id: 5,
		Cells: [][]position.Position{
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
		},
	}

	b.Move(0, 3)
	return b
}

// T Block
func NewTBlock() *Block {
	b := &Block{
		Id: 6,
		Cells: [][]position.Position{
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
		},
	}

	b.Move(0, 3)
	return b
}

// Z Block
func NewZBlock() *Block {
	b := &Block{
		Id: 7,
		Cells: [][]position.Position{
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
		},
	}

	b.Move(0, 3)
	return b
}