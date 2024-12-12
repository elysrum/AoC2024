package day13

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {

	input := strings.NewReader(`125 17`)

	result := partB(input)

	require.Equal(t, 65601038650482, result)
}
