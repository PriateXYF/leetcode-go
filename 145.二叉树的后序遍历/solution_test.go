package leetcode

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}
	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root.Left != nil {
			traverse(root.Left)
		}
		if root.Right != nil {
			traverse(root.Right)
		}
		res = append(res, root.Val)
	}
	traverse(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreePostorderTraversal(t *testing.T) {

}
