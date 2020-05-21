/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsets = function(nums) {
  let ans = [];
  for (let i = 0; i < nums.length; i++) {
    let ansSub = [];
    for (let j = i; j < nums.length; j++) {
      ansSub.push(nums[j]);
      // ans.push(ansSub);
      console.log(ansSub);
    }
  }
  return ans;
};

console.log(subsets([1,2,3]));