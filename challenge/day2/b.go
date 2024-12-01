package day2

import (
	"fmt"
	"io"
	"strings"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

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

	result := 0
	for i := 0; i < len(list1); i++ {
		count := 0
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				count++

			}
		}
		result += list1[i] * count
	}

	return result
}
