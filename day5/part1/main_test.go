package main

import "testing"

func TestGetRowNum(t *testing.T) {
	tests := []struct {
		str string
		val int
	}{
		{str: "BFFFBBFRRR", val: 70},
		{str: "FFFBBBFRRR", val: 14},
		{str: "BBFFBBFRLL", val: 102},
	}

	for _, test := range tests {
		val := getRowNum(test.str)

		if val != test.val {
			t.Errorf("expected: %v, got: %v", test.val, val)
		}
	}
}

func TestGetColNum(t *testing.T) {
	tests := []struct {
		str string
		val int
	}{
		{str: "BFFFBBFRRR", val: 7},
		{str: "FFFBBBFRRR", val: 7},
		{str: "BBFFBBFRLL", val: 4},
	}

	for _, test := range tests {
		val := getColNum(test.str)

		if val != test.val {
			t.Errorf("expected: %v, got: %v", test.val, val)
		}
	}
}

func TestDiv(t *testing.T) {
	lower := 0
	upper := 127

	tests := []struct {
		str   string
		lower int
		upper int
	}{
		{str: "F", lower: 0, upper: 63},
		{str: "FB", lower: 32, upper: 63},
		{str: "FBF", lower: 32, upper: 47},
		{str: "FBFB", lower: 40, upper: 47},
		{str: "FBFBB", lower: 44, upper: 47},
		{str: "FBFBBF", lower: 44, upper: 45},
	}

	for _, test := range tests {
		lower, upper := div(test.str, lower, upper)

		if lower != test.lower || upper != test.upper {
			t.Errorf("expected: %v - %v, got: %v - %v", test.lower, test.upper, lower, upper)
		}
	}
}

func TestGetID(t *testing.T) {
	tests := []struct {
		seat seat
		id   int
	}{
		{seat: seat{row: 70, col: 7}, id: 567},
		{seat: seat{row: 14, col: 7}, id: 119},
		{seat: seat{row: 102, col: 4}, id: 820},
	}

	for _, test := range tests {
		id := test.seat.id()

		if id != test.id {
			t.Errorf("expected: %v, got: %v", test.id, id)
		}
	}
}
