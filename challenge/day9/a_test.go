package day9

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {

	input := strings.NewReader(`2333133121414131402`)

	result := partA(input)

	require.Equal(t, 1928, result)
}
