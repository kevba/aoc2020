package main

import "testing"

func TestValidBirthyear(t *testing.T) {
	tests := []struct {
		val   string
		valid bool
	}{
		{"1920", true},
		{"1919", false},
		{"2002", true},
		{"2003", false},
		{"0", false},
		{"", false},
	}

	for _, test := range tests {
		if test.valid != validBirthYear(test.val) {
			t.Errorf("%v: should be valid: %v ", test.val, test.valid)
		}
	}
}

func TestValidIssueYear(t *testing.T) {
	tests := []struct {
		val   string
		valid bool
	}{
		{"2010", true},
		{"2009", false},
		{"2020", true},
		{"2021", false},
		{"0", false},
		{"", false},
	}

	for _, test := range tests {
		if test.valid != validIssueYear(test.val) {
			t.Errorf("%v: should be valid: %v ", test.val, test.valid)
		}
	}
}

func TestValidExpireYear(t *testing.T) {
	tests := []struct {
		val   string
		valid bool
	}{
		{"2020", true},
		{"2019", false},
		{"2030", true},
		{"2031", false},
		{"0", false},
		{"", false},
	}

	for _, test := range tests {
		if test.valid != validExpireYear(test.val) {
			t.Errorf("%v: should be valid: %v ", test.val, test.valid)
		}
	}
}

func TestValidHeight(t *testing.T) {
	tests := []struct {
		val   string
		valid bool
	}{
		{"150cm", true},
		{"193cm", true},
		{"149cm", false},
		{"194cm", false},
		{"59in", true},
		{"76in", true},
		{"58in", false},
		{"77in", false},
		{"77", false},
		{"", false},
	}

	for _, test := range tests {
		if test.valid != validHeigth(test.val) {
			t.Errorf("%v: should be valid: %v ", test.val, test.valid)
		}
	}
}

func TestValidHairColor(t *testing.T) {
	tests := []struct {
		val   string
		valid bool
	}{
		{"#123abc", true},
		{"#123abz", false},
		{"123abc", false},
		{"", false},
	}

	for _, test := range tests {
		if test.valid != validHairColor(test.val) {
			t.Errorf("%v: should be valid: %v ", test.val, test.valid)
		}
	}
}

func TestValidEyeColor(t *testing.T) {
	tests := []struct {
		val   string
		valid bool
	}{
		{"amb", true},
		{"blu", true},
		{"brn", true},
		{"gry", true},
		{"grn", true},
		{"hzl", true},
		{"oth", true},
		{"wat", false},
		{"", false},
	}

	for _, test := range tests {
		if test.valid != validEyeColor(test.val) {
			t.Errorf("%v: should be valid: %v ", test.val, test.valid)
		}
	}
}

func TestValidPID(t *testing.T) {
	tests := []struct {
		val   string
		valid bool
	}{
		{"000000001", true},
		{"0123456789", false},
		{"", false},
	}

	for _, test := range tests {
		if test.valid != validPID(test.val) {
			t.Errorf("%v: should be valid: %v ", test.val, test.valid)
		}
	}
}
