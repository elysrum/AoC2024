package day5

import (
	"fmt"
	"io"
	"strings"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {

	var dataPart []string
	for dataLine := range challenge.Sections(input) {
		dataPart = append(dataPart, dataLine)
	}

	pageRules := make(map[string]bool)

	for _, rule := range strings.Split(dataPart[0], "\n") {
		pageRules[rule] = true
	}

	result := 0
	for _, line := range strings.Split(strings.Trim(dataPart[1], "\n"), "\n") {
		pages := strings.Split(line, ",")

		failed := false
		for i := 0; !failed && i <= len(pages)-1; i++ {
			for j := i + 1; !failed && j <= len(pages)-1; j++ {

				lookup := pages[i] + "|" + pages[j]

				if !pageRules[lookup] {
					failed = true
				}
			}
		}
		if !failed {
			result += util.MustAtoI(pages[(len(pages) / 2)])
		}
	}

	return result
}
