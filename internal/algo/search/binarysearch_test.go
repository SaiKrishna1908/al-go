package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tc := []struct {
		arr         []int
		target      int
		expectedPos int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			3,
			2,
		},
		{
			[]int{7, 10, 13, 18, 20},
			20,
			4,
		},
		{
			[]int{15, 18, 20, 21, 23},
			19,
			1,
		},
		{
			[]int{16, 20, 24, 56, 65, 90},
			57,
			4,
		},
	}

	for _, tt := range tc {
		targetPosition := BinarySearch(tt.arr, tt.target)

		if targetPosition != tt.expectedPos {
			t.Errorf("Failed expected=%d, got=%d", tt.expectedPos, targetPosition)
		}
	}
}
