/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
    if (nums.length == 0) {
        return 0;
    }
    if (nums.length == 1) {
        return nums[0];
    }
    if (nums.length == 2) {
        return Math.max(nums[0], nums[1]);
    }
    let money = [];
    money[0] = nums[0];
    money[1] = nums[1];
    money[2] = nums[2] + nums[0];
    for (let i = 3; i < nums.length; i++) {
        money[i] = Math.max(money[i-2], money[i-3]) + nums[i];
    }
    return Math.max(money[nums.length-1], money[nums.length-2]);
};

console.log(rob([1,2,3,1])); // 4
console.log(rob([2,1,1,2])); // 4
console.log(rob([2,7,9,3,1])); // 12
