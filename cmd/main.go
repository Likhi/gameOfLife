package main

import (
	"fmt"
	gol "gameOfLife"
	"time"
)

//todo more oscillator test cases
//todo update same space on terminal screen (i.e no new line)

const (
	_period     = 500 * time.Millisecond
	_iterations = 50
)

func main() {
	//g := gol.NewGrid(61, 100)
	g := gol.NewOscillator(gol.BeaconPeriodTwo)
	fmt.Println("start")
	g.Print()
	for i := 0; i < _iterations; i++ {
		fmt.Println("Iteration:", i) // Show which iteration we're on
		g = g.Run()                  // Run the simulation once
		g.Print()                    // Show simulation after one iteration
		time.Sleep(_period)          // Slow your roll
	}
}
