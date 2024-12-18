package day18

import (
	"AoC2024/challenge"
	"AoC2024/util"
	"container/list"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 18, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

type Point struct {
	row int
	col int
}

// An Item is something we manage in a priority queue.
type Item struct {
	location Point // The value of the item; arbitrary.
	steps    int   // The priority of the item in the queue.

}

func partA(input io.Reader) int {

	gridSize := 70
	itemCount := 1024

	data := challenge.Lines(input)

	start := Point{0, 0}
	end := Point{gridSize, gridSize}

	grid := make([][]rune, gridSize+1)

	for i := range gridSize + 1 {
		grid[i] = make([]rune, gridSize+1)
	}

	coords := make([]Point, 0)

	for inputLine := range data {

		coord := strings.Split(inputLine, ",")

		x := util.MustAtoI(coord[0])
		y := util.MustAtoI(coord[1])

		coords = append(coords, Point{row: y, col: x})

	}

	// Set the grid up
	for _, coord := range coords[:itemCount] {

		grid[coord.row][coord.col] = 1
	}

	result := 0
	q := list.New()

	movement := []Point{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	q.PushBack(Item{location: start, steps: 0})
	seen := make(map[Point]bool)
	seen[start] = true

	for q.Len() > 0 {
		// Equivalent to Popleft
		item := q.Front().Value.(Item)
		q.Remove(q.Front())

		result = item.steps

		// Check each direction from here
		for _, newDirection := range movement {

			newLocation := Point{item.location.row + newDirection.row, item.location.col + newDirection.col}

			if newLocation.row < 0 ||
				newLocation.col < 0 ||
				newLocation.row > gridSize ||
				newLocation.col > gridSize {
				continue
			}
			if seen[newLocation] {
				continue
			}
			if grid[newLocation.row][newLocation.col] == 1 {
				continue
			}
			if newLocation.row == end.row && newLocation.col == end.col {
				result = item.steps + 1
				return result
			}

			seen[newLocation] = true
			q.PushBack(Item{steps: item.steps + 1, location: newLocation})

		}
	}

	return result
}
