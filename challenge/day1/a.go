package day1

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {

	data := challenge.Lines(input)

	var list1, list2 []int
	for inputLine := range data {
		linePart1, linePart2, found := strings.Cut(inputLine, " ")

		if found {
			//fmt.Printf("%v - %v \n", linePart1, linePart2)
			//	fmt.Printf("%v \n", strings.Split(dataSet[i], " ")[1])

			list1 = append(list1, util.MustAtoI(strings.TrimSpace(linePart1)))
			list2 = append(list2, util.MustAtoI(strings.TrimSpace(linePart2)))
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	result := 0
	for i := 0; i < len(list1); i++ {
		val := list1[i] - list2[i]
		if val < 0 {
			val = -val
		}
		result += val
	}

	return result
}
