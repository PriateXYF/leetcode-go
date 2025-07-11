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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nodes := head
	for nodes.Next != nil {
		if nodes.Val == nodes.Next.Val {
			nodes.Next = nodes.Next.Next
			continue
		}
		nodes = nodes.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {

}
