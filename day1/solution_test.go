package main

import "testing"
import "reflect"

type Case struct {
	input    []int
	expected []int
}

func TestSolution(t *testing.T) {
	cases := []Case{
		Case{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{120, 60, 40, 30, 24},
		},
		Case{
			input:    []int{3, 2, 1},
			expected: []int{2, 3, 6},
		},
	}

	for _, cs := range cases {
		res := solution(cs.input)
		if !reflect.DeepEqual(cs.expected, res) {
			t.Errorf("Expected: %v, Received: %v, on input: %v", cs.expected, res, cs.input)
		}
	}
}
