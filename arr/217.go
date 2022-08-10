package arr

// ContainsDuplicate
// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
func ContainsDuplicate(nums []int) bool {
	hash := map[int]bool{}
	for _, v := range nums {
		if hash[v] {
			return true
		}

		hash[v] = true
	}

	return false
}
