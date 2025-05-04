package linkedlist

// 206. Reverse Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	currentNode := head
	var previousNode *ListNode
	for currentNode != nil {
		extra := currentNode.Next

		currentNode.Next = previousNode

		currentNode, previousNode = extra, currentNode
	}
	return previousNode
}
