package day9

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {

	input := strings.NewReader(`2333133121414131402`)

	result := partB(input)

	require.Equal(t, 2858, result)
}
