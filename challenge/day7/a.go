package day7

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

	data := challenge.Lines(input)

	result := 0

	var grid [][]rune

	NORTH := Point{-1, 0}
	EAST := Point{0, 1}
	SOUTH := Point{1, 0}
	WEST := Point{0, -1}

	trail := make(map[Point]int)

	var guard = Point{}

	direction := NORTH

	rowIndex := 0
	for inputLine := range data {
		row := make([]rune, len(inputLine))
		for cellIndex, cell := range inputLine {
			row[cellIndex] = cell
			// find the guard
			if cell == '^' || cell == 'v' || cell == '<' || cell == '>' {
				switch cell {
				case '^':
					direction = NORTH
				case '>':
					direction = EAST
				case 'v':
					direction = SOUTH
				case '<':
					direction = WEST
				}
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
		trail[guard] += 1

		// Is next step out of bounds?  If so finish
		if guard.col+direction.col == numCols || guard.col+direction.col < 0 ||
			guard.row+direction.row == numRows || guard.row+direction.row < 0 {
			break
		}
		// Is next step an obstacle, if so, change direction
		if grid[guard.row+direction.row][guard.col+direction.col] == '#' {
			switch {
			case direction == NORTH:
				direction = EAST
			case direction == EAST:
				direction = SOUTH
			case direction == SOUTH:
				direction = WEST
			case direction == WEST:
				direction = NORTH
			}
		}

		// Make next step
		guard.row += direction.row
		guard.col += direction.col
	}

	result = len(trail)

	return result
}
