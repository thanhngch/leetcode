function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

/**
 * @param {TreeNode} root
 * @return {number}
 */
var maxDepth = function(root) {
    if (root == null) {
        return 0;
    }
    return 1 + Math.max(maxDepth(root.left), maxDepth(root.right));
};

var main = function() {
  let root = new TreeNode(1);
  root.left = new TreeNode(2);
  root.right = new TreeNode(2);
  root.left.left = new TreeNode(4);
  root.left.right = new TreeNode(3);
  root.right.left = new TreeNode(3);
  root.right.right = new TreeNode(4);
  console.log(maxDepth(root));
};

main();