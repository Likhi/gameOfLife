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

type Osc int

const (
	BlinkPeriodTwo Osc = iota + 1
	BeaconPeriodTwo
	ToadPeriodTwo
	PulsarPeriod3
	PentaDecathalonPeriod15
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
			row = append(row, newCell())
		}
		ret = append(ret, row)
	}

	return ret
}

// NewOscillator creates a grid seeded with an oscillator
// 5x5 grid with "blinker".
// todo instead of a string, use an enum (const) that determine oscillator type
func NewOscillator(o Osc) Grid {
	switch o {
	case BlinkPeriodTwo:
		row1 := []cell{deadCell(), deadCell(), deadCell(), deadCell(), deadCell()}
		row2 := []cell{deadCell(), deadCell(), deadCell(), deadCell(), deadCell()}
		row3 := []cell{deadCell(), aliveCell(), aliveCell(), aliveCell(), deadCell()}
		row4 := []cell{deadCell(), deadCell(), deadCell(), deadCell(), deadCell()}
		row5 := []cell{deadCell(), deadCell(), deadCell(), deadCell(), deadCell()}

		blinkPeriodTwo := Grid{row1, row2, row3, row4, row5}
		return blinkPeriodTwo
	case BeaconPeriodTwo:
		row1 := []cell{deadCell(), deadCell(), deadCell(), deadCell(), deadCell(), deadCell()}
		row2 := []cell{deadCell(), deadCell(), deadCell(), deadCell(), deadCell(), deadCell()}
		row3 := []cell{deadCell(), deadCell(), aliveCell(), aliveCell(), aliveCell(), deadCell()}
		row4 := []cell{deadCell(), aliveCell(), aliveCell(), aliveCell(), deadCell(), deadCell()}
		row5 := []cell{deadCell(), deadCell(), deadCell(), deadCell(), deadCell(), deadCell()}
		row6 := []cell{deadCell(), deadCell(), deadCell(), deadCell(), deadCell(), deadCell()}

		beaconPeriodTwo := Grid{row1, row2, row3, row4, row5, row6}
		return beaconPeriodTwo
	default:
		panic("Unsupported oscillator!")
	}

}

// deadGrid returns a new Grid with the same size of x but with only _dead cell{}s
func (x *Grid) deadGrid() Grid {
	width := len((*x)[0])
	height := len(*x)

	ret := NewGrid(height, width)

	return ret
}

// newCell returns a cell randomly selected to be alive or dead
func newCell() cell {
	value := rand.Intn(100)
	//fmt.Println(value)
	if value < _probabilityAlive {
		return cell{s: _alive}
	}
	return cell{s: _dead}
}

// deadCell returns a cell{} with the state _dead
func deadCell() cell {
	return cell{s: _dead}
}

// aliveCell returns a cell{} with the state _alive
func aliveCell() cell {
	return cell{s: _alive}
}

// Print "pretty prints" a Grid
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

// Run calculates the next iteration on the Grid and returns a deep copy of it
func (x *Grid) Run() Grid {
	// create grid "next" for holding next cells
	next := x.deadGrid()

	//todo run simulation on *Grid, copy it and return it
	for m, row := range *x {
		for n, c := range row {
			//iterate through each cell c in the original grid and calculate next cell
			upLeft := x.getState(m-1, n-1)
			up := x.getState(m-1, n)
			upRight := x.getState(m-1, n+1)
			left := x.getState(m, n-1)
			right := x.getState(m, n+1)
			downLeft := x.getState(m+1, n-1)
			down := x.getState(m+1, n)
			downRight := x.getState(m+1, n+1)

			neighborStates := map[string]state{
				"upLeft": upLeft, "up": up, "upRight": upRight,
				"left": left, "right": right,
				"downLeft": downLeft, "down": down, "downRight": downRight}

			liveCount := 0
			for _, v := range neighborStates {
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
				if liveCount == 3 {
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
	//fmt.Println("grid size", height, "x", width)

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

	//fmt.Println("mUW", m, "nUW", n)
	//fmt.Println("mW", mWrapped, "nW", nWrapped, "state", (*x)[mWrapped][nWrapped].s)

	return (*x)[mWrapped][nWrapped].s
}

func (x *Grid) updateState(m int, n int, c cell) {
	(*x)[m][n] = c
}
