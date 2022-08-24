package main

import (
	"fmt"
	gol "gameOfLife"
	"time"
)

//todo more oscillator test cases
//todo update same space on terminal screen (i.e no new line)

func main() {
	//g := gol.NewGrid(61, 100)
	g := gol.NewOscillator(gol.BeaconPeriodTwo)
	history := []gol.Grid{}
	fmt.Println("start")
	g.Print()

	const _iterations = 50

	for i := 0; i < _iterations; i++ {
		// save last grid state
		history = append(history, g)
		fmt.Println("Iteration:", i)
		g = g.Run()
		g.Print()
		time.Sleep(500 * time.Millisecond)
	}
}
