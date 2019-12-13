/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
    let remember = false;
    for (let i = digits.length - 1; i > -1; i--) {
        if (digits[i] == 9) {
            remember = true;
            digits[i] = 0;
        } else {
            remember = false;
            digits[i]++;
        }
        if (!remember) {
            return digits;
        }
    }
    return [1].concat(digits);
};

console.log(plusOne([1,2,3])) // 1,2,4
console.log(plusOne([4,3,2,1])) // 4,3,2,2
console.log(plusOne([8,9])) // 9 0
console.log(plusOne([8,9,9])) // 9 0 0
console.log(plusOne([9,9,9])) // 1,0,0,0