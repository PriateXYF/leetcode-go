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
func getDecimalValue(head *ListNode) int {
	res := 0
	for ; head != nil; head = head.Next {
		res |= head.Val
		res = res << 1
	}
	return res >> 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConvertBinaryNumberInALinkedListToInteger(t *testing.T) {

}
