/**
 * @param {string} s
 * @param {string} t
 * @return {character}
 */
var findTheDifference = function(s, t) {
    const LIMIT_CHAR_LENGTH = 26;
    const A_CHAR_CODE = 97;
    var aToZ = Array(LIMIT_CHAR_LENGTH).fill(0);

    for (let i = 0; i < t.length; i++) {
        let index = t[i].charCodeAt(0) - A_CHAR_CODE;
        aToZ[index]++;
    }
    
    for (let i = 0; i < s.length; i++) {
        let index = s[i].charCodeAt(0) - A_CHAR_CODE;
        aToZ[index]--;
    }

    for (let i = 0; i < LIMIT_CHAR_LENGTH; i++) {
        if (aToZ[i] === 1) {
            return String.fromCharCode(i + A_CHAR_CODE);
        }
    }
    throw new Error('Not found');
};
var findTheDifference2 = function(s, t) {
    const A_CHAR_CODE = 97;
    let char = 0;
    for (let i = 0; i < s.length; i++) {
        let sIndex = s[i].charCodeAt(0) - A_CHAR_CODE;
        let tIndex = t[i].charCodeAt(0) - A_CHAR_CODE;
        char = char ^ sIndex ^ tIndex;
    }
    char = char ^ (t[t.length - 1].charCodeAt(0) - A_CHAR_CODE);
    return String.fromCharCode(char + A_CHAR_CODE);
};

console.log(findTheDifference2('abcd', 'zabcd'));