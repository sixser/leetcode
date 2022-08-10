package arr

// Rotate
// Given an array, rotate the array to the right by k steps, where k is non-negative.
func Rotate(nums []int, k int) {
	reverse := func(nums []int, start int, end int) {
		for start < end {
			nums[start], nums[end-1] = nums[end-1], nums[start]
			start++
			end--
		}
	}

	k = k % len(nums)
	reverse(nums, 0, len(nums))
	reverse(nums, 0, k)
	reverse(nums, k, len(nums))
}
