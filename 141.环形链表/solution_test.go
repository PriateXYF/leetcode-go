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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	//	创建一个辅助节点
	helpNode := &ListNode{
		Val:  -999,
		Next: nil,
	}
	for head.Next != nil {
		// 如果有节点的 Next 指向了辅助节点,说明回到了原来的节点上,有环
		if head.Next.Next == helpNode || head.Next == helpNode {
			return true
		}
		nextNode := head.Next
		// 将已经遍历过的节点 Next 指向辅助节点
		head.Next = helpNode
		head = nextNode
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLinkedListCycle(t *testing.T) {

}
