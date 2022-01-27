const selectionSort = require('./selection_sort')

const arr = [4,5,6,9,1,3,2,8,3]

test("Sort array by ascending", () => {
    expect(selectionSort(arr)).toEqual([1,2,3,3,4,5,6,8,9])
})