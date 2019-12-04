package board

import (
	"strconv"
)

type coordinate struct {
	i int
	j int
}

func (c coordinate) add(i int, j int) coordinate {
	return coordinate{c.i + i, c.j + j}
}


type cell int

func (c cell) isAlive() bool {
	return c == 1
}

func (c cell) getPrintable() string {
	stringified := strconv.Itoa(int(c))
	if c == 1 {
		return makeYellow(stringified)
	}
	return makeBlack(stringified)
}

func (c cell) nextState(numAliveNeigbors int) cell {
	isAlive := c == 1

	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	if isAlive && numAliveNeigbors > 3 {
		return 0
	}

	// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	if isAlive && numAliveNeigbors < 2 {
		return 0
	}
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	if !isAlive && numAliveNeigbors == 3 {
		return 1
	}

	// Any live cell with two or three live neighbours lives on to the next generation.
	return c
}

func makeBlack(str string) string {
	return "\u001b[30m" + str + "\u001b[39m"
}

func makeYellow(str string) string {
	return "\u001b[33m" + str + "\u001b[39m"
}