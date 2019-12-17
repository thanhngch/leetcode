function ListNode(val) {
    this.val = val;
    this.next = null;
}
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function(head) {
    let node = head;
    while (node) {
        if (node.next && node.val == node.next.val) {
            node.next = node.next.next;
        } else  {
            node = node.next;
        }
    }
    return head;
};

var printNode = function(node) {
    let data = [];
    while(node) {
        data.push(node.val);
        node = node.next;
    }
    console.log(data.join(' -> '));
};

var main = function() {
    let node = new ListNode(1);
    node.next = new ListNode(1);
    node.next.next = new ListNode(1);
    node.next.next.next = new ListNode(1);
    node.next.next.next.next = new ListNode(3);
    node.next.next.next.next.next = new ListNode(3);
    node = deleteDuplicates(node);
    printNode(node);
};

main();
