package main

import "fmt"

// ListNode is list of node
type ListNode struct {
	Val  int
	Next *ListNode
}

func printNode(l *ListNode) {
	var node = l
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func setupNode1() *ListNode {
	var node = &ListNode{Val: 1}
	node.Next = &ListNode{Val: 3}
	node.Next.Next = &ListNode{Val: 5}
	node.Next.Next.Next = &ListNode{Val: 7}
	return node
}

func setupNode2() *ListNode {
	var node = &ListNode{Val: 2}
	node.Next = &ListNode{Val: 4}
	node.Next.Next = &ListNode{Val: 6}
	node.Next.Next.Next = &ListNode{Val: 8}
	return node
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var node = &ListNode{}
	var headNode = node
	for {
		if l1 != nil && (l2 == nil || l1.Val < l2.Val) {
			node.Next = &ListNode{}
			node.Next.Val = l1.Val
			l1 = l1.Next
		} else if l2 != nil {
			node.Next = &ListNode{}
			node.Next.Val = l2.Val
			l2 = l2.Next
		} else {
			break
		}
		node = node.Next
	}
	return headNode.Next
}

func main() {
	var node1 = setupNode1()
	printNode(node1)
	var node2 = setupNode2()
	printNode(node2)

	printNode(mergeTwoLists(node1, node2))
}
