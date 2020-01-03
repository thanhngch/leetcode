function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

/**
 * @param {TreeNode} root
 * @return {number}
 */
var minDepth = function (root) {
    if (root == null) {
        return 0;
    }
    if (root.left == null && root.right == null) {
        return 1;
    }
    if (root.left && root.right == null) {
        return minDepth(root.left) + 1;
    }
    if (root.left == null && root.right) {
        return minDepth(root.right) + 1;
    }
    return Math.min(minDepth(root.left) + 1, minDepth(root.right) + 1);
};

var main = function () {
    let t1 = new TreeNode(1);
    t1.left = new TreeNode(2);
    // t1.right = new TreeNode(3);
    console.log(minDepth(t1));

    // printNode(t);
}

main();
