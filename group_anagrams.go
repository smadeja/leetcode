package leetcode

func groupAnagrams(strs []string) [][]string {
	grps := make(map[[26]byte][]string)

	for _, s := range strs {
		dgst := stringDigest(s)
		grps[dgst] = append(grps[dgst], s)
	}

	res := make([][]string, 0, len(grps))

	for _, v := range grps {
		res = append(res, v)
	}

	return res
}

func stringDigest(s string) [26]byte {
	dgst := [26]byte{}

	for _, c := range s {
		dgst[c-97]++
	}

	return dgst
}
