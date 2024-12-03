package day2

import (
	"fmt"
	"io"
	"strings"

	"AoC2024/challenge"
	"AoC2024/util"
	"AoC2024/util/gmath"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {

	data := challenge.Lines(input)

	countSafe := 0
	val, oldVal := 0, 0
	for inputLine := range data {
		lineParts := strings.Split(inputLine, " ")

		safe := true
		first := true
		increasing := 0
		for _, element := range lineParts {
			val = util.MustAtoI(element)

			diff := gmath.Abs(oldVal - val)
			if first {
				first = false
				oldVal = val
				continue
			} else if increasing == 0 {
				if (oldVal - val) > 0 {
					increasing = 1
				} else {
					increasing = -1
				}
			}

			if safe &&
				(diff >= 1 && diff <= 3) &&
				((increasing == 1 && (oldVal-val) > 0) || (increasing == -1 && (oldVal-val) < 0)) {
				safe = true
			} else {
				safe = false
				continue
			}

			oldVal = val
		}
		if safe {
			countSafe++
		}

	}
	return countSafe
}
