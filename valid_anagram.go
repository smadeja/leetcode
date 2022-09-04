package neetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		charCount[t[i]]--

		if charCount[t[i]] < 0 {
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
