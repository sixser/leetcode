package arr

// SearchInsert
// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
func SearchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if nums[l] > target {
		return l
	} else if nums[r] < target {
		return r + 1
	}

	for l < r {
		if nums[(l+r)/2] == target {
			return (l + r) / 2
		} else if nums[(l+r)/2] < target {
			l = (l+r)/2 + 1
		} else {
			r = (l+r)/2 - 1
		}
	}

	if nums[l] < target {
		return l + 1
	} else {
		return l
	}
}
