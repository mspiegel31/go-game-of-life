package board

// TODO: refactor for better code organization
// export GameBoard et al.
// break cells and boards into separate modules?

// TODO: tests :)
// TODO: smooth out drawing; redraw changes only
// TODO: memoize cell state; only recompute on change
// TODO:completely re-implement with sparse arrays :)

import (
	"fmt"
	"math/rand"
	"strings"
)

type GameBoard struct {
	viewAnchor coordinate
	viewport   int
	aliveCells map[coordinate]cell
}

func Init(size int) GameBoard {
	anchor := coordinate{50, 50}
	aliveCells := make(map[coordinate]cell)
	board := GameBoard{anchor, size, aliveCells}

	// seed values in viewport randomly
	for i := anchor.i; i < anchor.i+size; i++ {
		for j := anchor.j; j < anchor.j+size; j++ {
			data := rand.Intn(2)
			if data == 1 {
				location := coordinate{i, j}
				aliveCells[location] = cell(data)
			}
		}
	}
	return board

}

func (board *GameBoard) NextState() {
	updates := board.identifyUpdates()
	board.applyUpdates(updates)
}

func (board *GameBoard) Print() {
	// TODO: move cursor to changed cell only
	// TODO:  diff changed cell somehow?
	printable := make([]string, board.viewport)
	printerIdx := 0
	for i := board.viewAnchor.i; i < board.viewAnchor.i+board.viewport; i++ {
		row := []string{}
		for j := board.viewAnchor.j; j < board.viewAnchor.j + board.viewport; j++ {
			val := board.getCell(coordinate{i, j})
			row = append(row, val.getPrintable())
		}
		printable[printerIdx] = strings.Join(row, " ")
		printerIdx++
	}
	//FIXME:  this is re-printing instead of overwriting
	fmt.Print(strings.Join(printable, "\n") + "\r")
}

func (board *GameBoard) getCell(location coordinate) cell {
	return board.aliveCells[location]
}

func (board *GameBoard) getNextCell(location coordinate) cell {
	neighbors := board.identifyNeighbors(location)
	numAliveNeighbors := 0
	for _, neighborLocation := range neighbors {
		neighbor := board.getCell(neighborLocation)
		numAliveNeighbors += int(neighbor)
	}

	return board.getCell(location).nextState(numAliveNeighbors)
}

func (board *GameBoard) identifyUpdates() map[coordinate]cell {
	updates := make(map[coordinate]cell)
	
	for location, currentCell := range board.aliveCells {
		nextCell := board.getNextCell(location)
		if nextCell != currentCell {
			updates[location] = nextCell
		}
		
		//FIXME: something about this logic is causing the updates to remain the same after a few iterations
		neighbors := board.identifyNeighbors(location)
		for _, neighborLocation := range neighbors {
			candidateCell := board.getCell(neighborLocation)
			if !candidateCell.isAlive() {
				nextCell := board.getNextCell(neighborLocation)
				if nextCell != candidateCell {
					updates[location] = nextCell
				}
			}
		}
	}
	return updates
}

func (board *GameBoard) applyUpdates(updates map[coordinate]cell) {
	for location, nextCell := range updates {
		if !nextCell.isAlive() {
			delete(board.aliveCells, location)
		} else {
			board.aliveCells[location] = nextCell
		}
	}
}

func (board *GameBoard) identifyNeighbors(coord coordinate) []coordinate {
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
