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
func findSecondMinimumValue(root *TreeNode) int {
	res, rootVal := -1, root.Val
	var dfs func(*TreeNode)
	dfs = func(tree *TreeNode) {
		if tree == nil || (tree.Val > res && res != -1) {
			return
		}
		if tree.Val > rootVal {
			res = tree.Val
		}
		dfs(tree.Left)
		dfs(tree.Right)
	}
	dfs(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSecondMinimumNodeInABinaryTree(t *testing.T) {

}
