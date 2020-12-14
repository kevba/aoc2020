package main

import "testing"

func TestCountSkipable(t *testing.T) {
	tests := []struct {
		adapters []int
		count    int
	}{
		{[]int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22}, 3},
		{[]int{0, 1, 4, 7, 10, 12, 15, 16, 19, 22}, 0},
		{[]int{0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31,
			32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49, 52}, 15},
	}

	for _, test := range tests {
		count := countSkippable(test.adapters)
		if count != test.count {
			t.Errorf("expected %v got %v", test.count, count)
		}
	}
}
