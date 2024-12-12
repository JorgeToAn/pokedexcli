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
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "bulbasaur SQUIRTLE Charmander",
			expected: []string{"bulbasaur", "squirtle", "charmander"},
		},
		{
			input:    "TesT",
			expected: []string{"test"},
		},
		{
			input:    "   map   ",
			expected: []string{"map"},
		},
		{
			input:    "MAP",
			expected: []string{"map"},
		},
		{
			input:    "test 1 2 3",
			expected: []string{"test", "1", "2", "3"},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			words := cleanInput(c.input)
			if len(words) != len(c.expected) {
				t.Errorf("args length of %v does not match the expected length of %v", len(words), len(c.expected))
			}
			for i, word := range words {
				if word != c.expected[i] {
					t.Errorf("expected word to be %v but got %v", c.expected[i], word)
				}
			}
		})
	}
}
