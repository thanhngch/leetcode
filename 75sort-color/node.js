/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var sortColors = function(nums) {
  let num1 = 0;
  let num2 = 0;
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] === 0) {
      num1++;
    }
    if (nums[i] === 1) {
      num2++;
    }
  }
  for (let i = 0; i < num1; i++) {
    nums[i] = 0;
  }
  for (let i = num1; i < num1 + num2; i++) {
    nums[i] = 1;
  }
  for (let i = num1 + num2; i < nums.length; i++) {
    nums[i] = 2;
  }
};

var colors = [2,0,2,1,1,0];
sortColors(colors);
console.log(colors);
