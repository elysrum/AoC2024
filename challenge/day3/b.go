package day3

import (
	"fmt"
	"io"
	"strings"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

	data := challenge.Lines(input)

	result := 0
	for inputLine := range data {
		lineParts := strings.Split(inputLine, " ")
		fmt.Printf("%v\n", lineParts)
	}
	return result
}
