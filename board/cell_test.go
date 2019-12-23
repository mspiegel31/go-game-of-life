package board

import (
	"testing"
)

func Test_cell_isAlive(t *testing.T) {
	tests := []struct {
		name     string
		c        cell
		expected bool
	}{
		{"should return true if it is alive", 1, true},
		{"should return true if it is alive", 0, false},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if actual := testcase.c.isAlive(); actual != testcase.expected {
				t.Errorf("cell.isAlive() = %v, want %v", actual, testcase.expected)
			}
		})
	}
}

func Test_cell_nextState(t *testing.T) {
	type args struct {
		numAliveNeigbors int
	}

	type testcases struct {
		name string
		c	cell
		args args
		want cell
	}

	runTestCases := func(tests []testcases) {
		for _, testCase := range tests {
			t.Run(testCase.name, func(t *testing.T) {
				if got := testCase.c.nextState(testCase.args.numAliveNeigbors); got != testCase.want {
					t.Errorf("cell.nextState() = %v, want %v", got, testCase.want)
				}
			})
		}
	}

	t.Run("Dead Cells", func(t *testing.T) {
		tests:= []testcases {
			{"0 neighbors", 0, args{0}, 0},
			{"1 neighbors", 0, args{1}, 0},
			{"2 neighbors", 0, args{2}, 0},
			{"3 neighbors", 0, args{3}, 1},
			{"4 neighbors", 0, args{4}, 0},
			{"5 neighbors", 0, args{5}, 0},
			{"6 neighbors", 0, args{6}, 0},
			{"7 neighbors", 0, args{7}, 0},
			{"8 neighbors", 0, args{8}, 0},
		}

		runTestCases(tests)
	})
	
	t.Run("Alive Cells", func(t *testing.T) {
		tests:= []testcases{
			{"0 neighbors", 1, args{0}, 0},
			{"1 neighbors", 1, args{1}, 0},
			{"2 neighbors", 1, args{2}, 1},
			{"3 neighbors", 1, args{3}, 1},
			{"4 neighbors", 1, args{4}, 0},
			{"5 neighbors", 1, args{5}, 0},
			{"6 neighbors", 1, args{6}, 0},
			{"7 neighbors", 1, args{7}, 0},
			{"8 neighbors", 1, args{8}, 0},
		}

		runTestCases(tests)
	})
}


func Test_coordinate_add(t *testing.T) {
	existing := coordinate{0, 1}
	expected := coordinate{-1, 0}

	actual := existing.add(-1, -1)

	if expected != actual {
		t.Errorf("coordinate.add() = %v, want %v", actual, expected)
	}
}