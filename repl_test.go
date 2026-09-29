package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "Bulbasaur", "PIKACHU"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Println(actual)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			fmt.Printf("%v, %v", len(actual), len(c.expected))
			t.Errorf("err: input and expected do not match in length")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			fmt.Printf("%v, %v", word, expectedWord)
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("err: word does not match expected word")
				t.Errorf("Expected: %v, input: %v", expectedWord, word)
			}
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
