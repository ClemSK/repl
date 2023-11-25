package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   ", // empty strings
			expected: []string{},
		},
		{
			input:    "  this   ", // empty strings
			expected: []string{"this"},
		},
		{
			input:    "  this and that  ", // empty strings
			expected: []string{"this", "and", "that"},
		},
		{
			input:    "  ThiS AnD ThaT  ", // empty strings
			expected: []string{"this", "and", "that"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
