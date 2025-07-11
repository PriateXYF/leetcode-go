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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	point1 := list1
	point2 := list2
	// 头结点
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	nodes := head
	for point1 != nil && point2 != nil {
		if point1.Val >= point2.Val {
			nodes.Next = &ListNode{
				Val:  point2.Val,
				Next: nil,
			}
			point2 = point2.Next
		} else {
			nodes.Next = &ListNode{
				Val:  point1.Val,
				Next: nil,
			}
			point1 = point1.Next
		}
		nodes = nodes.Next
	}
	if point1 == nil {
		nodes.Next = point2
	} else if point2 == nil {
		nodes.Next = point1
	}
	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeTwoSortedLists(t *testing.T) {

}
