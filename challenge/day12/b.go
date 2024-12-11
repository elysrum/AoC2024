package day12

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"AoC2024/challenge"
	"AoC2024/util"

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

type Store struct {
	stone int
	steps int
}

var cache map[Store]int

func partB(input io.Reader) int {

	data := challenge.Lines(input)

	var stones []int
	for line := range data {
		for _, stone := range strings.Split(line, " ") {

			stones = append(stones, util.MustAtoI(stone))

		}
	}

	cache = make(map[Store]int)

	result := 0
	for _, stone := range stones {
		result += processStones(stone, 75)
	}

	return result
}

func processStones(stone int, steps int) int {

	if steps == 0 {
		return 1
	}

	result, ok := cache[Store{stone: stone, steps: steps}]

	if ok {
		return result
	}

	stringStone := strconv.Itoa(stone)
	result = 0

	if stone == 0 {
		result = processStones(1, steps-1)
	} else if (len(stringStone) % 2) == 0 {

		mid := len(stringStone) / 2
		left := util.MustAtoI(stringStone[0:mid])
		right := util.MustAtoI(stringStone[mid:])

		result = processStones(left, steps-1) + processStones(right, steps-1)
	} else {
		result = processStones(stone*2024, steps-1)
	}

	cache[Store{stone: stone, steps: steps}] = result

	return result

}
