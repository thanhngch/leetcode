/**
 * @param {string[]} words
 * @return {string[]}
 */
var findWords = function (words) {
  const rows = [
    'qwertyuiopQWERTYUIOP',
    'asdfghjklASDFGHJKL',
    'zxcvbnmZXCVBNM'
  ];
  const map = {};
  for (let i = 0; i < rows[0].length; i++) {
    rows[0][i] && (map[rows[0][i]] = 0);
    rows[1][i] && (map[rows[1][i]] = 1);
    rows[2][i] && (map[rows[2][i]] = 2);
  }
  const result = [];
  for (let i = 0; i < words.length; i++) {
    let index = map[words[i][0]];
    let inserted = true;
    for(let j = 1; j < words[i].length;j++) {
      if (map[words[i][j]] != index) {
        inserted = false;
        break;
      }
    }
    if (inserted && index !== undefined) {
      result.push(words[i]);
    }
  }
  return result;
};
console.log(findWords(["Aasdfghjkl","Qwertyuiop","zZxcvbnm"]));
