package main

// TODO: refactor for better code organization
// export gameboard et al.
// break cells and boards into separate modules?

// TODO: tests :)
// TODO: smooth out drawing; redraw changes only
// TODO: memoize cell state; only recompute on change
// TODO:completely re-implement with sparse arrays :)

import (
	"strings"
	"fmt"
	"strconv"
)

type coordinate struct {
	i int
	j int
}

func (c coordinate) add(i int, j int) coordinate {
	return coordinate{c.i + i, c.j + j}
}

type cell struct {
	data int
}

func (c cell) getPrintable() string {
	stringified := strconv.Itoa(c.data)
	if c.data == 1 {
		return makeYellow(stringified)
	}
	return makeBlack(stringified)
}

func (c cell) nextState(numAliveNeigbors int) cell {
	isAlive := c.data == 1

	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	if isAlive && numAliveNeigbors > 3 {
		return cell{0}
	}

	// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	if isAlive && numAliveNeigbors < 2 {
		return cell{0}
	}
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	if !isAlive && numAliveNeigbors == 3 {
		return cell{1}
	}

	// Any live cell with two or three live neighbours lives on to the next generation.
	return c
}

type gameBoard struct {
	viewAnchor coordinate
	viewport   int
	aliveCells map[coordinate]cell
}

func (board gameBoard) nextBoard() gameBoard {
	return board
}

func (board gameBoard) identifyNeighbors(coord coordinate) []coordinate {
	neighbors := []coordinate{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			neighbor := coord.add(i, j)
			if neighbor != coord {
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}

func (board gameBoard) print() {
	// TODO: move cursor to changed cell only
	// TODO:  diff changed cell somehow?
	printable := make([]string, board.viewport)
	printerIdx := 0
	for i := board.viewAnchor.i; i < board.viewAnchor.i+board.viewport; i++ {
		row := []string{}
		for j := board.viewAnchor.j; j < board.viewAnchor.j+board.viewport; j++ {
			location := coordinate{i, j}
			val := board.aliveCells[location]
			row = append(row, val.getPrintable())
		}
		printable[printerIdx] = strings.Join(row, " ")
		printerIdx += 1
	}
	//FIXME:  this is re-printing instead of overwriting
	fmt.Print(strings.Join(printable, "\n") + "\r")
}

func makeBlack(str string) string {
	return "\u001b[30m" + str + "\u001b[39m"
}

func makeYellow(str string) string {
	return "\u001b[33m" + str + "\u001b[39m"
}
