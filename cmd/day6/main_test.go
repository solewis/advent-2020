package main

import "testing"

func TestCustomForm(t *testing.T) {
	t.Run("counts number of yes distinctly", func(t *testing.T) {
		f := customsForm{[]string{"abcx", "abcy", "abcz"}}
		if f.totalYeses() != 6 {
			t.Errorf("Expected 6 yesses but was %d", f.totalYeses())
		}
		f = customsForm{[]string{"a", "b", "c"}}
		if f.totalYeses() != 3 {
			t.Errorf("Expected 3 yesses but was %d", f.totalYeses())
		}
		f = customsForm{[]string{"a", "a", "a"}}
		if f.totalYeses() != 1 {
			t.Errorf("Expected 1 yesses but was %d", f.totalYeses())
		}
		f = customsForm{[]string{"abc"}}
		if f.totalYeses() != 3 {
			t.Errorf("Expected 3 yesses but was %d", f.totalYeses())
		}
	})

	t.Run("counts number of consensus yes", func(t *testing.T) {
		f := customsForm{[]string{"abcx", "abcy", "abcz"}}
		if f.totalConsensusYeses() != 3 {
			t.Errorf("Expected 3 yesses but was %d", f.totalConsensusYeses())
		}
		f = customsForm{[]string{"a", "b", "c"}}
		if f.totalConsensusYeses() != 0 {
			t.Errorf("Expected 0 yesses but was %d", f.totalConsensusYeses())
		}
		f = customsForm{[]string{"a", "a", "a"}}
		if f.totalConsensusYeses() != 1 {
			t.Errorf("Expected 1 yesses but was %d", f.totalConsensusYeses())
		}
		f = customsForm{[]string{"abc"}}
		if f.totalConsensusYeses() != 3 {
			t.Errorf("Expected 3 yesses but was %d", f.totalConsensusYeses())
		}
	})
}
