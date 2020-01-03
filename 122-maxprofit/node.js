/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    if (prices.length < 2) {
        return 0;
    }
    let profix = 0;
    let min = prices[0];
    let i = 0;
    while (i < prices.length - 1) {
        for (; i < prices.length - 1; i++) {
            if (prices[i] >= prices[i+1]) {
                profix += Math.max(prices[i] - min, 0);
                min = prices[i+1];
                i++;
                break;
            }
        }
    }
    if (prices[prices.length - 2] < prices[prices.length - 1]) {
        profix += Math.max(prices[prices.length - 1] - min, 0);
    }
    return profix;
};

console.log(maxProfit([7,1,5,3,6,4])); // 7
console.log(maxProfit([1,2,3,4,5])); // 4
console.log(maxProfit([7,6,4,3,2,1])); // 0
console.log(maxProfit([])); // 0
console.log(maxProfit([1])); // 0
console.log(maxProfit([1,2])); // 1
console.log(maxProfit([2,1])); // 0