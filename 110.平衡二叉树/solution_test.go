package leetcode

import (
	"math"
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
// 计算二叉树的深度
func depth(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	return max(depth(tree.Right), depth(tree.Left)) + 1
}

// 判断一颗树是否为平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right == nil && root.Left == nil {
		return true
	}
	if root.Right == nil || root.Left == nil {
		return (depth(root.Right) + depth(root.Left)) <= 1
	}
	if math.Abs(float64(depth(root.Right)-depth(root.Left))) > 1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBalancedBinaryTree(t *testing.T) {

}
