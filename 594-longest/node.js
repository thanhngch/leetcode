var lastIndexOf = function(arr) {
  const map = {};
  for(let i = arr.length - 1; i > -1; i--) {
    if (!map[arr[i]]) {
      map[arr[i]] = i;
    }
  }
  return map;
};

/**
 * @param {number[]} nums
 * @return {number}
 */
var findLHS = function (nums) {
  const numsSoft = nums.sort((a,b) => a - b);
  const mapLastIndex = lastIndexOf(numsSoft);
  let maxNumber = 0;
  for (let i = 0; i < numsSoft.length; i++) {
    let indexNext = mapLastIndex[numsSoft[i] + 1] || -1;
    if (maxNumber < indexNext - i + 1) {
      maxNumber = indexNext - i + 1;
    }
    i = mapLastIndex[numsSoft[i]];
  }
  return maxNumber;
};

var findLHS2 = function (nums) {
  let map = new Map();
  for (let i = 0; i < nums.length; i++) {
    if(map.has(nums[i])) {
      map.set(nums[i], map.get(nums[i]) + 1);
    } else {
      map.set(nums[i], 1);
    }
  }
  let maxNumber = 0;
  map.forEach((value, key) => {
    if (map.has(key + 1)) {
      maxNumber = Math.max(map.get(key) + map.get(key + 1), maxNumber);
    }
  });
  return maxNumber;
}

console.log(findLHS2([1,2,3,3,1,-14,13,4])); // 3
console.log(findLHS2([1,3,2,2,5,2,3,7])); // 5
console.log(findLHS2([2,2,2,2])); // 0
console.log(findLHS2([1,2,2,1])); // 4
