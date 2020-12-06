package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestValidCount(t *testing.T) {
	lines := parse.Lines("test_input.txt")
	passports := parsePassports(lines)
	if countValid(passports, passport.isValid) != 2 {
		t.Error("Expected 2 valid passports")
	}
}

func TestValidStrict(t *testing.T) {
	t.Run("birth year at least 1920", func(t *testing.T) {
		p1, p2 := validPassport(), validPassport()
		p1.birthYear = 1919
		p2.birthYear = 1920
		if p1.isValidStrict() {
			t.Error("Expected invalid passport for birth year 1919")
		}
		if !p2.isValidStrict() {
			t.Error("Expected valid passport for birth year 1920")
		}
	})
	t.Run("birth year at most 2002", func(t *testing.T) {
		p1, p2 := validPassport(), validPassport()
		p1.birthYear = 2003
		p2.birthYear = 2002
		if p1.isValidStrict() {
			t.Error("Expected invalid passport for birth year 2003")
		}
		if !p2.isValidStrict() {
			t.Error("Expected valid passport for birth year 2002")
		}
	})
	t.Run("issue year at least 2010", func(t *testing.T) {
		p1, p2 := validPassport(), validPassport()
		p1.issueYear = 2009
		p2.issueYear = 2010
		if p1.isValidStrict() {
			t.Errorf("Expected invalid passport for year %d", p1.issueYear)
		}
		if !p2.isValidStrict() {
			t.Errorf("Expected valid passport for year %d", p2.issueYear)
		}
	})
	t.Run("issue year at most 2020", func(t *testing.T) {
		p1, p2 := validPassport(), validPassport()
		p1.issueYear = 2021
		p2.issueYear = 2020
		if p1.isValidStrict() {
			t.Errorf("Expected invalid passport for year %d", p1.issueYear)
		}
		if !p2.isValidStrict() {
			t.Errorf("Expected valid passport for year %d", p2.issueYear)
		}
	})
	t.Run("expiration year at least 2020", func(t *testing.T) {
		p1, p2 := validPassport(), validPassport()
		p1.expirationYear = 2019
		p2.expirationYear = 2020
		if p1.isValidStrict() {
			t.Errorf("Expected invalid passport for year %d", p1.expirationYear)
		}
		if !p2.isValidStrict() {
			t.Errorf("Expected valid passport for year %d", p2.expirationYear)
		}
	})
	t.Run("expiration year at most 2030", func(t *testing.T) {
		p1, p2 := validPassport(), validPassport()
		p1.expirationYear = 2031
		p2.expirationYear = 2030
		if p1.isValidStrict() {
			t.Errorf("Expected invalid passport for year %d", p1.expirationYear)
		}
		if !p2.isValidStrict() {
			t.Errorf("Expected valid passport for year %d", p2.expirationYear)
		}
	})
	t.Run("height valid", func(t *testing.T) {
		t.Run("unit in cm or in", func(t *testing.T) {
			p1, p2, p3 := validPassport(), validPassport(), validPassport()
			p1.height = "10ft"
			p2.height = "160"
			p3.height = "x"
			if p1.isValidStrict() || p2.isValidStrict() || p3.isValidStrict() {
				t.Error("Failed height unit check")
			}
		})
		t.Run("cm at least 150 and at most 193", func(t *testing.T) {
			p1, p2, p3, p4 := validPassport(), validPassport(), validPassport(), validPassport()
			p1.height = "149cm"
			p2.height = "150cm"
			p3.height = "193cm"
			p4.height = "194cm"
			if p1.isValidStrict() || !p2.isValidStrict() || !p3.isValidStrict() || p4.isValidStrict() {
				t.Error("cm height check failed")
			}
		})
		t.Run("in at least 59 and at most 76", func(t *testing.T) {
			p1, p2, p3, p4 := validPassport(), validPassport(), validPassport(), validPassport()
			p1.height = "58in"
			p2.height = "59in"
			p3.height = "76in"
			p4.height = "77in"
			if p1.isValidStrict() || !p2.isValidStrict() || !p3.isValidStrict() || p4.isValidStrict() {
				t.Error("in height check failed")
			}
		})
	})
	t.Run("hair color valid", func(t *testing.T) {
		t.Run("starts with #", func(t *testing.T) {
			p := validPassport()
			p.hairColor = "!123abc"
			if p.isValidStrict() {
				t.Error("Expected invalid hair color that doesn't start with #")
			}
		})
		t.Run("6 characters", func(t *testing.T) {
			p1, p2 := validPassport(), validPassport()
			p1.hairColor = "#123ab"
			p2.hairColor = "#123abcd"
			if p1.isValidStrict() || p2.isValidStrict() {
				t.Error("Hair color 6 character check failed")
			}
		})
		t.Run("0-9 and a-f allowed", func(t *testing.T) {
			p1, p2, p3, p4, p5 := validPassport(), validPassport(), validPassport(), validPassport(), validPassport()
			p1.hairColor = "#012345"
			p2.hairColor = "#6789ab"
			p3.hairColor = "#cdef01"
			p4.hairColor = "#ghijkl"
			p5.hairColor = "#01234?"
			if !p1.isValidStrict() || !p2.isValidStrict() || !p3.isValidStrict() || p4.isValidStrict() || p5.isValidStrict() {
				t.Error("hair color character check failed")
			}
		})
	})
	t.Run("eye color valid", func(t *testing.T) {
		p1, p2, p3, p4, p5, p6, p7, p8 := validPassport(), validPassport(), validPassport(), validPassport(),
			validPassport(), validPassport(), validPassport(), validPassport()
		p1.eyeColor = "amb"
		p2.eyeColor = "blu"
		p3.eyeColor = "brn"
		p4.eyeColor = "gry"
		p5.eyeColor = "grn"
		p6.eyeColor = "hzl"
		p7.eyeColor = "oth"
		p8.eyeColor = "xxx"
		if !p1.isValidStrict() || !p2.isValidStrict() || !p3.isValidStrict() || !p4.isValidStrict() ||
			!p5.isValidStrict() || !p6.isValidStrict() || !p7.isValidStrict() || p8.isValidStrict() {
			t.Error("eye color check failed")
		}
	})
	t.Run("passport id valid", func(t *testing.T) {
		t.Run("is length 9", func(t *testing.T) {
			p1, p2, p3 := validPassport(), validPassport(), validPassport()
			p1.passportId = "123456789"
			p2.passportId = "12345678"
			p3.passportId = "0123456789"
			if !p1.isValidStrict() || p2.isValidStrict() || p3.isValidStrict() {
				t.Error("failed length check")
			}
		})
		t.Run("is all digits", func(t *testing.T) {
			p1, p2 := validPassport(), validPassport()
			p1.passportId = "123456789"
			p2.passportId = "12345678a"
			if !p1.isValidStrict() || p2.isValidStrict() {
				t.Error("failed digit check")
			}
		})
	})
}

func validPassport() passport {
	return passport{
		birthYear:      1940,
		issueYear:      2015,
		expirationYear: 2024,
		height:         "160cm",
		hairColor:      "#123abc",
		eyeColor:       "brn",
		passportId:     "012345678",
	}
}
