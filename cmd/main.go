package main

import (
	"fmt"
	gol "gameOfLife"
)

func main() {
	//g := gol.NewGrid(5, 5)
	g := gol.NewOscillator("o")
	history := []gol.Grid{g}
	fmt.Println("start")
	g.Print()

	const _iterations = 1

	for i := 0; i < _iterations; i++ {
		fmt.Println(i)
		next := g.Run()
		history = append(history, next) // runs iteration and saves it to the history

		//next.Print()
		//time.Sleep(1 * time.Second)
	}

}
