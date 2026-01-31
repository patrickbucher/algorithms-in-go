package sorting

import (
	"testing"
)

func TestInsertionSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, InsertionSort, numbers)
}

func TestInsertionSortLarge(t *testing.T) {
	testRandInt(t, InsertionSort, 1000)
}
