package selection_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSortAsc(t *testing.T) {
	arr := []int{6, 2, 8, 3, 1, 5, 9, 4, 2}
	SelectionSortAsc(arr)
	assert.Equal(t, []int{1, 2, 2, 3, 4, 5, 6, 8, 9}, arr)
}
