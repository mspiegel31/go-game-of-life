package main

import "fmt"

type simulationParams struct {
	ticks       int
	boardSize   int
	renderDelay int
}

var DEFAULTS simulationParams = simulationParams{500, 30, 100}

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
}
