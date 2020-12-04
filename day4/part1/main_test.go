package main

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		line string
		p    passport
	}{
		{
			line: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			p: passport{
				byr: "1937",
				iyr: "2017",
				eyr: "2020",
				ecl: "gry",
				hgt: "183cm",
				hcl: "#fffffd",
				pid: "860033327",
				cid: "147",
			},
		},
	}

	for _, test := range tests {
		p := parseLine(test.line)

		if p != test.p {
			t.Errorf("passport not properly parsed: %v", p)
		}
	}
}
