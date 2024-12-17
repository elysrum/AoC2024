package day16

import (
	"container/heap"
	"fmt"
	"io"

	"AoC2024/challenge"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 16, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

	data := challenge.Lines(input)

	var start, end Point

	var grid [][]rune

	rowIndex := 0
	for inputLine := range data {
		row := make([]rune, len(inputLine))
		for cellIndex, cell := range inputLine {
			row[cellIndex] = cell
			// find the guard
			if cell == 'S' {
				start.row = rowIndex
				start.col = cellIndex
			} else if cell == 'E' {

				end.row = rowIndex
				end.col = cellIndex
			}

		}
		grid = append(grid, row)
		rowIndex += 1
	}

	result := 0
	pq := make(PriorityQueue, 0)

	visit := Visit{location: Point{row: start.row, col: start.col},
		direction: Point{row: 0, col: 1}}

	pq = append(pq, &Item{visit: visit, priority: 0})
	seen := make(map[Visit]bool)
	seen[visit] = true

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		seen[item.visit] = true
		result = item.priority
		direction := item.visit.direction

		if grid[item.visit.location.row][item.visit.location.col] == 'E' {
			return item.priority
		}

		// Add Keep Heading in same direction
		newPriority := (*item).priority + 1
		newLocation := Point{(*item).visit.location.row + direction.row, (*item).visit.location.col + direction.col}
		newDirection := Point{(*item).visit.direction.row, (*item).visit.direction.col}

		newVisit := Visit{location: newLocation, direction: newDirection}

		if !seen[newVisit] && grid[newLocation.row][newLocation.col] != '#' {
			heap.Push(&pq, &Item{priority: newPriority, visit: newVisit})

		}

		// Add rotate 90
		newPriority = (*item).priority + 1000
		newLocation = Point{(*item).visit.location.row, (*item).visit.location.col}
		newDirection = Point{row: (*item).visit.direction.col, col: -(*item).visit.direction.row}

		newVisit = Visit{location: newLocation, direction: newDirection}

		if !seen[newVisit] && grid[newLocation.row][newLocation.col] != '#' {
			heap.Push(&pq, &Item{priority: newPriority, visit: newVisit})

		}

		// Add rotate 270
		newPriority = (*item).priority + 1000
		newLocation = Point{(*item).visit.location.row, (*item).visit.location.col}
		newDirection = Point{row: -(*item).visit.direction.col, col: (*item).visit.direction.row}

		newVisit = Visit{location: newLocation, direction: newDirection}

		if !seen[newVisit] && grid[newLocation.row][newLocation.col] != '#' {
			heap.Push(&pq, &Item{priority: newPriority, visit: newVisit})

		}

	}

	return result
}
