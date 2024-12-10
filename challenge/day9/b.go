package day9

import (
	"fmt"
	"io"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 9, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

	type Point struct {
		row int
		col int
	}

	data := challenge.Lines(input)

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
			for j := 0; j < len(antennaLocations); j++ {
				if i == j {
					continue
				}
				r1 := antennaLocations[i].row
				c1 := antennaLocations[i].col

				r2 := antennaLocations[j].row
				c2 := antennaLocations[j].col

				// Difference between antenna
				dr := r2 - r1
				dc := c2 - c1

				targetRow := r1
				targetCol := c1

				for 0 <= targetRow && targetRow < R &&
					0 <= targetCol && targetCol < C {
					antinodes[Point{row: targetRow, col: targetCol}] = true
					targetRow += dr
					targetCol += dc
				}

			}
		}
	}

	return len(antinodes)
}
