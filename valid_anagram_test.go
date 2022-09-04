package neetcode

import (
	"fmt"
	"testing"
)

type testCase struct {
	stringOne string
	stringTwo string
	expected  bool
}

func TestIsAnagram(t *testing.T) {
	testCases := []testCase{
		{
			stringOne: "anagram",
			stringTwo: "nagaram",
			expected:  true,
		},
		{
			stringOne: "rat",
			stringTwo: "car",
			expected:  false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test#%d", i+1), func(t *testing.T) {
			actual := isAnagram(tc.stringOne, tc.stringTwo)

			if actual != tc.expected {
				t.Errorf("isAnagram(%#v, %#v) = %t; expected %t",
					tc.stringOne,
					tc.stringTwo,
					actual,
					tc.expected,
				)
			}
		})
	}
}
