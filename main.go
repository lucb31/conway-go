package main

import (
	"fmt"

	"github.com/lucb31/conway-go/logic"
)

const xSize, ySize = 3, 3

func main() {
	state, err := logic.ParseBoolState([]bool{false, true, false, false, true, false, false, true, false}, xSize, ySize)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Hello world")
	logic.PrintState(state)
	state, err = logic.Epoch(state)
	if err != nil {
		fmt.Println(err.Error())
	}
	logic.PrintState(state)
}
