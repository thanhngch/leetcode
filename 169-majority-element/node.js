/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function(nums) {
    let major = {};
    let maxElement;
    for (let i = 0; i < nums.length; i++) {
        if (major[nums[i]]) {
            major[nums[i]]++;
        } else {
            major[nums[i]] = 1;
        }
        if (!maxElement || major[maxElement] < major[nums[i]]) {
            maxElement = nums[i];
        }
    }
    return maxElement;
};

console.log(majorityElement([2])); // 2
console.log(majorityElement([3,2,3])); // 3
console.log(majorityElement([2,2,1,1,1,2,2])); // 2