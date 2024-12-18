package day18

import (
	"container/list"
	"fmt"
	"io"
	"strings"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 18, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

	gridSize := 70

	data := challenge.Lines(input)

	coords := make([]Point, 0)

	for inputLine := range data {

		coord := strings.Split(inputLine, ",")

		x := util.MustAtoI(coord[0])
		y := util.MustAtoI(coord[1])

		coords = append(coords, Point{row: y, col: x})

	}

	low := 0
	high := len(coords) - 1

	for low < high {

		mid := ((low + high) / 2)

		if binSearch(coords, gridSize, mid+1) {
			low = mid + 1
		} else {
			high = mid
		}
	}

	fmt.Printf("%d,%d\n", coords[low].col, coords[low].row)

	return coords[low].row*10 + coords[low].col

}

func binSearch(coords []Point, gridSize int, itemCount int) bool {

	start := Point{0, 0}
	end := Point{gridSize, gridSize}

	grid := make([][]rune, gridSize+1)

	for i := range gridSize + 1 {
		grid[i] = make([]rune, gridSize+1)
	}

	// Set the grid up
	for _, coord := range coords[:itemCount] {

		grid[coord.row][coord.col] = 1
	}

	q := list.New()

	movement := []Point{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	q.PushBack(Item{location: start, steps: 0})
	seen := make(map[Point]bool)
	seen[start] = true

	for q.Len() > 0 {
		// Equivalent to Popleft
		item := q.Front().Value.(Item)
		q.Remove(q.Front())

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
				return true
			}

			seen[newLocation] = true
			q.PushBack(Item{steps: item.steps + 1, location: newLocation})

		}
	}

	return false
}
