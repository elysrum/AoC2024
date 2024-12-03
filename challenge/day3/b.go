package day3

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

	data := challenge.Lines(input)

	re := regexp.MustCompile("(?:mul\\((\\d+),(\\d+)\\)|don't\\(\\)|do\\(\\))")

	result := 0
	multiply := true
	for inputLine := range data {
		matches := re.FindAllStringSubmatch(inputLine, -1)
		for _, match := range matches {

			if multiply && strings.HasPrefix(match[0], "mul(") {

				num1 := util.MustAtoI(match[1])
				num2 := util.MustAtoI(match[2])

				result += num1 * num2
			} else if strings.HasPrefix(match[0], "don't(") {
				multiply = false
			} else if strings.HasPrefix(match[0], "do(") {
				multiply = true
			}
		}
	}
	return result
}
