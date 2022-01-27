package selection_sort

// Difficulty: O(n*2)
func SelectionSortAsc(arr []int) {
	var ln = len(arr)
	for i := 0; i < ln; i++ {
		var minIdx = i
		for j := i; j < ln; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
