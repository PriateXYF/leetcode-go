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

func postorder(root *Node) (res []int) {
	if root == nil {
		return []int{}
	}
	for _, child := range root.Children {
		seq := postorder(child)
		res = append(res, seq...)
	}
	res = append(res, root.Val)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNAryTreePostorderTraversal(t *testing.T) {

}
