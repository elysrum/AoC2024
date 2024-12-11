package day10

import (
	"fmt"
	"io"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 10, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

type Point struct {
	row int
	col int
}

var grid [][]int
var found map[Point]bool

func partA(input io.Reader) int {

	data := challenge.Lines(input)

	result := 0

	for inputLine := range data {
		row := make([]int, len(inputLine))
		for cellIndex, cell := range inputLine {
			row[cellIndex] = util.MustAtoI(string(cell))
		}
		grid = append(grid, row)

	}

	Rows := len(grid)
	Cols := len(grid[0])

	// Loop until guard goes out of bounds
	for row := 0; row < Rows; row++ {
		for col := 0; col < Cols; col++ {

			if grid[row][col] != 0 {
				// Not a starting Point, next cell
				continue
			}

			_ = partAWalkTrail(row, col, 0, true)

			result += len(found)

		}
	}

	return result
}

// Walk from starting point to end of trail
// return true if we hit a 9
func partAWalkTrail(row, col, prevValue int, start bool) int {
	retVal := 0

	R := len(grid)
	C := len(grid[0])

	currValue := grid[row][col]

	if start {
		found = make(map[Point]bool)
	}

	if !start && currValue != prevValue+1 {
		retVal = 0
	} else {

		if currValue == 9 {
			loc := Point{row: row, col: col}

			if !found[loc] {
				found[loc] = true

			}
			retVal = 1
		} else {

			// Check North
			if (row - 1) >= 0 {
				retVal += partAWalkTrail(row-1, col, currValue, false)
			}
			// Check East
			if (col + 1) < C {
				retVal += partAWalkTrail(row, col+1, currValue, false)
			}

			// Check South
			if (row + 1) < R {
				retVal += partAWalkTrail(row+1, col, currValue, false)
			}

			// Check West
			if (col - 1) >= 0 {
				retVal += partAWalkTrail(row, col-1, currValue, false)
			}

		}
	}
	return retVal
}
