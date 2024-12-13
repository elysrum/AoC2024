package day13

import (
	"AoC2024/challenge"
	"AoC2024/util"
	"fmt"
	"io"
	"math"
	"regexp"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 13, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {

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

		countA := (tx*by - ty*bx) / (ax*by - ay*bx)
		countB := (tx - ax*countA) / bx

		// Are countA and countB whole numbers?
		truncA := math.Trunc(countA)
		truncB := math.Trunc(countB)
		if truncA == countA && truncB == countB {
			if truncA <= 100 && truncB <= 100 {

				totalCost += int(countA*3 + countB)
			}
		}

	}
	return totalCost
}
