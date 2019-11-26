package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	rawArgs := os.Args[1:]
	if len(rawArgs) > 3 {
		panic(fmt.Sprintf("passed %d;  require at most 3 arguments", len(rawArgs)))
	}

	runSimulation(parseArgs(rawArgs))
}

func parseArgs(rawArgs []string) simulationParams {
	args := make([]int, len(rawArgs))

	for i := 0; i < len(rawArgs); i++ {
		arg, err := strconv.Atoi(rawArgs[i])
		if err != nil {
			panic(err)
		}
		args[i] = arg
	}

	//TODO: handle defaults
	return NewSimulationParams(args...)
}

