package entities

import (
	"math/rand"
	"time"
)

// World Struct yes
type World struct {
	Row   int
	Col   int
	Cells [][]Cell
}

// GenCells : Generate world matrix
func GenCells(row, col int) [][]Cell {
	cells := make([][]Cell, row)
	for i := 0; i < row; i++ {
		cells[i] = make([]Cell, col)
	}
	return cells
}

// NewWorld : Generate new world
func NewWorld(row, col int) World {
	var w World

	w.Row = row
	w.Col = col

	w.Cells = GenCells(row, col)

	return w
}

// Line preset
func (w World) Line() {
	// var midCol int
	// var midRow int
	// if w.Col%2 == 0 {
	// 	midCol = (w.Col / 2)
	// } else {
	// 	midCol = int(math.Ceil(float64(w.Col) / 2.0))
	// }
	// if w.Row%2 == 0 {
	// 	midRow = (w.Row / 2)
	// } else {
	// 	midRow = int(math.Ceil(float64(w.Row) / 2.0))
	// }

	midCol := (w.Col / 2)
	midRow := (w.Row / 2)

	w.Cells[midRow][midCol-4].SetAlive(true)
	w.Cells[midRow][midCol-3].SetAlive(true)
	w.Cells[midRow][midCol-2].SetAlive(true)
	w.Cells[midRow][midCol-1].SetAlive(true)
	w.Cells[midRow][midCol].SetAlive(true)
	w.Cells[midRow][midCol+1].SetAlive(true)
	w.Cells[midRow][midCol+2].SetAlive(true)
	w.Cells[midRow][midCol+3].SetAlive(true)
	w.Cells[midRow][midCol+4].SetAlive(true)
	w.Cells[midRow][midCol+5].SetAlive(true)

}

// Exploder preset
func (w World) Exploder() {
	// var midCol int
	// var midRow int
	// if w.Col%2 == 0 {
	// 	midCol = (w.Col / 2)
	// } else {
	// 	midCol = int(math.Ceil(float64(w.Col) / 2.0))
	// }
	// if w.Row%2 == 0 {
	// 	midRow = (w.Row / 2)
	// } else {
	// 	midRow = int(math.Ceil(float64(w.Row) / 2.0))
	// }

	midCol := (w.Col / 2)
	midRow := (w.Row / 2)

	w.Cells[midRow-2][midCol-2].SetAlive(true)
	w.Cells[midRow-2][midCol].SetAlive(true)
	w.Cells[midRow-2][midCol+2].SetAlive(true)
	w.Cells[midRow-1][midCol-2].SetAlive(true)
	w.Cells[midRow-1][midCol+2].SetAlive(true)
	w.Cells[midRow][midCol-2].SetAlive(true)
	w.Cells[midRow][midCol+2].SetAlive(true)
	w.Cells[midRow+1][midCol-2].SetAlive(true)
	w.Cells[midRow+1][midCol+2].SetAlive(true)
	w.Cells[midRow+2][midCol-2].SetAlive(true)
	w.Cells[midRow+2][midCol].SetAlive(true)
	w.Cells[midRow+2][midCol+2].SetAlive(true)

}

//Tumbler preset
func (w World) Tumbler() {
	// var midCol int
	// var midRow int
	// if w.Col%2 == 0 {
	// 	midCol = (w.Col / 2)
	// } else {
	// 	midCol = int(math.Ceil(float64(w.Col) / 2.0))
	// }
	// if w.Row%2 == 0 {
	// 	midRow = (w.Row / 2)
	// } else {
	// 	midRow = int(math.Ceil(float64(w.Row) / 2.0))
	// }

	midCol := (w.Col / 2)
	midRow := (w.Row / 2)

	w.Cells[midRow-2][midCol-1].SetAlive(true)
	w.Cells[midRow-2][midCol-2].SetAlive(true)
	w.Cells[midRow-2][midCol+1].SetAlive(true)
	w.Cells[midRow-2][midCol+2].SetAlive(true)
	w.Cells[midRow-1][midCol-1].SetAlive(true)
	w.Cells[midRow-1][midCol-2].SetAlive(true)
	w.Cells[midRow-1][midCol+1].SetAlive(true)
	w.Cells[midRow-1][midCol+2].SetAlive(true)
	w.Cells[midRow][midCol-1].SetAlive(true)
	w.Cells[midRow][midCol+1].SetAlive(true)
	w.Cells[midRow+1][midCol-3].SetAlive(true)
	w.Cells[midRow+1][midCol-1].SetAlive(true)
	w.Cells[midRow+1][midCol+1].SetAlive(true)
	w.Cells[midRow+1][midCol+3].SetAlive(true)
	w.Cells[midRow+2][midCol-3].SetAlive(true)
	w.Cells[midRow+2][midCol-1].SetAlive(true)
	w.Cells[midRow+2][midCol+1].SetAlive(true)
	w.Cells[midRow+2][midCol+3].SetAlive(true)
	w.Cells[midRow+3][midCol-3].SetAlive(true)
	w.Cells[midRow+3][midCol-2].SetAlive(true)
	w.Cells[midRow+3][midCol+2].SetAlive(true)
	w.Cells[midRow+3][midCol+3].SetAlive(true)

}

// Glider preset
func (w World) Glider() {
	// var midCol int
	// var midRow int
	// if w.Col%2 == 0 {
	// 	midCol = (w.Col / 2)
	// } else {
	// 	midCol = int(math.Ceil(float64(w.Col) / 2.0))
	// }
	// if w.Row%2 == 0 {
	// 	midRow = (w.Row / 2)
	// } else {
	// 	midRow = int(math.Ceil(float64(w.Row) / 2.0))
	// }

	midCol := (w.Col / 2)
	midRow := (w.Row / 2)

	w.Cells[midRow-1][midCol].SetAlive(true)
	w.Cells[midRow][midCol+1].SetAlive(true)
	w.Cells[midRow+1][midCol-1].SetAlive(true)
	w.Cells[midRow+1][midCol].SetAlive(true)
	w.Cells[midRow+1][midCol+1].SetAlive(true)
}

// Start : generate random beginning
func (w World) Start() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	times := r.Intn(w.Col * w.Row)

	for t := 0; t < times; t++ {
		i, j := rand.Intn(w.Row), rand.Intn(w.Col)

		w.Cells[i][j].SetAlive(true)
	}

}

// GetNearby : Return a list of all neighbours for a given cell
func (w World) GetNearby(r, c int) []Cell {
	var cells []Cell

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			newR, newC := r+i, c+j

			if newR >= w.Row {
				newR = 0
			} else if newR < 0 {
				newR = w.Row - 1
			}

			if newC >= w.Col {
				newC = 0
			} else if newC < 0 {
				newC = w.Col - 1
			}

			if !(newR == r && newC == c) {
				cells = append(cells, w.Cells[newR][newC])
			}
		}
	}

	return cells
}

// NextGen : Advence the world 1 generation
func (w World) NextGen() World {
	newCells := GenCells(w.Row, w.Col)

	for i := 0; i < w.Row; i++ {
		for j := 0; j < w.Col; j++ {
			newCells[i][j] = w.Cells[i][j].Step(w.GetNearby(i, j))
		}
	}

	w.Cells = newCells

	return w
}
