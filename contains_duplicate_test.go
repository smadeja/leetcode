package leetcode

import (
	"fmt"
	"testing"
)

type testCase struct {
	input    []int
	expected bool
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []testCase{
		{
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test#%d", i+1), func(t *testing.T) {
			actual := containsDuplicate(tc.input)

			if actual != tc.expected {
				t.Errorf("containsDuplicate(%#v) = %t; expected %t", tc.input, actual, tc.expected)
			}
		})
	}
}
