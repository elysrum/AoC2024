package day16

import (
	"AoC2024/challenge"
	"container/heap"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 16, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

type Point struct {
	row int
	col int
}

type Visit struct {
	location  Point
	direction Point
}

// An Item is something we manage in a priority queue.
type Item struct {
	visit    Visit // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func partA(input io.Reader) int {

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
