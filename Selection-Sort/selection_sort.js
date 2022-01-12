function findSmallest(array) {
    let minVal = array[0]
    let minIdx = 0

    for (let i = 0; i < array.length; i++) {
        if (array[i] < minVal) {
            minVal = array[i]
            minIdx = i
        }
    }
    return minIdx
}

// Difficulty: log(x^2)
function selectionSortAsc(array) {
    const len = array.length
    const sortedArr = []

    for (let i = 0; i < len; i++) {
        const smallestIdx = findSmallest(array)
        sortedArr.push(...array.splice(smallestIdx, 1))
    }
    return sortedArr
}

module.exports = selectionSortAsc