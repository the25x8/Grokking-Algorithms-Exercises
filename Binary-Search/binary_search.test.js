const binarySearch = require("./binary_search")

const arr = ["Egor", "David", "Zack", "Alex", "Job", "Darell"]
arr.sort()

test("Find index of the last element in array, equals = 5",  () => {
    expect(binarySearch(arr, "Zack")).toBe(5)
})

test("Find index of the first element in array, equals = 0",  () => {
    expect(binarySearch(arr, "Alex")).toBe(0)
})

test("Find non-existent element in array, equals = -1",  () => {
    expect(binarySearch(arr, "hello")).toBe(-1)
})