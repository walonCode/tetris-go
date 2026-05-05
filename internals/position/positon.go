package position

type Position struct {
	Row int
	Column int
}

func New(row, column int )*Position{
	return &Position{
		Row: row,
		Column: column,
	}
}