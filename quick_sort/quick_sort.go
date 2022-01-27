package quick_sort

import "math"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else if len(arr) == 2 {
		if arr[0] < arr[1] {
			arr[1], arr[0] = arr[0], arr[1]
			return arr
		}
	}

	bearingEl := arr[int(math.Floor(float64(len(arr)/2)))]
	leftArr := make([]int, 0, len(arr))
	rightArr := make([]int, 0, len(arr))
	for _, i := range arr {
		if i > bearingEl {
			rightArr = append(rightArr, i)
		} else {
			leftArr = append(leftArr, i)
		}
	}

	newArr := make([]int, 0)
	newArr = append(newArr, QuickSort(leftArr)...)
	newArr = append(newArr, bearingEl)
	newArr = append(newArr, QuickSort(rightArr)...)
	return newArr
}
