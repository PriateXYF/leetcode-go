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
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	// 先找到不为指定值的头结点
	for head.Val == val {
		head = head.Next
		if head == nil {
			return nil
		}
	}
	node := head
	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveLinkedListElements(t *testing.T) {

}
