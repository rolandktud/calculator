// unit tests for main program
package main

import "testing"

// testing the helper function TestIsInt
func TestIsInt(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{1.0, true},
		{1.1, false},
	}

	for _, test := range tests {
		if got := isInt(test.input); got != test.want {
			t.Errorf("isInt(%f) = %v", test.input, got)
		}
	}
}
