package day4

import (
	"fmt"
	"io"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {

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

			if numCols-3 > cIndex &&
				grid[rIndex][cIndex] == 'X' &&
				grid[rIndex][cIndex+1] == 'M' &&
				grid[rIndex][cIndex+2] == 'A' &&
				grid[rIndex][cIndex+3] == 'S' {
				result += 1
			}
			if numCols-3 > cIndex &&
				grid[rIndex][cIndex] == 'S' &&
				grid[rIndex][cIndex+1] == 'A' &&
				grid[rIndex][cIndex+2] == 'M' &&
				grid[rIndex][cIndex+3] == 'X' {
				result += 1
			}
			if numRows-3 > rIndex &&
				grid[rIndex][cIndex] == 'X' &&
				grid[rIndex+1][cIndex] == 'M' &&
				grid[rIndex+2][cIndex] == 'A' &&
				grid[rIndex+3][cIndex] == 'S' {
				result += 1
			}
			if numRows-3 > rIndex &&
				grid[rIndex][cIndex] == 'S' &&
				grid[rIndex+1][cIndex] == 'A' &&
				grid[rIndex+2][cIndex] == 'M' &&
				grid[rIndex+3][cIndex] == 'X' {
				result += 1
			}
			if numRows-3 > rIndex && numCols-3 > cIndex &&
				grid[rIndex][cIndex] == 'X' &&
				grid[rIndex+1][cIndex+1] == 'M' &&
				grid[rIndex+2][cIndex+2] == 'A' &&
				grid[rIndex+3][cIndex+3] == 'S' {
				result += 1
			}
			if numRows-3 > rIndex && numCols-3 > cIndex &&
				grid[rIndex][cIndex] == 'S' &&
				grid[rIndex+1][cIndex+1] == 'A' &&
				grid[rIndex+2][cIndex+2] == 'M' &&
				grid[rIndex+3][cIndex+3] == 'X' {
				result += 1
			}
			if numRows-3 > rIndex && cIndex >= 3 &&
				grid[rIndex][cIndex] == 'X' &&
				grid[rIndex+1][cIndex-1] == 'M' &&
				grid[rIndex+2][cIndex-2] == 'A' &&
				grid[rIndex+3][cIndex-3] == 'S' {
				result += 1
			}
			if numRows-3 > rIndex && cIndex >= 3 &&
				grid[rIndex][cIndex] == 'S' &&
				grid[rIndex+1][cIndex-1] == 'A' &&
				grid[rIndex+2][cIndex-2] == 'M' &&
				grid[rIndex+3][cIndex-3] == 'X' {
				result += 1
			}

		}
	}
	return result
}
