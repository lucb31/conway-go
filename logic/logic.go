package logic

import (
	"errors"
)

var InvalidStateDimensions = errors.New("Invalid state dimensions")

type State struct {
	xSize int
	ySize int
	Vals  [][]bool
}

func (s *State) calcNeighbors(row int, col int) (int, error) {
	neighbors := 0
	for r := max(0, row-1); r < min(s.ySize, row+2); r++ {
		for c := max(0, col-1); c < min(s.xSize, col+2); c++ {
			// Ignore myself
			if r == row && c == col {
				continue
			}
			if s.Vals[r][c] {
				neighbors++
			}
		}
	}
	return neighbors, nil
}

func calcNextState(alive bool, neighbors int) bool {
	if alive {
		if neighbors < 2 {
			return false
		} else if neighbors > 3 {
			return false
		}
	} else if neighbors == 3 {
		return true
	}
	// Do nothing
	return alive
}

func (s *State) Epoch() error {
	nextState := InitEmptyState(s.xSize, s.ySize)
	for row := 0; row < s.ySize; row++ {
		for col := 0; col < s.xSize; col++ {
			// Check neighbors
			neighbors, err := s.calcNeighbors(row, col)
			if err != nil {
				return err
			}
			// Update state
			nextState[row][col] = calcNextState(s.Vals[row][col], neighbors)
		}
	}
	s.Vals = nextState
	return nil
}
