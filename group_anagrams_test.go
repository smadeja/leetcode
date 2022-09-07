package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

type testCase struct {
	input    []string
	expected [][]string
}

func testCases() []testCase {
	return []testCase{
		testCase{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},

			expected: [][]string{
				[]string{"ate", "eat", "tea"},
				[]string{"bat"},
				[]string{"nat", "tan"},
			},
		},
		testCase{
			input:    []string{""},
			expected: [][]string{[]string{""}},
		},
		testCase{
			input:    []string{"a"},
			expected: [][]string{[]string{"a"}},
		},
	}
}

func TestGroupAnagrams(t *testing.T) {
	for i, tc := range testCases() {
		t.Run(fmt.Sprintf("Test#%d", i+1), func(t *testing.T) {
			act := groupAnagrams(tc.input)
			sortStringSliceSlice(act)

			if !equalStringSliceSlices(act, tc.expected) {
				t.Errorf("groupAnagrams: got %v (sorted), expected %v", act, tc.expected)
			}
		})
	}
}

func sortStringSliceSlice(s [][]string) {
	for _, v := range s {
		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
	}

	sort.Slice(s, func(i, j int) bool { return strings.Join(s[i], "") < strings.Join(s[j], "") })
}

func equalStringSliceSlices(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !equalStringSlices(a[i], b[i]) {
			return false
		}
	}

	return true
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
