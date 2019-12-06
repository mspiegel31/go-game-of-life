package board

// TODO: tests :)
// TODO: smooth out drawing; redraw changes only
// TODO: memoize cell state; only recompute on change

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
			if state := rand.Intn(2); state == 1 {
				location := coordinate{i, j}
				aliveCells[location] = cell(state)
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
		for j := board.viewAnchor.j; j < board.viewAnchor.j+board.viewport; j++ {
			cell := board.getCell(coordinate{i, j})
			row = append(row, cell.getPrintable())
		}
		printable[printerIdx] = strings.Join(row, " ")
		printerIdx++
	}
	//FIXME:  this is re-printing instead of overwriting
	fmt.Print(strings.Join(printable, "\n") + "\r\n")
	fmt.Printf("# alive cells: %d", len(board.aliveCells))
}

func (board *GameBoard) getCell(location coordinate) cell {
	return board.aliveCells[location]
}

func (board *GameBoard) getNextCell(location coordinate) cell {
	numAliveNeighbors := 0
	for _, neighborLocation := range board.identifyNeighbors(location) {
		neighbor := board.getCell(neighborLocation)
		numAliveNeighbors += int(neighbor)
	}

	return board.getCell(location).nextState(numAliveNeighbors)
}

func (board *GameBoard) identifyUpdates() map[coordinate]cell {
	updates := make(map[coordinate]cell)

	for location, currentCell := range board.aliveCells {
		// save us a write
		if nextCell := board.getNextCell(location); currentCell != nextCell {
			updates[location] = nextCell
		}

		for _, neighborLocation := range board.identifyNeighbors(location) {
			candidateCell := board.getCell(neighborLocation)
			nextCell := board.getNextCell(neighborLocation)

			//check dead neighbors only;  we'll get the live ones before the loop is done
			if !candidateCell.isAlive() && nextCell != candidateCell {
				updates[neighborLocation] = nextCell
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
