package leetcode

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)

	for _, v := range nums {
		cnt[v]++
	}

	freq := make([][]int, len(nums))
	maxi := 0

	for k, v := range cnt {
		i := v - 1
		freq[i] = append(freq[i], k)

		if i > maxi {
			maxi = i
		}
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
