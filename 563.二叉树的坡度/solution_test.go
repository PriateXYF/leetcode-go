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
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftTilt, rightTilt := findTilt(root.Left), findTilt(root.Right)
	leftVal, rightVal := 0, 0
	if root.Left != nil {
		leftVal = root.Left.Val
	}
	if root.Right != nil {
		rightVal = root.Right.Val
	}
	// 更新当前节点的值
	root.Val += leftVal + rightVal
	tilt := abs(leftVal - rightVal)
	return tilt + leftTilt + rightTilt
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeTilt(t *testing.T) {

}
