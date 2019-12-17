package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ListNode is list of node
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil {
		if node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}
func printNode(node *ListNode) {
	data := []string{}
	for node != nil {
		data = append(data, strconv.Itoa(node.Val))
		node = node.Next
	}
	fmt.Printf("%v", strings.Join(data, " -> "))
}

func main() {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 1}
	node.Next.Next = &ListNode{Val: 1}
	node.Next.Next.Next = &ListNode{Val: 2}
	node.Next.Next.Next.Next = &ListNode{Val: 2}
	node = deleteDuplicates(node)
	printNode(node)
	// fmt.Println("Hello")
}
