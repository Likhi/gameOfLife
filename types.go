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

func NewCell() cell {
	value := rand.Intn(100)
	//fmt.Println(value)
	if value < 10 {
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

func (x *Grid) Run() Grid {
	// runs one iteration of the simulation on the grid
	// returns a copy grid after simulation

	r := make(Grid, len(*x)) //just for testing

	// todo run simulation on *Grid, copy it and return it

	copy(r, *x)

	return r
}
