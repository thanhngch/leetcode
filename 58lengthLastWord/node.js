/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    let end = s.length;
    let i;
    for (i = s.length - 1; i > -1 ; i--) {
        if (s[i] == ' ') {
            end = i;
        } else {
            break;
        }
    }
    for ( ; i > -1 ; i--) {
        if (s[i] == ' ') {
            break;
        }
    }
    return end - i - 1;
};
console.log(lengthOfLastWord('Hello World')); // 5
console.log(lengthOfLastWord('Hello ')); // 5
console.log(lengthOfLastWord('  Hello  ')); // 5
console.log(lengthOfLastWord('')); // 0
console.log(lengthOfLastWord(' ')); // 0
