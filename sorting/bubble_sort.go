package sorting

import "cmp"

func BubbleSort[T cmp.Ordered](numbers []T) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if i == j {
				continue
			}
			if cmp.Less(numbers[j], numbers[j-1]) {
				Swap(numbers, j-1, j)
			}
		}
	}
}
