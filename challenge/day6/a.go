package day6

import (
	"fmt"
	"io"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {

	type Point struct {
		row int
		col int
	}

	type Guard struct {
		direction Point
		position  Point
	}

	data := challenge.Lines(input)

	result := 0

	var grid [][]rune

	NORTH := Point{-1, 0}
	EAST := Point{0, 1}
	SOUTH := Point{1, 0}
	WEST := Point{0, -1}

	trail := make(map[Guard]bool)

	var guard = Point{}

	rowIndex := 0
	for inputLine := range data {
		row := make([]rune, len(inputLine))
		for cellIndex, cell := range inputLine {
			row[cellIndex] = cell
			// find the guard
			if cell == '^' || cell == 'v' || cell == '<' || cell == '>' {
				guard.row = rowIndex
				guard.col = cellIndex
			}
		}
		grid = append(grid, row)
		rowIndex += 1
	}

	numRows := len(grid)
	numCols := len(grid[0])

	// Loop until guard goes out of bounds
	for {
		// add current location to trail
		trail[]
	}

	return result
}
