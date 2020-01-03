/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit1 = function(prices) {
    if (prices.length == 0) {
        return 0;
    }
    let max = prices[prices.length - 1];
    let min = prices[0];
    let diff = 0;
    for (let i = prices.length - 1; i > 0; i--) {
        max = Math.max(prices[i], max);
        min = prices[0];
        for (let j = 0; j < i; j++) {
            min = Math.min(prices[j], min);
        }
        diff = Math.max(diff, max - min);
    }
    return diff;
};

/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    if (prices.length == 0) {
        return 0;
    }
    let min = prices[0];
    let diff = 0;
    for (let i = 0; i < prices.length; i++) {
        min = Math.min(prices[i], min);
        diff = Math.max(diff, prices[i] - min);
    }
    return diff;
};
console.log(maxProfit([7,1,5,3,6,4])); // 5
console.log(maxProfit([7,6,4,3,1])); // 0
console.log(maxProfit([1, 7])); // 6
console.log(maxProfit([7,1,7,0,5])); // 6
