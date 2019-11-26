package main

import (
	"math/rand"
	"fmt"
)

type simulationParams struct {
	ticks       int
	boardSize   int
	renderDelay int
}

type coordinate struct {
	i int
	j int
}

type cell struct {
	data     int
	location coordinate
	neigbors []coordinate
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
	board := initBoard(params.boardSize)
	printBoard(board)
}

func initBoard(size int) [][]cell {
	getRandomInt := func() int { return rand.Intn(2) }
	board := make([][]cell, size)
	for i := range board {
		board[i] = make([]cell, size)
		for j := range board[i] {
			location := coordinate{i, j}
			neighbors := identifyNeighbors(location, size)
			board[i][j] = cell{getRandomInt(), location, neighbors}
		}
	}
	return board
}

func identifyNeighbors(loc coordinate, size int) []coordinate {
	return nil
}


func printBoard(board [][]cell) {
	dataOnly:= make([][]int, len(board))
	for i := range board {
		for j:= range board[i] {
			//FIXME: this is printing pointer values.  
			dataOnly[i][j] = board[i][j].data
		}
	}
	fmt.Print(dataOnly)
}