package first

import (
	"testing"
)

type twoNumberSumTestCase struct {
	name     string
	numbers  []int
	target   int
	expected []int
}

func TestTwoNumberSum(t *testing.T) {
	testCases := []twoNumberSumTestCase{
		{
			name:     "Example 1",
			numbers:  []int{2, 3, -1, -4, 10, 15},
			target:   11,
			expected: []int{-4, 15},
		},
		{
			name:     "Example 2",
			numbers:  []int{},
			target:   3,
			expected: []int{},
		},
		{
			name:     "Example 3",
			numbers:  []int{2, 1, 5, -4, 9, -10, 4},
			target:   6,
			expected: []int{2, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			answer := TwoNumberSum(tc.numbers, tc.target)

			if len(answer) != len(tc.expected) {
				t.Errorf("expectation %v, but got %v", tc.expected, answer)
			}

			for i := 0; i < len(answer); i++ {
				if answer[i] != tc.expected[i] {
					t.Errorf("expectation %v, but got %v", tc.expected, answer)
				}
			}
		})
	}
}
