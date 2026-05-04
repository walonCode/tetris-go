package grid

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/walonCode/tetris-go/colors"
)

type Grid struct {
	numRows int 
	numCols int
	cellSize int
	colors []rl.Color 
	grid [20][10]int
}

//public function
func (g *Grid) Initialize(){
	for row := 0; row < g.numRows; row++{
		for col := 0; col < g.numCols; col++{
			g.grid[row][col] = 0
		}
	}
}

func (g *Grid)Print(){
	for row := 0; row < g.numRows; row++{
		for col := 0; col < g.numCols; col++{
			fmt.Printf("%d ",g.grid[row][col])
		}
		fmt.Print("")
	}
}

func (g *Grid)Draw(){
	for row := 0; row < g.numRows; row++{
		for col := 0; col < g.numCols; col++{
			cellValue := g.grid[row][col]
			rl.DrawRectangle(int32(col * g.cellSize) + 11, int32(row * g.cellSize) + 11, int32(g.cellSize) - 1, int32(g.cellSize) - 1, g.colors[cellValue] )
		}
	}
}

func (g *Grid) IsCellOutside(row,column int)bool{
	if row >= 0 && row < g.numRows && column >=0 && column < g.numCols{
		return false
	}
	return true
}

func (g *Grid) IsCellEmpty(row, column int)bool{
	if g.grid[row][column] == 0 {
		return true
	}

	return false
}

func (g *Grid)ClearFullRows()int{
	completed := 0
	for row := g.numRows - 1; row >= 0; row-- {
		if g.isRowFull(row){
			g.clearRow(row)
			completed++
		}else if completed > 0 {
			g.moveRowDown(row, completed)
		}
	}
	return completed
}

// private function
func (g *Grid) isRowFull(row int)bool{
	for col:=0; col<g.numCols;col++{
		if g.grid[row][col] == 0 {
			return false
		}
	}

	return true
}

func (g *Grid)clearRow(row int){
	for col:=0; col<g.numCols;col++{
		g.grid[row][col] = 0
	}
}
func (g *Grid)moveRowDown(row,numRows int){
	for col:=0; col<g.numCols;col++{
	 	g.grid[row + numRows][col] = g.grid[row][col];
        g.grid[row][col] = 0;
	}
}

func New()*Grid{
	return &Grid{
		numRows: 20,
		numCols: 10,
		cellSize: 30,
		colors: *colors.GetColors(),
	}
}