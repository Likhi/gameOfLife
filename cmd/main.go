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

	const _iterations = 5

	for i := 0; i < _iterations; i++ {
		fmt.Println(i)
		next := g.Run()
		history = append(history, next) // runs iteration and saves it to the history
		next.Print()
	}

}
