package day8

import (
	"fmt"
	"io"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 8, Problem A",
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

	// Create Grid from input data
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

	antennas := make(map[rune][]Point)

	// Iterate over Grid, store each antenna and every location it is found at
	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			location := Point{row, col}
			ant := grid[row][col]
			if ant != '.' {
				antLocs := antennas[ant]
				antLocs = append(antLocs, location)
				antennas[ant] = antLocs
			}
		}
	}

	antinodes := make(map[Point]bool)

	// iterate over antennas and then check each point against each other point
	for _, antennaLocations := range antennas {

		for i := 0; i < len(antennaLocations); i++ {
			for j := i + 1; j < len(antennaLocations); j++ {

				r1 := antennaLocations[i].row
				c1 := antennaLocations[i].col

				r2 := antennaLocations[j].row
				c2 := antennaLocations[j].col

				r3 := (2 * r1) - r2
				c3 := (2 * c1) - c2

				r4 := (2 * r2) - r1
				c4 := (2 * c2) - c1

				antinodes[Point{row: r3, col: c3}] = true
				antinodes[Point{row: r4, col: c4}] = true

			}
		}
	}

	// iterate over all antinodes and only count those in bounds
	for point, _ := range antinodes {
		if 0 <= point.row && point.row < R &&
			0 <= point.col && point.col < C {
			result += 1
		}
	}

	return result
}
