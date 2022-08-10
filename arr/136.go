package arr

// SingleNumber
// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
func SingleNumber(nums []int) int {
	i := 0
	for _, v := range nums {
		i ^= v
	}

	return i
}
