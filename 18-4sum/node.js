var checkDuplicate = function(ans, item) {
  for (let i = 0; i < ans.length; i++) {
    const a = ans[i];
    if (a[0] == item[0] && a[1] == item[1] && a[2] == item[2] && a[3] == item[3]) {
      return true;
    }
  }
  return false;
}
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[][]}
 */
var fourSum = function(nums, target) {
  const ans = [];
  for(let i = 0; i < nums.length - 3; i++) {
    for(let j = i+1; j < nums.length - 2; j++) {
      for(let k = j+1; k < nums.length - 1; k++) {
        for(let l = k+1; l < nums.length; l++) {
          if (nums[i] + nums[j] + nums[k] + nums[l] == target) {
            let item = [nums[i], nums[j], nums[k], nums[l]].sort((a,b) => a - b);
            if(!checkDuplicate(ans, item)) {
              ans.push(item);
            }
          }
        }
      }
    }
  }
  return ans;
};
let ans = fourSum([-3,-2,-1,0,0,1,2,3], 0);
console.log(ans, ans.length);
