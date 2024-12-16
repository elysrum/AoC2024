package day15

import (
	"AoC2024/challenge"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 15, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

type Point struct {
	row int
	col int
}

func partA(input io.Reader) int {

	blocks := make([]string, 0)
	data := challenge.Sections(input)

	for line := range data {
		blocks = append(blocks, line)

	}

	// result := 0

	var grid [][]rune

	var robot = Point{}

	inputLine := strings.Split(blocks[0], "\n")
	for rowIndex, line := range inputLine {
		row := make([]rune, len(inputLine))

		for cellIndex, cell := range line {
			if cell == '@' {
				robot.row = rowIndex
				robot.col = cellIndex
				row[cellIndex] = '.'

			} else {
				row[cellIndex] = cell
			}
		}
		grid = append(grid, row)
	}

	steps := blocks[1]
	NORTH := Point{-1, 0}
	EAST := Point{0, 1}
	SOUTH := Point{1, 0}
	WEST := Point{0, -1}

	var direction Point

	for _, step := range steps {
		switch step {
		case '^':
			direction = NORTH
		case '>':
			direction = EAST
		case 'v':
			direction = SOUTH
		case '<':
			direction = WEST
		case '\n':
			continue
		}

		newRow := robot.row + direction.row
		newCol := robot.col + direction.col

		// We can move freely
		if grid[newRow][newCol] == '.' {
			robot.row = newRow
			robot.col = newCol
			continue
		}
		// We've hit a wall, nothing doing here.
		if grid[newRow][newCol] == '#' {
			continue
		}

		// We've hit a barrel, push it along
		if grid[newRow][newCol] == 'O' {

			foundSpace := false
			moreGrid := true
			testRow := newRow
			testCol := newCol
			for moreGrid && !foundSpace {
				testRow += direction.row
				testCol += direction.col

				// We've hit a wall, stop looking
				if grid[testRow][testCol] == '#' {
					moreGrid = false
				}
				// found space
				if grid[testRow][testCol] == '.' {
					grid[testRow][testCol] = 'O'
					grid[newRow][newCol] = '.'
					robot.row = newRow
					robot.col = newCol
					foundSpace = true
				}

			}

		}

	}

	result := 0

	for rindx, row := range grid {
		for cindx, cell := range row {
			if cell == 'O' {
				result += 100*rindx + cindx

			}

		}
	}

	return result
}

func printGrid(grid [][]rune, robot Point) int {

	count := 0
	for ri, row := range grid {
		for ci, cell := range row {

			if cell == 'O' {
				count++
			}
			if ri == robot.row && ci == robot.col {
				fmt.Printf("@")
			} else {
				fmt.Printf("%v", string(cell))
			}
		}
		fmt.Printf("\n")
	}
	return count
}
