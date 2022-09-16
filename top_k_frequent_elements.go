package leetcode

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	maxi := 0

	for _, v := range nums {
		c := cnt[v] + 1
		cnt[v] = c

		if c > maxi {
			maxi = c
		}
	}

	freq := make([][]int, maxi+1)

	for k, v := range cnt {
		i := v - 1
		freq[i] = append(freq[i], k)
	}

	res := make([]int, 0, k)

	for i := maxi; i >= 0; i-- {
		for _, v := range freq[i] {
			res = append(res, v)
		}

		if len(res) == k {
			break
		}
	}

	return res
}
