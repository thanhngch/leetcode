var findNumber = function(numbers, target, start, end) {
    if (start == end) {
        return -1;
    }
    let middleIndex = Number.parseInt((start + end) / 2);
    if (numbers[middleIndex] == target) {
        return middleIndex;
    } else if (numbers[middleIndex] < target) {
        return findNumber(numbers, target, middleIndex + 1, end);
    } else if (numbers[middleIndex] > target) {
        return findNumber(numbers, target, start, middleIndex);
    }
}
/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum2 = function(numbers, target) {
    for (let i = 0; i < numbers.length; i++) {
        if (numbers[i] > target) {
            return [];
        }
        let findIndex = findNumber(numbers, target - numbers[i], i, numbers.length);
        if (findIndex != -1) {
            return [i, findIndex];
        }
    }
    return [];
};
/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(numbers, target) {
    var hashMap = {};
    for (var i = 0; i < numbers.length; i++) {
        var find = target - numbers[i];
        if (hashMap[find] != undefined) {
            return [hashMap[find] + 1, i + 1];
        } else {
            hashMap[numbers[i]] = i;
        }
    }
    return [];
};

var main = function() {
    console.log(twoSum([2,7,11,15], 9));
};

main();
