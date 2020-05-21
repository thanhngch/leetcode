/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
  let max = 0;
  let l = 0;
  let r = height.length - 1;
  while (l < r) {
    max = Math.max(max, Math.min(height[l], height[r]) * (r - l));
    if (height[l] < height[r]) {
      l++;
    } else {
      r--;
    }
  }
  return max;
};
// var maxArea = function(height) {
//   let max = 0;
//   let maxHeight = 0;
//   for (let i = height.length - 1; i > -1; i--) {
//     if (maxHeight < height[i]) {
//       maxHeight = height[i];
//       for (let j = 0; j < i; j++) {
//         max = Math.max(max, Math.min(height[i], height[j]) * (i-j));
//       }
//     }
//   }
//   return max;
// }
console.log(maxArea([1,8,6,2,5,4,8,3,1])); // 40
console.log(maxArea([1,8,6,2,5,4,8,3,7])); // 49
console.log(maxArea([1,8,6,2,5,4,8,3])); // 40
console.log(maxArea([1,8,6,2,5,4,8])); // 40
console.log(maxArea([1,8,6,2,5,4])); // 16
console.log(maxArea([1,8,6,2,5])); // 15
console.log(maxArea([1,8,6,2])); // 6
console.log(maxArea([1,8,6])); // 6
console.log(maxArea([1,8])); // 1
console.log(maxArea([1,2,1])); // 2
console.log(maxArea([1,2,4,3])); // 4