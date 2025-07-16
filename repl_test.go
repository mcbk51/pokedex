package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Charmander Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},

		{
			input:    " Giants Cowboys Eagles Commanders ",
			expected: []string{"giants", "cowboys", "eagles", "commanders"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch for input %q: expected %d words, got %d", c.input, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Word mismatch for input %q at index %d: expected %q, got %q", c.input, i, expectedWord, word)
			}
		}
	}
}
