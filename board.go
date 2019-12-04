package main

// TODO: refactor for better code organization
	// export gameboard et al.
	// break cells and boards into separate modules?

// TODO: tests :)
// TODO: smooth out drawing; redraw changes only
// TODO: memoize cell state; only recompute on change
// TODO:completely re-implement with sparse arrays :)

import (
	"fmt"
	"strconv"
	"strings"
)

type cellMapper func(currCell cell) cell

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
	viewport int
	state    [][]cell
}

func (board gameBoard) nextBoard() gameBoard {
	return board
}

func (board gameBoard) identifyNeighbors(coord coordinate) []coordinate {
	neighbors := []coordinate{}
	return neighbors
}

func (board gameBoard) print() {
	dataOnly := make([]string, board.viewport)
	for i := 0; i < board.viewport; i++ {
		row := make([]string, board.viewport)
		for j := 0; j < board.viewport; j++ {
			row[j] = board.state[i][j].getPrintable()
		}
		dataOnly[i] = strings.Join(row, " ")
	}
	fmt.Print(strings.Join(dataOnly, "\n") + "\r")
}

func (board gameBoard) getCell(coord coordinate) cell {
	return board.state[coord.i][coord.j]
}

func makeBlack(str string) string {
	return "\u001b[30m" + str + "\u001b[39m"
}

func makeYellow(str string) string {
	return "\u001b[33m" + str + "\u001b[39m"
}
