package leetcode

import (
	"reflect"
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
func dfs(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
	}
	dfs(root.Left, leaves)
	dfs(root.Right, leaves)
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaves1, leaves2 []int
	dfs(root1, &leaves1)
	dfs(root2, &leaves2)
	return reflect.DeepEqual(leaves1, leaves2)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLeafSimilarTrees(t *testing.T) {

}
