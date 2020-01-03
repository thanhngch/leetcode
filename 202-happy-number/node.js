/**
 * @param {number} n
 * @return {number}
 */
var nextNumber = function(n) {
    let sum = 0;
    while (n != 0) {
        sum += (n % 10) ** 2;
        n = Number.parseInt(n / 10);
    }
    return sum;
};

/**
 * @param {number} n
 * @return {boolean}
 */
var isHappy1 = function(n) {
    let map = {};
    while (n != 1) {
        n = nextNumber(n);
        if (map[n]) {
            return false;
        }
        map[n] = true;
    }
    return true;
};
/**
 * @param {number} n
 * @return {boolean}
 */
var isHappy = function(n) {
    let map = [];
    while (n != 1) {
        n = nextNumber(n);
        if (map.indexOf(n) > -1) {
            return false;
        }
        map.push(n);
    }
    return true;
};
var main = function() {
    // console.log(isHappy(19)); // true
    // console.log(isHappy(4)); // true
    let arr = [];
    for (let i = 0; i < 100; i++) {
        if (isHappy(i)) {
            arr.push(i);
        }
    }
    console.log(arr);
};

main();
