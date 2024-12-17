package day17

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {

	input := strings.NewReader(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)

	result := partB(input)

	require.Equal(t, 117440, result)
}
