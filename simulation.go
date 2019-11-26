package main

import (
	"fmt"
	"math/rand"
)

type simulationParams struct {
	ticks       int
	boardSize   int
	renderDelay int
}

var DEFAULTS simulationParams = simulationParams{500, 30, 100}
var MARGIN int = 6

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
	fmt.Printf("running simulation with params %#v", params)
	board := initBoard(params.boardSize)
	board.print()
}

func initBoard(size int) board {
	state := make([][]cell, size)
	for i := range state {
		state[i] = make([]cell, size)
		for j := range state[i] {
			location := coordinate{i, j}
			neighbors := identifyNeighbors(location, size)
			state[i][j] = cell{rand.Intn(2), location, neighbors}
		}
	}
	return board{size, size - MARGIN, state}
}

func identifyNeighbors(loc coordinate, size int) []coordinate {
	neighbors := make([]coordinate, size)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			//TODO: implement me!
		}
	}
	return neighbors
}
