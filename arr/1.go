package arr

// TwoSum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
func TwoSum(nums []int, target int) []int {
	h := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if j, ok := h[target-nums[i]]; ok {
			return []int{j, i}
		}

		h[nums[i]] = i
	}

	return nil
}
