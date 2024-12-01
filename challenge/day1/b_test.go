package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {

	input := strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3
`)

	result := partB(input)

	require.Equal(t, 31, result)
}
