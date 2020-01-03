/**
 * Definition for singly-linked list.
 */ 
function ListNode(val) {
    this.val = val;
    this.next = null;
}

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
    let map = {};
    while (head) {
        if (map[head.val]) {
            for (let i = 0; i < map[head.val].length; i++) {
                if (map[head.val][i] === head) {
                    return true;
                }
            }
        }
        if (!map[head.val]) {
            map[head.val] = [head];
        } else {
            map[head.val].push(head);
        }
        head = head.next;
    }
    return false;
};

var main = function() {
    let l1 = new ListNode(3);
    let l2 = new ListNode(2);
    let l3 = new ListNode(0);
    let l4 = new ListNode(-4);
    l1.next = l2;
    l2.next = l3;
    l3.next = l4;
    l4.next = l2;

    console.log(hasCycle(l1));
};

main();
