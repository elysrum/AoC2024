package day17

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {

	input := strings.NewReader(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)

	result := partA(input)

	require.Equal(t, "4,6,3,5,6,3,5,2,1,0", result)
}

func TestAPartBReverse(t *testing.T) {

	input := strings.NewReader(`Register A: 202975183645226
Register B: 0
Register C: 0

Program: 2,4,1,1,7,5,0,3,1,4,4,4,5,5,3,0`)

	result := partA(input)

	require.Equal(t, "2,4,1,1,7,5,0,3,1,4,4,4,5,5,3,0", result)
}
