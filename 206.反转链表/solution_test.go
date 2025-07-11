package leetcode

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 反转链表（迭代法）
func reverseList2(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 反转链表（递归法）
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		// 这个实际上就是 tail
		return head
	}
	tail := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tail
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedList(t *testing.T) {

}
