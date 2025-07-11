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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	res := head
	var val = func(n *ListNode) (int, *ListNode) {
		if n == nil {
			return 0, nil
		}
		return n.Val, n.Next
	}
	var addNode = func(n *ListNode, val int) *ListNode {
		var node ListNode
		node.Val = val
		n.Next = &node
		return n.Next
	}
	for carry, ans := 0, 0; l1 != nil || l2 != nil; {
		var v1, v2 int
		v1, l1 = val(l1)
		v2, l2 = val(l2)
		value := v1 + v2 + carry
		carry, ans = value/10, value%10
		res = addNode(res, ans)
		if l1 == nil && l2 == nil && carry == 1 {
			res = addNode(res, carry)
		}
	}
	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{1, nil}
	l2 := ListNode{1, nil}
	addTwoNumbers(&l1, &l2)
}
