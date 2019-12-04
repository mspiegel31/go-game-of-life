package main

import (
	"fmt"
	"math/rand"
	"time"
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
	board := initBoard(params.boardSize)
	board.print()

	for index := 0; index < params.ticks; index++ {
		time.Sleep(time.Duration(params.renderDelay) * time.Millisecond)
		board = board.nextBoard()
		board.print()
	}

}

func initBoard(size int) gameBoard {
	anchor := coordinate{50, 50}
	aliveCells := make(map[coordinate]cell)
	board := gameBoard{anchor, size, aliveCells}

	// seed values in viewport randomly
	for i := anchor.i; i < anchor.i+size; i++ {
		for j := anchor.j; j < anchor.j+size; j++ {
			data := rand.Intn(2)
			if data == 1 {
				location := coordinate{i, j}
				aliveCells[location] = cell{data}
			}
		}
	}
	return board
}