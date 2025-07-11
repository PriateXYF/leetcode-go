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

func getListLength(head *ListNode) (len int) {
	for node := head; node != nil; node = node.Next {
		len++
	}
	return len
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := getListLength(headA), getListLength(headB)
	nodeA, nodeB := headA, headB
	// 需要忽略的节点数量
	skip := lenA - lenB
	// 如果 A 比较长
	if skip > 0 {
		for skip > 0 {
			nodeA = nodeA.Next
			skip--
		}
	} else { // 如果 B 比较长
		for skip < 0 {
			nodeB = nodeB.Next
			skip++
		}
	}
	if nodeA == nodeB {
		return nodeA
	}
	// 开始比较 A 和 B
	for nodeA != nil {
		if nodeA.Next == nodeB.Next && nodeA.Next != nil {
			return nodeA.Next
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIntersectionOfTwoLinkedLists(t *testing.T) {

}
