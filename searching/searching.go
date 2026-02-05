package searching

import "cmp"

// LinearSearch looks for the needle in the pre-sorted haystack and returns the
// first index of its occurrence from the left, or -1 if it isn't found.
func LinearSearch[T comparable](haystack []T, needle T) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return i
		}
	}
	return -1
}

// BinarySearch looks for the needle in the pre-sorted haystack and returns the
// index of its occurrence, or -1 if it isn't found.
func BinarySearch[T cmp.Ordered](haystack []T, needle T) int {
	n := len(haystack)
	i, j := 0, n-1
	for i <= j {
		m := i + (j-i+1)/2
		if needle == haystack[m] {
			return m
		}
		if cmp.Less(needle, haystack[m]) {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return -1
}
