package main

import (
	"fmt"
	"testing"
)

type inputResult struct {
	name string
	args []string
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		result struct {
			name string
			args []string
		}
	}{
		{
			input: "test",
			result: inputResult{
				name: "test",
				args: []string{},
			},
		},
		{
			input: "   map   ",
			result: inputResult{
				name: "map",
				args: []string{},
			},
		},
		{
			input: "MAP",
			result: inputResult{
				name: "map",
				args: []string{},
			},
		},
		{
			input: "test 1 2 3",
			result: inputResult{
				name: "test",
				args: []string{"1", "2", "3"},
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			name, args := cleanInput(c.input)
			if name != c.result.name {
				t.Errorf("expected name to be %s but got %s", c.result.name, name)
			}
			if len(args) != len(c.result.args) {
				t.Errorf("args length of %v does not match the expected length of %v", len(args), len(c.result.args))
			}
			for i, arg := range args {
				if arg != c.result.args[i] {
					t.Errorf("expected arg to be %v but got %v", c.result.args[i], arg)
				}
			}
		})
	}
}
