package main

import (
	"os"
	"raderh2o/brainfuck_interpreter/executor"
	"raderh2o/brainfuck_interpreter/parser"
)

func main() {
    data, err := os.ReadFile("app.bf")
    if err != nil {
        panic(err)
    }
    cells := make([]uint8, 1)
    current := 0

    parsedBf := parser.ParseBf(string(data))
    executor.ExecuteBf(parsedBf, &current, cells)
}
