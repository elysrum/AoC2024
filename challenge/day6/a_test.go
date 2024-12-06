package day6

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {

	input := strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)

	result := partA(input)

	require.Equal(t, 41, result)
}
