package leetcode

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums     []int
	k        int
	expected []int
}

func TestTopKFrequent(t *testing.T) {
	tcs := []testCase{
		testCase{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		testCase{
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		testCase{
			nums:     []int{1, 2, 2, 3, 3, 3},
			k:        2,
			expected: []int{3, 2},
		},
		testCase{
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
			k:        10,
			expected: []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test#%d", i+1), func(t *testing.T) {
			act := topKFrequent(tc.nums, tc.k)

			if !sameElements(act, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, act)
			}
		})
	}
}

func sameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	f := make(map[int]int)

	for _, v := range a {
		f[v]++
	}

	for _, v := range b {
		f[v]--

		if f[v] < 0 {
			return false
		}
	}

	for _, v := range f {
		if v != 0 {
			return false
		}
	}

	return true
}
