package gameOfLife

import "math/rand"

// put helper functions in here

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
