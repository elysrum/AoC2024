package day4

import (
	"fmt"
	"io"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {

	data := challenge.Lines(input)

	result := 0

	for inputLine := range data {
		fmt.Printf("%v\n", inputLine)

	}
	return result
}
