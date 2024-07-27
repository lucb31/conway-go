package logic

import (
	"fmt"
	"testing"
)

func stateEqual(expected [][]bool, received [][]bool) (bool, error) {
	for row, rowData := range expected {
		for col, _ := range rowData {
			if expected[row][col] != received[row][col] {
				return false, fmt.Errorf("Inequality at position [%d][%d]. Expected %v, but received %v", row, col, expected[row][col], received[row][col])
			}
		}
	}
	return true, nil
}

func runTestWithStateData(t *testing.T, states [][]uint8, xSize int, ySize int) {
	// Starting state
	state, err := ParseIntState(states[0], xSize, ySize)
	if err != nil {
		t.Fatal(err.Error())
	}
	// Calculate & compare states starting at epoch 1
	for epoch := 1; epoch < len(states); epoch++ {
		fmt.Printf("Epoch %d\n", epoch)
		nextState, err := Epoch(state)
		if err != nil {
			t.Fatal(err.Error())
		}
		// Compare with stored state
		expected, err := ParseIntState(states[epoch], xSize, ySize)
		if err != nil {
			t.Fatal(err.Error())
		}
		_, err = stateEqual(expected, nextState)
		if err != nil {
			PrintState(expected)
			fmt.Println("---")
			PrintState(nextState)
			t.Fatal(err.Error())
		}
		// Overwrite state to be used in next iteration
		state = nextState
	}
}

func TestBlinkerPattern(t *testing.T) {
	states := [][]uint8{
		{
			0, 1, 0,
			0, 1, 0,
			0, 1, 0,
		},
		{
			0, 0, 0,
			1, 1, 1,
			0, 0, 0,
		},
		{
			0, 1, 0,
			0, 1, 0,
			0, 1, 0,
		},
	}
	runTestWithStateData(t, states, 3, 3)
}

func TestToadPattern(t *testing.T) {
	states := [][]uint8{
		{
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 0, 0,
			0, 1, 0, 0, 1, 0,
			0, 1, 0, 0, 1, 0,
			0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
			0, 0, 1, 1, 1, 0,
			0, 1, 1, 1, 0, 0,
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 0, 0,
			0, 1, 0, 0, 1, 0,
			0, 1, 0, 0, 1, 0,
			0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
		},
	}
	runTestWithStateData(t, states, 6, 6)
}
