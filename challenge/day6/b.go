package day6

import (
	"fmt"
	"io"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	type Point struct {
		row       int
		col       int
		direction int
	}

	type Seen struct {
		row       int
		col       int
		direction int
	}

	data := challenge.Lines(input)

	result := 0

	var grid [][]rune

	NORTH := Point{-1, 0, 0}
	EAST := Point{0, 1, 1}
	SOUTH := Point{1, 0, 2}
	WEST := Point{0, -1, 3}

	var guardPos = Seen{}
	var guardStart = Seen{}

	startDirection := NORTH

	rowIndex := 0
	for inputLine := range data {
		row := make([]rune, len(inputLine))
		for cellIndex, cell := range inputLine {
			row[cellIndex] = cell
			// find the guard
			if cell == '^' || cell == 'v' || cell == '<' || cell == '>' {
				switch cell {
				case '^':
					startDirection = NORTH
				case '>':
					startDirection = EAST
				case 'v':
					startDirection = SOUTH
				case '<':
					startDirection = WEST
				}
				guardStart.row = rowIndex
				guardStart.col = cellIndex
				guardStart.direction = startDirection.direction
			}
		}
		grid = append(grid, row)
		rowIndex += 1
	}

	numRows := len(grid)
	numCols := len(grid[0])

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {

			guardPos.row = guardStart.row
			guardPos.col = guardStart.col
			guardPos.direction = guardStart.direction

			direction := startDirection

			oldCell := grid[row][col]
			grid[row][col] = '#'

			trail := make(map[Seen]bool)

			// Loop until guard goes out of bounds
		infiniteLoop:
			for {
				// add current location to trail
				if trail[guardPos] {
					// We've been here before in this direction, therefore we are looping
					result += 1
					grid[row][col] = oldCell

					break infiniteLoop
				}
				trail[guardPos] = true

				// Is next step out of bounds?  If so finish
				if guardPos.col+direction.col == numCols || guardPos.col+direction.col < 0 ||
					guardPos.row+direction.row == numRows || guardPos.row+direction.row < 0 {

					grid[row][col] = oldCell

					break infiniteLoop
				}
				for grid[guardPos.row+direction.row][guardPos.col+direction.col] == '#' {
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
				guardPos.row += direction.row
				guardPos.col += direction.col
				guardPos.direction = direction.direction
			}

		}
	}

	return result
}
