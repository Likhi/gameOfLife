package main

import (
	"fmt"
	gol "gameOfLife"
	"time"
)

func main() {
	g := gol.NewGrid(10, 10)
	//g := gol.NewOscillator("o")
	history := []gol.Grid{}
	fmt.Println("start")
	g.Print()

	const _iterations = 5

	for i := 0; i < _iterations; i++ {
		// save last grid state
		history = append(history, g)
		fmt.Println(i)
		g = g.Run()
		g.Print()
		time.Sleep(1 * time.Second)
	}

}
