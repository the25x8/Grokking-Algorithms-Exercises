package binary_search

// Difficulty: O(log(n))
func BinarySearchStr(list []string, value string) int {
	// We skip it this time
	// return sort.SearchStrings(list, value)

	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2

		if list[mid] == value {
			return mid
		} else if list[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
