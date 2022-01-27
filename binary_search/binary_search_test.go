package binary_search

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var arr = []string{"Egor", "David", "Zack", "Alex", "Job", "Darell"}

func init() {
	sort.Strings(arr)
}

func TestBinarySearch_FindLastElement(t *testing.T) {
	idx := BinarySearchStr(arr, "Alex")
	assert.Equal(t, 0, idx)
}

func TestBinarySearch_FindFirstElement(t *testing.T) {
	idx := BinarySearchStr(arr, "Zack")
	assert.Equal(t, len(arr)-1, idx)
}

func TestBinarySearch_EmptyResult(t *testing.T) {
	idx := BinarySearchStr(arr, "hello")
	assert.Equal(t, -1, idx)
}
