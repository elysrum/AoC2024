package day14

import (
	"fmt"
	"io"

	"regexp"

	"AoC2024/challenge"
	"AoC2024/util"

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

	type Robot struct {
		x  int
		y  int
		vx int
		vy int
	}

	re, err := regexp.Compile("-?\\d+")
	if err != nil {
		panic(err)
	}

	R := 103
	C := 101

	S := R * C

	robots := make([]Robot, 0)
	data := challenge.Lines(input)

	for line := range data {

		items := re.FindAllString(line, -1)

		robots = append(robots, Robot{
			x:  util.MustAtoI(items[0]),
			y:  util.MustAtoI(items[1]),
			vx: util.MustAtoI(items[2]),
			vy: util.MustAtoI(items[3]),
		})
	}

	for s := 0; s < S; s++ {

		grid := make([][]int, R)

		for idx := range grid {
			grid[idx] = make([]int, C)
		}

		for _, robot := range robots {
			newX := (robot.x + robot.vx*s) % C
			newY := (robot.y + robot.vy*s) % R

			if newX < 0 {
				newX = C + newX
			}
			if newY < 0 {
				newY = R + newY
			}

			grid[newY][newX] += 1
		}

		fmt.Printf("%v\n", s)
		for x := 0; x < C; x++ {
			for y := 0; y < R; y++ {
				if grid[y][x] == 0 {
					fmt.Printf(".")
				} else {
					fmt.Printf("#")
				}

			}
			fmt.Printf("\n")
		}
	}

	// Capture output and examine with mark 1 eye ball
	return 0
}
