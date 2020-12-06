package main

import "testing"

func TestGetYes(t *testing.T) {
	tests := []struct {
		g        group
		yesCount int
	}{
		{g: group{answers: []string{"abc"}}, yesCount: 3},
		{g: group{answers: []string{"a", "b", "c"}}, yesCount: 0},
		{g: group{answers: []string{"ab", "bc"}}, yesCount: 1},
		{g: group{answers: []string{"a", "a", "a", "a"}}, yesCount: 1},
		{g: group{answers: []string{"b"}}, yesCount: 1},
	}

	for _, test := range tests {
		yesCount := test.g.getYes()
		if len(yesCount) != test.yesCount {
			t.Errorf("%v - expected: %v, got: %v", test.g.answers, test.yesCount, yesCount)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		groups   []group
		yesCount int
	}{
		{
			groups: []group{
				{answers: []string{"abc"}},
				{answers: []string{"a", "b", "c"}},
				{answers: []string{"ab", "bc"}},
				{answers: []string{"a", "a", "a", "a"}},
				{answers: []string{"b"}},
			},
			yesCount: 6,
		},
	}

	for _, test := range tests {
		yesCount := solve(test.groups)
		if yesCount != test.yesCount {
			t.Errorf("expected: %v, got: %v", test.yesCount, yesCount)
		}
	}
}
