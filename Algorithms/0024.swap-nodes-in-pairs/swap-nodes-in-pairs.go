package problem0024

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode is definition for singly-linked list
type ListNode = kit.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}

//better 递归
func swapPairs2(head *ListNode) *ListNode {
	// 第一个节点和第二个节点为空时，直接返回原链表
	if head == nil || head.Next == nil {
		return head
	}

	// 声明firstNode，secondNode，并将head赋给firstNode，head的下一个节点赋给secondNode
	firstNode := head
	secondNode := head.Next

	//交换firstNode与second,
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}

//better 迭代
func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val : -1,
		Next : head,
	}
	//preNode用来记录两个节点交换后前面那个节点的前一个节点，用于最后的return。
	//①head或者head.Next为nil时应该返回head
	//②head或者head.Next不为nil时应该返回节点交换后的前一个节点
	//用preNode就能在最后返回的时候统一以上两种情况，只用dummy.Next即可
	preNode := dummy

	for head != nil && head.Next != nil {
		firstNode := head
		secondNode := head.Next

		//更新preNode的next
		preNode.Next = secondNode
		//交换节点
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		preNode = firstNode
		head = firstNode.Next
	}
	return dummy.Next
}
