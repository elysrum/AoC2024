package day11

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {

	input := strings.NewReader(`125 17`)

	result := partA(input)

	require.Equal(t, 55312, result)
}
