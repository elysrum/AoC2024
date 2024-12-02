package day2

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"AoC2024/challenge"
	"AoC2024/util"
	"AoC2024/util/gmath"

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

	countSafe := 0
outerloop:
	for inputLine := range data {
		origLineParts := strings.Split(inputLine, " ")
		lineParts := slices.Clone(origLineParts)
		numParts := len(lineParts)

		if checkLine(lineParts) {
			countSafe++
			continue
		} else {
			for i := 0; i < numParts; i++ {
				lineParts = slices.Delete(slices.Clone(origLineParts), i, i+1)

				if checkLine(lineParts) {
					countSafe++
					continue outerloop
				}
			}

		}

	}
	return countSafe
}

func checkLine(lineParts []string) bool {
	safe := true
	first := true
	increasing := 0
	val, oldVal := 0, 0

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
			break
		}

		oldVal = val
	}

	return safe
}
