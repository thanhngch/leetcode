var isAlplaNumberic = function(c) {
    if ((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')) {
        return true;
    }
    return false;
}
/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
    let i = 0;
    let j = s.length - 1;
    while (i < j) {
        let si = s[i].toLowerCase();
        let sj = s[j].toLowerCase();

        while (!(isAlplaNumberic(si))) {
            i++;
            si = s[i].toLowerCase();
            if (i >= j) {
                break;
            }
        }
        while (!(isAlplaNumberic(sj))) {
            j--;
            sj = s[j].toLowerCase();
            if (i >= j) {
                break;
            }
        }
        if (si === sj) {
            i++;
            j--;
        } else if (isAlplaNumberic(si)) {
            return false;
        }
    }
    return true;
};

// console.log(isPalindrome('A man, a plan, a canal: Panama')); // true
// console.log(isPalindrome('race a car')); // false
// console.log(isPalindrome('.,')); // true
console.log(isPalindrome('0P')); // false
// A man, a plan, a canal: Panama
