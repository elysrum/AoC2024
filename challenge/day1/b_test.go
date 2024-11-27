package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Skip("Not Implemented")

	input := strings.NewReader("foobar")

	result := partB(input)

	require.Equal(t, 42, result)
}
