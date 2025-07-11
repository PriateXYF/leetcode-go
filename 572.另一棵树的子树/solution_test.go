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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// 检测是否是子树
	var checkChild func(*TreeNode, *TreeNode) bool
	checkChild = func(tree *TreeNode, sub *TreeNode) bool {
		if (tree != nil && sub == nil) || (tree == nil && sub != nil) {
			return false
		}
		if tree == nil && sub == nil {
			return true
		}
		if tree.Val != sub.Val {
			return checkChild(tree.Left, subRoot) || checkChild(tree.Right, subRoot)
		} else {
			return checkChild(tree.Left, sub.Left) && checkChild(tree.Right, sub.Right)
		}
	}
	// 遍历检测每个节点
	var traverseCheck func(*TreeNode) bool
	traverseCheck = func(tree *TreeNode) bool {
		if tree == nil {
			return false
		}
		res := false
		if tree.Val == subRoot.Val {
			res = checkChild(tree, subRoot)
		}
		return traverseCheck(tree.Left) || traverseCheck(tree.Right) || res
	}

	return traverseCheck(root)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSubtreeOfAnotherTree(t *testing.T) {

}
