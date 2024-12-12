package day12

import (
	"AoC2024/challenge"
	"fmt"
	"io"
	"slices"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 12, Problem A",
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

	data := challenge.Lines(input)

	var grid [][]rune

	for inputLine := range data {
		row := make([]rune, len(inputLine))
		for cellIndex, cell := range inputLine {
			row[cellIndex] = cell
		}
		grid = append(grid, row)
	}

	R := len(grid)
	C := len(grid[0])

	found := make(map[Point]bool)
	regions := make([][]Point, 0)

	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {

			// add current grid pos to seen list
			loc := Point{row: row, col: col}

			if found[loc] {
				continue
			}

			currValue := grid[row][col]

			workingRegion := make([]Point, 1)
			workingRegion[0] = loc
			found[loc] = true

			idx := 0
			RG := len(workingRegion)

			for idx < RG {
				for _, newPoint := range []Point{{workingRegion[idx].row - 1, workingRegion[idx].col}, {workingRegion[idx].row + 1, workingRegion[idx].col},
					{workingRegion[idx].row, workingRegion[idx].col - 1}, {workingRegion[idx].row, workingRegion[idx].col + 1}} {

					if newPoint.row < 0 || newPoint.col < 0 || newPoint.row >= R || newPoint.col >= C {
						continue
					}
					if grid[newPoint.row][newPoint.col] != currValue {
						continue
					}
					if found[newPoint] {
						continue
					}
					found[newPoint] = true
					workingRegion = append(workingRegion, newPoint)

				}
				idx += 1
				RG = len(workingRegion)
			}

			regions = append(regions, workingRegion)

		}
	}

	perim := 0
	area := 0
	retVal := 0

	for _, region := range regions {
		perim = partACalcPerimter(region)
		area = len(region)
		retVal += perim * area
	}

	return retVal
}

func partACalcPerimter(region []Point) int {

	retVal := 0

	for _, tile := range region {

		// Assume tile is isolated
		retVal += 4

		for _, newTile := range []Point{{tile.row - 1, tile.col}, {tile.row + 1, tile.col},
			{tile.row, tile.col - 1}, {tile.row, tile.col + 1}} {
			if slices.Contains(region, newTile) {
				retVal -= 1
			}

		}

	}

	return retVal
}
