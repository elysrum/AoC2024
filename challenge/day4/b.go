package day4

import (
	"fmt"
	"io"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

	data := challenge.Lines(input)

	result := 0

	var grid [][]rune

	for inputLine := range data {
		row := make([]rune, len(inputLine))
		for cellIndex, cell := range inputLine {
			row[cellIndex] = cell
		}
		grid = append(grid, row)
	}

	numRows := len(grid)
	numCols := len(grid[0])

	for rIndex := 0; rIndex < numRows; rIndex++ {
		for cIndex := 0; cIndex < numCols; cIndex++ {

			if cIndex+2 < numCols && rIndex+2 < numRows &&
				grid[rIndex][cIndex] == 'M' && grid[rIndex+1][cIndex+1] == 'A' && grid[rIndex+2][cIndex+2] == 'S' &&
				grid[rIndex+2][cIndex] == 'M' && grid[rIndex][cIndex+2] == 'S' {
				result += 1
			}
			if cIndex+2 < numCols && rIndex+2 < numRows &&
				grid[rIndex][cIndex] == 'M' && grid[rIndex+1][cIndex+1] == 'A' && grid[rIndex+2][cIndex+2] == 'S' &&
				grid[rIndex+2][cIndex] == 'S' && grid[rIndex][cIndex+2] == 'M' {
				result += 1
			}
			if cIndex+2 < numCols && rIndex+2 < numRows &&
				grid[rIndex][cIndex] == 'S' && grid[rIndex+1][cIndex+1] == 'A' && grid[rIndex+2][cIndex+2] == 'M' &&
				grid[rIndex+2][cIndex] == 'M' && grid[rIndex][cIndex+2] == 'S' {
				result += 1
			}
			if cIndex+2 < numCols && rIndex+2 < numRows &&
				grid[rIndex][cIndex] == 'S' && grid[rIndex+1][cIndex+1] == 'A' && grid[rIndex+2][cIndex+2] == 'M' &&
				grid[rIndex+2][cIndex] == 'S' && grid[rIndex][cIndex+2] == 'M' {
				result += 1
			}
		}
	}
	return result
}
