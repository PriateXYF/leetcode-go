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
func diameterOfBinaryTree(root *TreeNode) (maxDiameter int) {
	var calcDiameter func(*TreeNode) int
	calcDiameter = func(node *TreeNode) (depth int) {
		if node == nil {
			return 0
		}
		leftDepth, rightDepth := calcDiameter(node.Left), calcDiameter(node.Right)
		if node.Left != nil {
			leftDepth += 1
		}
		if node.Right != nil {
			rightDepth += 1
		}
		maxDiameter = max(leftDepth+rightDepth, maxDiameter)
		return max(leftDepth, rightDepth)
	}
	calcDiameter(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDiameterOfBinaryTree(t *testing.T) {

}
