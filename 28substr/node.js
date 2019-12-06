/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function(haystack, needle) {
    return haystack.indexOf(needle);
};
var main = function() {
    console.log(strStr("123", "3"));
    console.log(strStr("hello", "ll"));
    console.log(strStr("aaaaa", "bba"));
};

main();
