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
	_dead state = iota
	_alive
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
		return cell{s: _alive}
	}
	return cell{s: _dead}
}

func (x *Grid) Print() {
	for _, row := range *x {
		for _, col := range row {
			if col.s == _dead {
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

	// create grid "next" for holding next cells
	next := make(Grid, len(*x))
	copy(next, *x)

	//todo run simulation on *Grid, copy it and return it
	for m, row := range *x {
		for n, c := range row {
			//iterate through each cell c in the original grid and calculate next cell
			upLeft := x.getState(m-1, n-1)
			up := x.getState(m-1, n)
			upRight := x.getState(m+1, n+1)
			left := x.getState(m, n-1)
			right := x.getState(m, n+1)
			downLeft := x.getState(m+1, n-1)
			down := x.getState(m, n+1)
			downRight := x.getState(m+1, n+1)

			neighbors := map[string]state{"upLeft": upLeft, "up": up, "upRight": upRight,
				"left": left, "right": right,
				"downLeft": downLeft, "down": down, "downRight": downRight}

			liveCount := 0
			for _, v := range neighbors {
				if v == _alive {
					liveCount += 1
				}
			}

			nextCell := cell{}
			//Any live cell with two or three live neighbours survives.
			if c.s == _alive {
				if (liveCount == 2) || (liveCount == 3) {
					nextCell = c
				} else {
					nextCell = cell{s: _dead}
				}
			}

			//Any dead cell with three live neighbours becomes a live cell.
			if c.s == _dead {
				if (liveCount == 3) {
					nextCell = cell{s: _alive}
				} else {
					nextCell = cell{s: _dead}
				}
			}

			//All other live cells die in the next generation. Similarly, all other dead cells stay dead.
			//else {
			//	nextCell = cell{s: _dead}
			//}

			//put nextCell in next grid at same location
			next.updateState(m, n, nextCell)

		}
	}

	return next
}

func (x *Grid) getState(m int, n int) state {
	width := len((*x)[0])
	height := len(*x)

	var mWrapped int
	var nWrapped int

	if m < 0 {
		mWrapped = height + m
	} else {
		mWrapped = m % height
	}

	if n < 0 {
		nWrapped = width + n
	} else {
		nWrapped = n % width
	}

	return (*x)[mWrapped][nWrapped].s
}

func (x *Grid) updateState(m int, n int, c cell) {
	(*x)[m][n] = c
}
