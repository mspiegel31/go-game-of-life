package board

import (
	"reflect"
	"testing"
)

func Benchmark_nextState_computationSpeed(b *testing.B) {
	board := Init(500)
	for index := 0; index < b.N; index++ {
		board.NextState()
	}
}
func TestGameBoard_identifyUpdates(t *testing.T) {
	tests := []struct {
		name      string
		boardSeed GameBoard
		want      map[coordinate]cell
	}{
		{"no updates", toGameBoard(staticBoard), make(map[coordinate]cell)},
		{"blinker board", toGameBoard(blinkerBoard), map[coordinate]cell{
			coordinate{1, 2}: 0,
			coordinate{3, 2}: 0,
			coordinate{2, 1}: 1,
			coordinate{2, 3}: 1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.boardSeed.identifyUpdates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameBoard.identifyUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBoard_applyUpdates(t *testing.T) {
	board := toGameBoard(blinkerBoard)

	updates := map[coordinate]cell {
		coordinate{0, 4}: 1,
		coordinate{4, 4}: 1,
		coordinate{1, 2}: 0,
		coordinate{2, 2}: 0,
		coordinate{3, 2}: 0,
	}

	expected := map[coordinate]cell {
		coordinate{0, 4}: 1,
		coordinate{4, 4}: 1,
	}

	board.applyUpdates(updates)

	if actual := board.aliveCells; !reflect.DeepEqual(expected, actual) {
		t.Errorf("GameBoard.identifyUpdates() = %v, want %v", actual, expected)

	}
}

func TestGameBoard_nextBoard(t *testing.T) {
	board := toGameBoard(blinkerBoard)
	expectedNextBoard := toGameBoard([][]cell {
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	})

	board.NextState()

	if !reflect.DeepEqual(board, expectedNextBoard) {
		t.Errorf("GameBoard.NextBoard() = %v, want %v", board, expectedNextBoard)
	}

}

func toGameBoard(arr [][]cell) GameBoard {
	viewAnchor := coordinate{0, 0}
	viewport := len(arr)
	aliveCells := make(map[coordinate]cell)

	for i, row := range arr {
		for j, c := range row {
			if c.isAlive() {
				aliveCells[coordinate{i, j}] = c
			}
		}
	}
	return GameBoard{viewAnchor, viewport, aliveCells}
}

var blinkerBoard [][]cell = [][]cell{
	{0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0},
}

var staticBoard [][]cell = [][]cell{
	{0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 1, 0, 1, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0},
}
