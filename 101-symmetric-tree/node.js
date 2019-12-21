/**
 * Definition for a binary tree node.
 */ 
function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}
var treeNodeToArr = function(arr, roots) {
    if (roots.length > 0) {
        let arrRoot = [];
        for (let i = 0; i < roots.length; i++) {
            if (roots[i]) {
                arr.push(roots[i].val);
                if (roots[i].left || roots[i].right) {
                    arrRoot.push(roots[i].left, roots[i].right);
                }
            } else {
                arr.push(null);
            }
        }
        treeNodeToArr(arr, arrRoot);
    }
}
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
// var isSymmetric = function(root) {
//     let arr = [];
//     treeNodeToArr(arr, [root]);
//     for (let i = 1; i <= arr.length; i = i * 2) {
//         let index = i - 1;
//         for (let j = index; j <= index * 3 / 2; j++) {
//             if (arr[j] !== arr[index*3 - j]) {
//                 return false;
//             }
//         }
//     }
//     return true;
// };

var isSymmetric = function(root) {
    let isEqual = function(node1, node2) {
        if (node1 == null && node2 == null) {
            return true;
        }
        if (node1 != null && node2 != null) {
            return node1.val == node2.val && isEqual(node1.left, node2.right) 
                && isEqual(node1.right, node2.left);
        }
        return false;
    };
    if (root == null) {
        return true;
    }
    return isEqual(root.left, root.right);
};

var main = function() {
    // console.log(isSymmetric([1,2,2,3,4,4,3]));
    // console.log(isSymmetric([1,2,2,null,3,null,3]));
    let root = new TreeNode(1);
    root.left = new TreeNode(2);
    root.right = new TreeNode(2);
    root.left.left = new TreeNode(4);
    root.left.right = new TreeNode(3);
    root.right.left = new TreeNode(3);
    root.right.right = new TreeNode(4);
    // let arr = [];
    // treeNodeToArr(arr, [root]);
    console.log(isSymmetric(root));
    // console.log(arr);
};

main();
// 0 1 2 3 4 5 6 7 8
// 1 2 4 8