package autocomplete

import "testing"

func TestComputeDistance(t *testing.T) {
	tests := []struct {
		a, b           string
		expectedResult int
	}{
		{"", "hello", 5},
		{"hello", "", 5},
		{"hello", "hello", 0},
		{"ab", "aa", 1},
		{"ab", "ba", 2},
		{"ab", "aaa", 2},
		{"bbb", "a", 3},
		{"kitten", "sitting", 3},
		{"distance", "difference", 5},
		{"levenshtein", "frankenstein", 6},
		{"resume and cafe", "resumes and cafes", 2},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
		{"levenshtein", "frankenstein", 6},
		{"levenshtein", "frankenstein", 6},
	}
	for i, d := range tests {
		t.Run("", func(t *testing.T) {
			n := ComputeDistance(d.a, d.b)
			if n != d.expectedResult {
				t.Errorf("Test[%d]: ComputeDistance(%q,%q) returned %v, want %v",
					i, d.a, d.b, n, d.expectedResult)
			}
		})
	}
}
