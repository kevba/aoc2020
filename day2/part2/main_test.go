package main

import "testing"

func TestGetValid(t *testing.T) {
	tests := []struct {
		p     passwordLine
		valid bool
	}{
		{
			p:     passwordLine{min: 1, max: 3, char: "a", password: "abcde"},
			valid: true,
		},
		{
			p:     passwordLine{min: 1, max: 3, char: "b", password: "cdefg"},
			valid: false,
		},
		{
			p:     passwordLine{min: 2, max: 9, char: "c", password: "ccccccccc"},
			valid: false,
		},
	}

	for _, test := range tests {
		valid := validatePassword(test.p)
		if valid != test.valid {
			t.Errorf("is valid: %v; expected: %v", valid, test.valid)
		}
	}
}
