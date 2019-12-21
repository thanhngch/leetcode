function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var levelOrderBottom = function(root) {
    let arr = [];
    let traversal = function(arr, level, node) {
        if (node) {
            traversal(arr, level + 1, node.left);
            traversal(arr, level + 1, node.right);
            if (node.left || node.right) {
                let subArr = [];
                node.left && subArr.push(node.left.val);
                node.right && subArr.push(node.right.val);
                if (!arr[level]) {
                    arr[level] = [];
                }
                arr[level] = arr[level].concat(subArr);
            }
        }
    };
    traversal(arr, 0, root);
    arr.reverse();
    if (root) {
        arr.push([root.val]);
    }
    return arr;
};

var main = function() {
    let root = new TreeNode(3);
    root.left = new TreeNode(9);
    root.right = new TreeNode(20);
    root.right.left = new TreeNode(15);
    root.right.right = new TreeNode(7);
    console.log(levelOrderBottom(root));
};

main();

//     1
//    2 3
//   4 n n 5
