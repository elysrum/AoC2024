package day7

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	data := challenge.Lines(input)

	result := 0

	for line := range data {

		strResult := strings.Split(line, ":")
		tgtResult := util.MustAtoI(strings.Trim(strResult[0], " "))
		strValues := strings.Split(strings.Trim(strResult[1], " "), " ")

		srcValues := make([]int, len(strValues))
		for idx, val := range strValues {

			srcValues[idx] = util.MustAtoI(strings.Trim(val, " "))
		}

		operators := len(srcValues) - 1

		startOp := make([]rune, operators)

		for i := 0; i < operators; i++ {
			startOp[i] = '+'
		}

		if calculateB(tgtResult, srcValues) {
			result += tgtResult
		}

	}
	return result

}

func calculateB(tgtValue int, srcValues []int) bool {

	newSrc := make([]int, 1)

	if len(srcValues) == 1 {
		return srcValues[0] == tgtValue
	}
	newSrc[0] = srcValues[0] + srcValues[1]
	if calculateB(tgtValue, slices.Concat(newSrc, srcValues[2:])) {
		return true
	}
	newSrc[0] = srcValues[0] * srcValues[1]
	if calculateB(tgtValue, slices.Concat(newSrc, srcValues[2:])) {
		return true
	}
	newSrc[0] = util.MustAtoI(fmt.Sprintf("%d%d", srcValues[0], srcValues[1]))
	if calculateB(tgtValue, slices.Concat(newSrc, srcValues[2:])) {
		return true
	}
	return false

}
