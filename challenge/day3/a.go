package day3

import (
	"fmt"
	"io"
	"regexp"

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

	re := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	result := 0
	for inputLine := range data {
		matches := re.FindAllStringSubmatch(inputLine, -1)
		for _, match := range matches {
			num1 := util.MustAtoI(match[1])
			num2 := util.MustAtoI(match[2])

			result += num1 * num2
		}
	}
	return result
}
