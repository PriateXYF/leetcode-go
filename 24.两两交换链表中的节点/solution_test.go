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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	head = second
	for second != nil {
		first.Next = second.Next
		second.Next = first
		first, second = second, first
		if second.Next == nil || second.Next.Next == nil {
			break
		} else {
			first = first.Next.Next
			second.Next = second.Next.Next
			second = second.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSwapNodesInPairs(t *testing.T) {

}
