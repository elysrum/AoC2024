package day13

import (
	"AoC2024/challenge"
	"AoC2024/util"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 13, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {

	data := challenge.Lines(input)

	var stones []int
	for line := range data {
		for _, stone := range strings.Split(line, " ") {

			fmt.Printf("%v\n", stone)
			stones = append(stones, util.MustAtoI(stone))

		}
	}

	for blink := 0; blink < 25; blink++ {

		idx := 0
		S := len(stones)
		for idx < S {

			stone := stones[idx]
			stringStone := strconv.Itoa(stone)

			if stone == 0 {
				stones[idx] = 1
			} else if (len(stringStone) % 2) == 0 {

				mid := len(stringStone) / 2
				left := util.MustAtoI(stringStone[0:mid])
				right := util.MustAtoI(stringStone[mid:])

				stones[idx] = left
				stones = slices.Insert(stones, idx+1, right)

				S = len(stones)

				// make sure we skip over the stone we just added
				idx += 1

			} else {
				stones[idx] = stone * 2024
			}

			idx += 1

		}

	}

	return len(stones)
}
