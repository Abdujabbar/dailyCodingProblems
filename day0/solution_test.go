package main

import "testing"

type Case struct {
	input    []int
	k        int
	expected bool
}

func TestSolution(t *testing.T) {
	cases := []Case{
		Case{
			input:    []int{1, 2, 3, 4},
			expected: true,
			k:        4,
		},
		Case{
			input:    []int{1, 5, 2, 4},
			expected: false,
			k:        30,
		},
		Case{
			input:    []int{10, 7, 3, 1},
			k:        17,
			expected: true,
		},
	}
	for _, cs := range cases {
		res := solution(cs.input, cs.k)
		if res != cs.expected {
			t.Errorf("Expected: %v, received: %v, on input: %v", cs.expected, res, cs.input)
		}
	}
}
