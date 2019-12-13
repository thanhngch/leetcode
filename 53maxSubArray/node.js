/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
    if (nums.length == 0) {
        return 0;
    }
    let maxNumsPrev;
    let maxNumsNext = nums[0];
    let maxNumber = nums[0];
    for (let i = 1; i < nums.length; i++) {
        maxNumsPrev = maxNumsNext;
        maxNumsNext = Math.max(nums[i], maxNumsPrev + nums[i]);
        maxNumber = Math.max(maxNumber, maxNumsNext);
    }
    return maxNumber;
};
console.log(maxSubArray([-2,1,-3,4,-1,2,1,-5,4]));