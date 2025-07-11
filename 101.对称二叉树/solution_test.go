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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isSame func(*TreeNode, *TreeNode) bool
	isSame = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return isSame(left.Left, right.Right) && isSame(left.Right, right.Left)
	}
	// 如果左右子树都存在
	if root.Right != nil && root.Left != nil {
		return isSame(root.Left, root.Right)
	} else if root.Right != nil || root.Left != nil {
		return false
	} else {
		return true
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSymmetricTree(t *testing.T) {

}
