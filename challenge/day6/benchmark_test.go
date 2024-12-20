package day6

import (
	"testing"

	"AoC2024/challenge"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partA(challenge.InputFile())
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partB(challenge.InputFile())
	}
}
