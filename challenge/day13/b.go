package day13

import (
	"fmt"
	"io"
	"math"
	"regexp"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 13, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	re, err := regexp.Compile("\\d+")
	if err != nil {
		panic(err)
	}

	totalCost := 0

	// Finally a math problem
	data := challenge.Sections(input)

	for block := range data {

		items := re.FindAllString(block, -1)

		ax := float64(util.MustAtoI(items[0]))
		ay := float64(util.MustAtoI(items[1]))
		bx := float64(util.MustAtoI(items[2]))
		by := float64(util.MustAtoI(items[3]))
		tx := float64(util.MustAtoI(items[4]))
		ty := float64(util.MustAtoI(items[5]))
		tx += 10000000000000.0
		ty += 10000000000000.0

		countA := (tx*by - ty*bx) / (ax*by - ay*bx)
		countB := (tx - ax*countA) / bx

		// Are countA and countB whole numbers?
		truncA := math.Trunc(countA)
		truncB := math.Trunc(countB)
		if truncA == countA && truncB == countB {

			totalCost += int(countA*3 + countB)
		}

	}
	return totalCost
}
