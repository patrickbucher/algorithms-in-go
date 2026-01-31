package sorting

import (
	"cmp"
	"math/rand"
)

func Swap[T cmp.Ordered](s []T, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

func RandomIntSlice(n int) []int {
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = rand.Intn(n)
	}
	return xs
}
