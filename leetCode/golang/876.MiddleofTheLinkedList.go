package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	return list[len(list)/2]
}

// sử dụng phương pháp đôi chạy
func middleNode1(head *ListNode) *ListNode {

	middle := head
	end := head
	for end != nil && end.Next != nil {
		fmt.Println(middle.Val)
		middle = middle.Next
		end = end.Next.Next
	}
	return middle
}

func printFromMiddleToEnd(head *ListNode) {
	middle := middleNode(head)
	current := middle

	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}
func main() {
	// Tạo một danh sách ví dụ
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	printFromMiddleToEnd(node1)

}
