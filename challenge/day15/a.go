package day15

import (
	"AoC2024/challenge"
	"AoC2024/util"
	"fmt"
	"io"

	"regexp"

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

func partA(input io.Reader) int {

	re, err := regexp.Compile("-?\\d+")
	if err != nil {
		panic(err)
	}

	R := 103
	C := 101
	// R = 7
	// C = 11

	S := 100

	grid := make([][]int, R)

	for idx := range grid {
		grid[idx] = make([]int, C)
	}

	data := challenge.Lines(input)

	for line := range data {

		items := re.FindAllString(line, -1)

		x := util.MustAtoI(items[0])
		y := util.MustAtoI(items[1])
		vx := util.MustAtoI(items[2])
		vy := util.MustAtoI(items[3])

		newX := (x + vx*S) % C
		newY := (y + vy*S) % R

		if newX < 0 {
			newX = C + newX
		}
		if newY < 0 {
			newY = R + newY
		}

		grid[newY][newX] += 1
	}

	midx := C / 2
	midy := R / 2

	quad1 := 0
	quad2 := 0
	quad3 := 0
	quad4 := 0

	for x := 0; x < C; x++ {
		for y := 0; y < R; y++ {
			if grid[y][x] > 0 {

				if x < midx {
					if y < midy {
						quad1 += grid[y][x]
					}
					if y > midy {
						quad2 += grid[y][x]
					}
				}
				if x > midx {
					if y < midy {
						quad3 += grid[y][x]
					}
					if y > midy {
						quad4 += grid[y][x]
					}

				}
			}
		}
	}

	return quad1 * quad2 * quad3 * quad4
}
