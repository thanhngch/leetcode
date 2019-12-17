
function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

 /**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {boolean}
 */
var isSameTree = function(p, q) {
    if (!p && !q) {
        return true;
    }
    if (p && q) {
        return p.val == q.val && isSameTree(p.left, q.left) &&
            isSameTree(p.right, q.right);
    }
    return false;
};

var printNode = function(p) {
    if (p) {
        console.log(p.val);
        printNode(p.left);
        printNode(p.right);
    }
};
var main = function () {
    let t1 = new TreeNode(1);
    t1.left = new TreeNode(2);
    t1.right = new TreeNode(3);

    let t2 = new TreeNode(1);
    t2.left = new TreeNode(2);
    t2.right = new TreeNode(4);
    console.log(isSameTree(t1, t2));

    // printNode(t);
}

main();
