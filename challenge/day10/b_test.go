package day10

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {

	input := strings.NewReader(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)

	result := partB(input)

	require.Equal(t, 81, result)
}
