package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Skip("Not Implemented")

	input := strings.NewReader("foobar")

	result := partA(input)

	require.Equal(t, 42, result)
}
