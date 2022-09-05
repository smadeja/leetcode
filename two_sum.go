package neetcode

func twoSum(nums []int, target int) []int {
	ind := make(map[int]int)

	for i, n := range nums {
		w := target - n

		if j, ok := ind[w]; ok {
			return []int{i, j}
		} else {
			ind[n] = i
		}
	}

	return nil
}
