package util

import "strconv"

func MustAtoI(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return result
}
func MustAtoUI(a string) uint64 {
	result, err := strconv.ParseUint(a, 10, 64)
	if err != nil {
		panic(err)
	}

	return result
}
