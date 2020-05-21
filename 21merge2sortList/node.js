function ListNode(val) {
    this.val = val;
    this.next = null;
}

var printNode = function(l) {
    var node = l;
    var nodePrint = [];
    while (node != null) {
        nodePrint.push(node.val + ' -> ');
        node = node.next;
    }
    console.log(nodePrint.join('') + 'null');
}

var setupNode1 = function() {
    let node = new ListNode(1);
    node.next = new ListNode(3);
    node.next.next = new ListNode(5);
    node.next.next.next = new ListNode(7);
    return node;
}

var setupNode2 = function() {
    let node = new ListNode(2);
    node.next = new ListNode(4);
    node.next.next = new ListNode(6);
    node.next.next.next = new ListNode(8);
    return node;
}

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {
    var node = new ListNode();
    var headNode = node;
    while (true) {
        if (l1 != null && (l2 == null || l1.val < l2.val)) {
            node.next = new ListNode();
            node.next.val = l1.val;
            l1 = l1.next;
        } else if (l2 != null) {
            node.next = new ListNode();
            node.next.val = l2.val;
            l2 = l2.next;
        } else {
            break;
        }
        node = node.next;
    }
    return headNode.next;
};


var main = function() {
    let node1 = setupNode1();
    printNode(node1);

    let node2 = setupNode2();
    printNode(node2);

    printNode(mergeTwoLists(node1, node2));
};

main();
