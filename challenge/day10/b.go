package day10

import (
	"fmt"
	"io"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 10, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

// NOTE:  STILL USING GLOBALS DEFINED IN PART A

func partB(input io.Reader) int {

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

			result += partBWalkTrail(row, col, 0, true)

		}
	}

	return result
}

// Walk from starting point to end of trail
// return true if we hit a 9
func partBWalkTrail(row, col, prevValue int, start bool) int {
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
				retVal += partBWalkTrail(row-1, col, currValue, false)
			}
			// Check East
			if (col + 1) < C {
				retVal += partBWalkTrail(row, col+1, currValue, false)
			}

			// Check South
			if (row + 1) < R {
				retVal += partBWalkTrail(row+1, col, currValue, false)
			}

			// Check West
			if (col - 1) >= 0 {
				retVal += partBWalkTrail(row, col-1, currValue, false)
			}

		}
	}
	return retVal
}
