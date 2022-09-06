package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

type testCase struct {
	numbers  []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	tcs := []testCase{
		{
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			numbers:  []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			numbers:  []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test#%d", i+1), func(t *testing.T) {
			act := twoSum(tc.numbers, tc.target)

			if !equal(act, tc.expected) {
				t.Errorf("twoSum(%#v, %#v) = %#v; expected %#v",
					tc.numbers,
					tc.target,
					act,
					tc.expected,
				)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	s1 := make([]int, len(a))
	copy(s1, a)

	s2 := make([]int, len(b))
	copy(s2, b)

	sort.Ints(s1)
	sort.Ints(s2)

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
