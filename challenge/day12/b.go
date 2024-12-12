package day12

import (
	"fmt"
	"io"
	"slices"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 12, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

type Corner struct {
	row float32
	col float32
}

func partB(input io.Reader) int {

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
		perim = partBCalcSides(region)
		area = len(region)
		retVal += perim * area
	}

	return retVal
}

func partBCalcSides(region []Point) int {

	allCorners := make([]Corner, 0)

	for _, tile := range region {

		for _, corner := range []Corner{{float32(tile.row) - 0.5, float32(tile.col) - 0.5}, {float32(tile.row) + 0.5, float32(tile.col) - 0.5},
			{float32(tile.row) + 0.5, float32(tile.col) + 0.5}, {float32(tile.row) - 0.5, float32(tile.col) + 0.5}} {
			if slices.Contains(allCorners, corner) {
				continue
			} else {
				allCorners = append(allCorners, corner)
			}
		}
	}
	actualCorners := 0

	for _, corner := range allCorners {

		validCorners := make([]bool, 4)

		for idx, newTile := range []Point{{int(corner.row - 0.5), int(corner.col - 0.5)}, {int(corner.row + 0.5), int(corner.col - 0.5)},
			{int(corner.row + 0.5), int(corner.col + 0.5)}, {int(corner.row - 0.5), int(corner.col + 0.5)}} {
			validCorners[idx] = slices.Contains(region, newTile)

		}
		number := 0
		for _, test := range validCorners {
			if test {
				number += 1
			}
		}

		if number == 1 {
			actualCorners += 1
		} else if number == 2 {
			if (validCorners[0] && !validCorners[1] && validCorners[2] && !validCorners[3]) || (!validCorners[0] && validCorners[1] && !validCorners[2] && validCorners[3]) {
				actualCorners += 2
			}
		} else if number == 3 {
			actualCorners += 1
		}
	}

	return actualCorners
}
