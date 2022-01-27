package quick_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert.Equal(
		t,
		[]int{0, 1, 2, 3, 4, 4, 5, 7, 8, 9},
		QuickSort([]int{8, 3, 1, 4, 9, 0, 2, 4, 5, 7}))
}

func TestQuickSort_ArrWithOneItem(t *testing.T) {
	assert.Equal(t, []int{5}, QuickSort([]int{5}))
}

func TestQuickSort_EmptyArr(t *testing.T) {
	assert.Nil(t, QuickSort([]int{}))
}
