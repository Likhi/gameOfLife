package gameOfLife

import (
	"fmt"
	"math/rand"
	"time"
)

// put types in here
type Grid [][]cell

type cell struct {
	s state
}

type state int

const (
	Dead state = iota
	Alive
)

const (
	_probabilityAlive = 10
)

func NewGrid(rows int, cols int) Grid {
	ret := Grid{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rows; i++ {
		var row []cell
		for j := 0; j < cols; j++ {
			row = append(row, NewCell())
		}
		ret = append(ret, row)
	}

	return ret
}

// NewCell returns a cell randomly selected to be alive or dead
func NewCell() cell {
	value := rand.Intn(100)
	//fmt.Println(value)
	if value < _probabilityAlive {
		return cell{s: Alive}
	}
	return cell{s: Dead}
}

func (x *Grid) Print() {
	for _, row := range *x {
		for _, col := range row {
			if col.s == Dead {
				fmt.Print("_ ")
			} else {
				fmt.Print("* ")
			}
		}

		// line break
		fmt.Println("")
	}
}

// Run calculates the next iteration on the Grid and returns a copy of it
func (x *Grid) Run() Grid {
	//runs one iteration of the simulation on the grid
	//returns a copy grid after simulation

	// create grid for holding next cells
	next := make(Grid, len(*x))
	copy(next, *x)

	//todo run simulation on *Grid, copy it and return it
	for m, row := range *x {
		for n, cell := range row {
			//todo iterate through each cell in the original grid and calculate next cell
			nextCell := NewCell()

			//put nextCell in next grid at same location
			x.updateState(m,n,nextCell)

		}
	}

	// copy updated grid into next
	copy(next, *x)

	return next
}

func (x *Grid) getState(m int, n int) state {
	return (*x)[m][n].s
}

func (x *Grid) updateState(m int, n int, c cell) {

}

func