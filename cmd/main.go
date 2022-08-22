package main

import (
	"fmt"
	gol "gameOfLife"
)

func main() {
	fmt.Println("oh yeah, bb")
	g := gol.NewGrid(10, 10)
	//fmt.Println(g)
	//g.Print()
	history := []gol.Grid{g}

	const _iterations = 1

	for i := 0; i < _iterations; i++ {
		fmt.Println(i)
		history = append(history, g.Run()) // runs iteration and saves it to the history
		//history[0].Print()
	}

}
