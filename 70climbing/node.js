/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    let a = 1;
    let b = 0;
    for (let i = 0; i < n; i++) {
        a = a + b;
        b = a - b;
    }
    return a;
};

var main = function() {
    console.log(climbStairs(2));
    console.log(climbStairs(3));
    console.log(climbStairs(4));
};

main();
