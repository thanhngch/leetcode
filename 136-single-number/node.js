/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
    return nums.reduce((acc, value) => acc ^ value, 0);
};

console.log(singleNumber([4,1,2,1,2])); // 4
console.log(singleNumber([2,2,1])); // 1
