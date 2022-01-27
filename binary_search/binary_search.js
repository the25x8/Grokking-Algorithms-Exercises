// Difficulty: O(log(n))
function binarySearch(list, value) {
    let low = 0
    let high = list.length - 1

    while (low <= high) {
        const mid = Math.floor((low + high) / 2)

        if (list[mid] === value) {
            return mid
        } else if (list[mid] > value) {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return -1
}

module.exports = binarySearch