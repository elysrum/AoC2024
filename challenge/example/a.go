package example

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"AoC2024/challenge"
	"AoC2024/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Example Day, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.InputFile()))
		},
	}
}

func a(input io.Reader) (result int) {
	v, _ := util.First(challenge.Ints(input))

	return v
}
