package main

// TODO: refactor for better code organization
// TODO: add convenience method for applying function to every cell

import (
	"fmt"
	"strconv"
	"strings"
)

type operator func(currCell cell) cell

type coordinate struct {
	i int
	j int
}

func (c coordinate) add(i int, j int) coordinate {
	return coordinate{c.i + i, c.j + j}
}


type cell struct {
	data      int
	location  coordinate
	neighbors []coordinate
}

func (c cell) nextState(numAliveNeigbors int) cell {
	isAlive := c.data == 1

	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	if isAlive && numAliveNeigbors > 3 {
		return cell{0, c.location, c.neighbors}
	}

	// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	if isAlive && numAliveNeigbors < 2 {
		return cell{0, c.location, c.neighbors}
	}
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	if !isAlive && numAliveNeigbors == 3 {
		return cell{1, c.location, c.neighbors}
	}

	// Any live cell with two or three live neighbours lives on to the next generation.
	return c
}

func (c cell) getPrintable() string {
	stringified := strconv.Itoa(c.data)
	if c.data == 1 {
		return makeYellow(stringified)
	}
	return makeBlack(stringified)
}

type gameBoard struct {
	size     int
	viewport int
	state    [][]cell
}

func (board gameBoard) updateCells(function operator) {
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			currentCell := board.state[i][j]
			board.state[i][j] = function(currentCell)
		}
	}
}

func (board gameBoard) nextBoard() gameBoard {
	nextState := make([][]cell, board.size)
	for i := 0; i < board.size; i++ {
		row := make([]cell, board.size)
		for j := 0; j < board.size; j++ {
			cell := board.getCell(i, j)
			numAlive := 0
			for _, neigbor := range cell.neighbors {
				numAlive += board.getCell(neigbor.i, neigbor.j).data
			}
			row[j] = cell.nextState(numAlive)
		}
		nextState[i] = row
	}
	board.state = nextState
	return board
}

func (b gameBoard) identifyNeighbors(coord coordinate) []coordinate {
	neighbors := []coordinate{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			potentialNeighbor := coord.add(i, j)
			if b.inBounds(potentialNeighbor) && potentialNeighbor != coord {
				neighbors = append(neighbors, potentialNeighbor)
			}
		}
	}
	return neighbors
}

func (b gameBoard) print() {
	dataOnly := make([]string, b.size)
	for i := 0; i < b.size; i++ {
		row := make([]string, b.size)
		for j := 0; j < b.size; j++ {
			row[j] = b.state[i][j].getPrintable()
		}
		dataOnly[i] = strings.Join(row, " ")
	}
	fmt.Print(strings.Join(dataOnly, "\n"))
}

func (b gameBoard) inBounds(coord coordinate) bool {
	onBoard := func(p int) bool { return p >= 0 && p < b.size }
	return onBoard(coord.i) && onBoard(coord.j)
}

func (b gameBoard) getCell(i int, j int) cell {
	return b.state[i][j]
}

func makeBlack(str string) string {
	return "\u001b[30m" + str + "\u001b[39m"
}

func makeYellow(str string) string {
	return "\u001b[33m" + str + "\u001b[39m"
}
