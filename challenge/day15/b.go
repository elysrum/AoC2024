package day15

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 13, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

	blocks := make([]string, 0)
	data := challenge.Sections(input)

	for line := range data {
		blocks = append(blocks, line)

	}

	var grid [][]rune

	var robot = Point{}

	inputLine := strings.Split(blocks[0], "\n")
	for rowIndex, line := range inputLine {
		row := make([]rune, len(inputLine)*2)

		for cellIndex, srcIndex := 0, 0; cellIndex < len(row) && srcIndex < len(line); srcIndex, cellIndex = srcIndex+1, cellIndex+2 {
			if line[srcIndex] == '@' {
				robot.row = rowIndex
				robot.col = cellIndex
				row[cellIndex] = '.'
				row[cellIndex+1] = '.'

			} else if line[srcIndex] == '#' {
				row[cellIndex] = '#'
				row[cellIndex+1] = '#'
			} else if line[srcIndex] == '.' {
				row[cellIndex] = '.'
				row[cellIndex+1] = '.'
			} else if line[srcIndex] == 'O' {
				row[cellIndex] = '['
				row[cellIndex+1] = ']'

			}
		}

		grid = append(grid, row)
	}

	B_printGrid(grid, robot)

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

		candidates := make([]Point, 1)
		candidates[0].row = robot.row
		candidates[0].col = robot.col
		candidateCount := 1
		foundSpace := true

		for idx := 0; idx < candidateCount; idx++ {

			newRow := candidates[idx].row + direction.row
			newCol := candidates[idx].col + direction.col

			cell := grid[newRow][newCol]

			if slices.Contains(candidates, Point{row: newRow, col: newCol}) {
				continue
			}
			// We've hit a wall, nothing doing here.
			if cell == '#' {
				foundSpace = false
				break
			}

			// We've hit a barrel, add it to list to push later
			if cell == '[' {

				candidates = append(candidates, []Point{{row: newRow, col: newCol}, {row: newRow, col: newCol + 1}}...)
				candidateCount += 2

			}
			if cell == ']' {
				candidates = append(candidates, []Point{{row: newRow, col: newCol}, {row: newRow, col: newCol - 1}}...)
				candidateCount += 2
			}
		}

		if !foundSpace {
			continue
		}
		gridCopy := make([][]rune, 0)
		for _, r := range grid {
			row := make([]rune, 0)
			for _, c := range r {
				row = append(row, c)
			}
			gridCopy = append(gridCopy, row)

		}

		grid[robot.row][robot.col] = '.'

		for _, point := range candidates[1:] {
			grid[point.row][point.col] = '.'
		}

		for _, point := range candidates[1:] {
			grid[point.row+direction.row][point.col+direction.col] = gridCopy[point.row][point.col]
		}

		robot.row += direction.row
		robot.col += direction.col

	}

	result := 0

	for rindx, row := range grid {
		for cindx, cell := range row {
			if cell == '[' {
				result += 100*rindx + cindx

			}

		}
	}

	return result
}
func B_printGrid(grid [][]rune, robot Point) int {

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
