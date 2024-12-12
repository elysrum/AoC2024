package day12

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB1(t *testing.T) {

	input := strings.NewReader(`AAAA
BBCD
BBCC
EEEC`)

	result := partB(input)

	require.Equal(t, 80, result)
}
func TestB2(t *testing.T) {

	input := strings.NewReader(`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`)

	result := partB(input)

	require.Equal(t, 236, result)
}
func TestB3(t *testing.T) {

	input := strings.NewReader(`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`)

	result := partB(input)

	require.Equal(t, 368, result)
}
func TestB4(t *testing.T) {

	input := strings.NewReader(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)

	result := partB(input)

	require.Equal(t, 1206, result)
}
