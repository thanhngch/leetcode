/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
    let i2 = 0;
    let arr = [];
    let iArr = 0;
    if (n == 0) {
        return nums1;
    }
    for (let i1 = 0; i1 < m + n; i1++) {
        if (iArr == arr.length && nums1[i1] <= nums2[i2] && i1 < m) {
            continue;
        } else {
            if (i1 < m) {
                arr.push(nums1[i1]);
            }
            if (i2 >= n || iArr < arr.length && arr[iArr] < nums2[i2]) {
                nums1[i1] = arr[iArr];
                iArr++;
            } else {
                nums1[i1] = nums2[i2];
                i2++;
            }
        }
    }
    return nums1;
};
// console.log(merge([1,2,3,0,0,0],3,[2,5,6],3)); // [ 1, 2, 2, 3, 5, 6 ]
// console.log(merge([1],1,[],0));
console.log(merge([2,0],1,[1],1));
