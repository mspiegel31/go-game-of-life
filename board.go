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

func (c coordinate) inBounds(boardSize int) bool {
	onBoard := func(p int) bool { return p >= 0 && p < boardSize }
	return onBoard(c.i) && onBoard(c.j)
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
	return "\u001b[33m" + str + "\u001b[39m";
  }