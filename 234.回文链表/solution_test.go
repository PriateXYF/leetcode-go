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
func isPalindrome(head *ListNode) bool {
	length := getListLength(head)
	var prev *ListNode
	// 逆转前半部分
	for range length / 2 {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	// 如果是奇数个元素
	if length%2 == 1 {
		head = head.Next
	}
	for head != nil && prev != nil {
		if head.Val != prev.Val {
			// TODO : 恢复链表
			return false
		}
		head, prev = head.Next, prev.Next
	}
	// TODO : 恢复链表
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPalindromeLinkedList(t *testing.T) {

}
