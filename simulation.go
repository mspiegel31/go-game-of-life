package main

import (
	"os"
	"os/exec"
	"fmt"
	"time"
	"github.com/mike/go-game-of-life/board"
)

type simulationParams struct {
	ticks       int
	boardSize   int
	renderDelay int
}

var DEFAULTS simulationParams = simulationParams{500, 30, 70}

func NewSimulationParams(args ...int) simulationParams {
	if len(args) == 1 {
		return simulationParams{args[0], DEFAULTS.boardSize, DEFAULTS.renderDelay}
	}
	if len(args) == 2 {
		return simulationParams{args[0], args[1], DEFAULTS.renderDelay}
	}
	if len(args) == 3 {
		return simulationParams{args[0], args[1], args[2]}
	}

	return DEFAULTS
}

func runSimulation(params simulationParams) {
	fmt.Printf("running simulation with params %#v \n", params)
	board := board.Init(params.boardSize)
	board.Print()

	for index := 0; index < params.ticks; index++ {
		time.Sleep(time.Duration(params.renderDelay) * time.Millisecond)
		board = board.NextBoard()
		clearScreen()
		board.Print()
	}

}

func clearScreen() {
	c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}