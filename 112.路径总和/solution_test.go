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
func hasPathSum(root *TreeNode, targetSum int) (res bool) {
	var sum func(*TreeNode, int) int
	// 中顺遍历计算每个叶子节点的路径和
	sum = func(node *TreeNode, total int) int {
		if res {
			return total
		}
		if node == nil {
			return total
		}
		total += node.Val
		total = sum(node.Left, total)
		total = sum(node.Right, total)
		// 判断是否是叶子节点
		if node.Right == nil && node.Left == nil {
			if total == targetSum {
				res = true
			}
		}
		total -= node.Val
		return total
	}
	sum(root, 0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPathSum(t *testing.T) {

}
