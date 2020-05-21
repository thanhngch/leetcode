/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function (a, b) {
    let result = '';
    let remember = 0;
    const charToInt = {
        '1': 1,
        '0': 0,
    };
    for (let i = a.length - 1, j = b.length - 1; 
        i > -1 || j > -1; i--, j--) {
        let ai = '0';
        if (i > -1) {
            ai = a[i];
        }
        let bj = '0';
        if (j > -1) {
            bj = b[j];
        }
        let sum = charToInt[ai] + charToInt[bj] + remember;
        switch (sum) {
            case 0:
                result = '0' + result;
                remember = 0;
                break;
            case 1:
                result = '1' + result;
                remember = 0;
                break;
            case 2:
                result = '0' + result;
                remember = 1;
                break;
            case 3:
                result = '1' + result;
                remember = 1;
                break;
        }
    }
    if (remember == 0) {
        return result;
    }
    return '1' + result;
};

var main = function () {
    console.log(addBinary('11', '100')); // 111
    console.log(addBinary('11', '11')); // 110
};

main();