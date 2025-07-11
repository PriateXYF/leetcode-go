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
func sumOfLeftLeaves(root *TreeNode) (res int) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}
	res += sumOfLeftLeaves(root.Left)
	res += sumOfLeftLeaves(root.Right)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSumOfLeftLeaves(t *testing.T) {

}
