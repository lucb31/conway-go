package logic

import (
	"fmt"
	"math/rand"
)

func ParseBoolState(flatState []bool, xSize int, ySize int) ([][]bool, error) {
	if len(flatState) != xSize*ySize {
		return [][]bool{}, fmt.Errorf("Invalid state dimensions. Expected %d, got %d", xSize*ySize, len(flatState))
	}
	picture := make([][]bool, ySize) // One row per unit of y.
	count := 0
	for row := range ySize {
		picture[row] = make([]bool, xSize)
		for col := range xSize {
			picture[row][col] = flatState[count]
			count++
		}
	}
	return picture, nil
}

func ParseIntState(flatState []uint8, xSize int, ySize int) ([][]bool, error) {
	boolState := make([]bool, xSize*ySize)
	for idx := range flatState {
		boolState[idx] = flatState[idx] == 1
	}
	return ParseBoolState(boolState, xSize, ySize)
}

func InitEmptyState(xSize int, ySize int) [][]bool {
	picture := make([][]bool, ySize) // One row per unit of y.
	// Allocate one large slice to hold all the pixels.
	pixels := make([]bool, xSize*ySize)
	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for i := range picture {
		picture[i], pixels = pixels[:xSize], pixels[xSize:]
	}
	return picture
}

func InitRandomState(xSize int, ySize int) [][]bool {
	picture := make([][]bool, ySize) // One row per unit of y.
	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for row := range ySize {
		picture[row] = make([]bool, xSize)
		for col := range ySize {
			picture[row][col] = rand.Intn(2) == 0
		}
	}
	return picture
}

func PrintState(state [][]bool) {
	for _, rowData := range state {
		for _, colData := range rowData {
			if colData {
				fmt.Print("X ")
			} else {
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}
}
