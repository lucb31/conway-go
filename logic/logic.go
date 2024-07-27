package logic

import (
	"errors"
	"fmt"
)

var InvalidStateDimensions = errors.New("Invalid state dimensions")

func calcNeighbors(state [][]bool, row int, col int) (int, error) {
	if len(state) < 1 {
		return 0, InvalidStateDimensions
	}
	xSize, ySize := len(state), len(state[0])
	neighbors := 0
	for r := max(0, row-1); r < min(xSize, row+2); r++ {
		for c := max(0, col-1); c < min(ySize, col+2); c++ {
			// Ignore myself
			//fmt.Println(row+1, col+1, xSize, ySize, max(0, row-1), min(xSize-1, row+1))
			fmt.Printf("Checking [%d][%d] [%v] \n", r, c, state[r][c])
			if r == row && c == col {
				continue
			}
			if state[r][c] {
				neighbors++
			}
		}
	}
	return neighbors, nil
}

func calcNextState(alive bool, neighbors int) bool {
	if alive {
		if neighbors < 2 {
			fmt.Println("Cell dies by underpopulation")
			return false
		}
		if neighbors > 3 {
			fmt.Println("Cell dies by overpopulation")
			return false
		}
	} else if neighbors == 3 {
		fmt.Println("Cell spawned by reproduction")
		return true
	}
	// Do nothing
	return alive
}

func Epoch(state [][]bool) ([][]bool, error) {
	xSize, ySize := len(state), len(state[0])
	nextState := InitState(xSize, ySize)
	for row := range xSize {
		for col := range ySize {
			// Check neighbors
			neighbors, err := calcNeighbors(state, row, col)
			if err != nil {
				return [][]bool{}, err
			}
			fmt.Printf("Cell [%d][%d] has %d neighbors\n", row, col, neighbors)
			// Update state
			nextState[row][col] = calcNextState(state[row][col], neighbors)
			fmt.Printf("----\n")
		}
	}
	return nextState, nil
}
