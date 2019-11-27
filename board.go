package main

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

func (b board) inBounds(coord coordinate) bool {
	onBoard := func(p int) bool { return p >= 0 && p < b.size }
	return onBoard(coord.i) && onBoard(coord.j)
}

type cell struct {
	data      int
	location  coordinate
	neighbors []coordinate
}

func (c cell) getPrintable() string {
	stringified := strconv.Itoa(c.data)
	if c.data == 1 {
		return makeYellow(stringified)
	}
	return makeBlack(stringified)
}

type board struct {
	size     int
	viewport int
	state    [][]cell
}

func (b board) identifyNeighbors(coord coordinate) []coordinate {
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


func (b board) print() {
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

func makeBlack(str string) string {
	return "\u001b[30m" + str + "\u001b[39m"
}

func makeYellow(str string) string {
	return "\u001b[33m" + str + "\u001b[39m"
}