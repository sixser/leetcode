package arr

// HeightChecker
// A school is trying to take an annual photo of all the students. The students are asked to stand in a single file line in non-decreasing order by height. Let this ordering be represented by the integer array expected where expected[i] is the expected height of the ith student in line.
// You are given an integer array heights representing the current order that the students are standing in. Each heights[i] is the height of the ith student in line (0-indexed).
// Return the number of indices where heights[i] != expected[i].
func HeightChecker(heights []int) int {
	i, j, l, e, tc := 0, 0, 0, 0, 0
	for i < len(heights) {
		if j >= len(heights) {
			if l > i || i >= l+e {
				tc++
			}

			i++
			j, l, e = 0, 0, 0
			continue
		}

		if heights[i] > heights[j] {
			l++
		} else if heights[i] == heights[j] {
			e++
		}

		j++
	}

	return tc
}
