package leetcode

import (
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) (res []int) {
	if root == nil {
		return []int{}
	}
	res = append(res, root.Val)
	for _, child := range root.Children {
		seq := preorder(child)
		res = append(res, seq...)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNAryTreePreorderTraversal(t *testing.T) {

}
