# Conway's Game of Life

[![asciicast](https://asciinema.org/a/318500.svg)](https://asciinema.org/a/318500)

## Rules
(from [wikipedia](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules))
The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in one of two possible states, alive or dead, (or populated and unpopulated, respectively). Every cell interacts with its eight neighbours, which are the cells that are horizontally, vertically, or diagonally adjacent. At each step in time, the following transitions occur:

    Any live cell with fewer than two live neighbors dies, as if by underpopulation.
    Any live cell with two or three live neighbors lives on to the next generation.
    Any live cell with more than three live neighbors dies, as if by overpopulation.
    Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.


## Motivations
This is a small implementation of Conway's Game of Life, primarily for the purposes of familiarizing myself with the Go language.  It's also a rather interesting problem for trying out different approaches for optimization, code organization and testing.

## Usage
### Build from source
1. `go build`
1. `./go-game-of-life` will run a simulation with the defaults:
    1. 500 iterations
    1. board size of 30X30
    1. render delay of 70 ms
1. you can optionally tweak test params by passing them at the command line, e.g.
```shell
./go-game-of-life <numIterations> <boardSize> <renderDelay>
```

### Tests
a minimal test suite is provided, to use it run `go test ./...` in the current directory