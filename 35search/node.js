/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function(nums, target) {
    if (nums[0] > target) {
        return 0;
    }
    if (nums[nums.length - 1] < target) {
        return nums.length;
    }
    return search(nums, target, 0, nums.length);
};

var search = function(nums, target, start, end) {
    if (start == end) {
        return start;
    }
    let middleIndex = Number.parseInt((start + end) / 2);
    if (nums[middleIndex] == target) {
        return middleIndex;
    } else if (nums[middleIndex] < target) {
        return search(nums, target, middleIndex + 1, end);
    } else if (nums[middleIndex] > target) {
        return search(nums, target, start, middleIndex);
    }
}

console.log(searchInsert([1,3,5,6], 5)); // 2
console.log(searchInsert([1,3,5,6], 2)); // 1
console.log(searchInsert([1,3,5,6], 7)); // 4
console.log(searchInsert([1,3,5,6], 0)); // 0
console.log(searchInsert([1,3,5,6], 4)); // 2
console.log(searchInsert([1,3,5,6], 6)); // 3
