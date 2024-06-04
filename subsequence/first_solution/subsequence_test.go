package first_solution

import "testing"

type testCasesSubsequence struct {
	name        string
	array       []int
	subsequence []int
	expectation bool
}

func TestSubsequence(t *testing.T) {

	cases := []testCasesSubsequence{

		{
			name:        "Example 1",
			array:       []int{1, 2, 3, 4, 5},
			subsequence: []int{3, 4, 5},
			expectation: true,
		},
		{
			name:        "Example 2",
			array:       []int{-1, 0, 3},
			subsequence: []int{4, 5},
			expectation: false,
		},
		{
			name:        "Example 3",
			array:       []int{-3, 6, 7},
			subsequence: []int{6},
			expectation: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			answer := Subsequence(tc.array, tc.subsequence)

			if answer != tc.expectation {
				t.Errorf("expectation %v, but got %v", tc.expectation, answer)
			}

		})

	}

}
