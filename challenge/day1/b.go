package day1

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

// func partB(input io.Reader) int {

// 	data := challenge.Lines(input)

// 	var list1, list2 []int

// 	for inputLine := range data {
// 		linePart1, linePart2, found := strings.Cut(inputLine, " ")

// 		if found {
// 			//fmt.Printf("%v - %v \n", linePart1, linePart2)
// 			//	fmt.Printf("%v \n", strings.Split(dataSet[i], " ")[1])

// 			list1 = append(list1, util.MustAtoI(strings.TrimSpace(linePart1)))
// 			list2 = append(list2, util.MustAtoI(strings.TrimSpace(linePart2)))
// 		}
// 	}

// 	result := 0
// 	for i := 0; i < len(list1); i++ {
// 		count := 0
// 		for j := 0; j < len(list2); j++ {
// 			if list1[i] == list2[j] {
// 				count++

// 			}
// 		}
// 		result += list1[i] * count
// 	}

// 	return result
// }

// Using a Map to pre-count occurrences of list 2 is much quicker 1.284326ms vs 472.943µs
func partB(input io.Reader) int {

	data := challenge.Lines(input)

	var list1 []int

	countMap := make(map[int]int)

	for inputLine := range data {
		linePart1, linePart2, found := strings.Cut(inputLine, " ")

		if found {
			val1 := util.MustAtoI(strings.TrimSpace(linePart1))
			val2 := util.MustAtoI(strings.TrimSpace(linePart2))
			list1 = append(list1, val1)
			countMap[val2] = countMap[val2] + 1
		}
	}

	result := 0
	for i := 0; i < len(list1); i++ {
		count := countMap[list1[i]]
		result += list1[i] * count
	}

	return result
}
