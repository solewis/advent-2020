package main

import "testing"

func TestPasswordIsValid(t *testing.T) {
	password1 := passwordFile{"a", "abcde", 1, 3}
	password2 := passwordFile{"b", "cdefg", 1, 3}
	password3 := passwordFile{"c", "ccccccccc", 2, 9}

	t.Run("sled rental rules", func(t *testing.T) {
		if password1.isValidSledRentalRules() == false {
			t.Error("Expected passwordFile 1 to be valid")
		}
		if password2.isValidSledRentalRules() == true {
			t.Error("Expected passwordFile 2 to be invalid")
		}
		if password3.isValidSledRentalRules() == false {
			t.Error("Expected passwordFile 3 to be valid")
		}
	})

	t.Run("toboggan rules", func(t *testing.T) {
		if password1.isValidTobogganRules() == false {
			t.Error("Expected passwordFile 1 to be valid")
		}
		if password2.isValidTobogganRules() == true {
			t.Error("Expected passwordFile 2 to be invalid")
		}
		if password3.isValidTobogganRules() == true {
			t.Error("Expected passwordFile 3 to be invalid")
		}
	})
}
