package main

import (
	"fmt"

	"github.com/lucb31/conway-go/logic"
)

const xSize, ySize = 10, 10

func main() {
	state := logic.InitRandomState(xSize, ySize)
	logic.PrintState(state)
	for epoch := range 10 {
		fmt.Println("---", epoch, "---")
		state, err := logic.Epoch(state)
		if err != nil {
			fmt.Println(err.Error())
		}
		logic.PrintState(state)
	}
}
