package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := [26]int{}

	for i := 0; i < len(s); i++ {
		charCount[s[i]-97]++
	}

	for i := 0; i < len(t); i++ {
		charIndex := t[i] - 97
		charCount[charIndex]--

		if charCount[charIndex] < 0 {
			return false
		}
	}

	for _, v := range charCount {
		if v != 0 {
			return false
		}
	}

	return true
}
