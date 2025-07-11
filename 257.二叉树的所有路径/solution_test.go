package leetcode

import (
	"strconv"
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
// 生成路径字符
func generatePathStr(nodes []int, tail int) (res string) {
	for _, node := range nodes {
		res += strconv.Itoa(node) + "->"
	}
	res += strconv.Itoa(tail)
	return res
}

// 统计二叉树根节点到叶子节点的所有路径
func binaryTreePaths(root *TreeNode) (res []string) {
	// 先序遍历二叉树
	var traverse func(*TreeNode, []int)
	traverse = func(tree *TreeNode, parents []int) {
		if tree == nil {
			return
		}
		// 如果当前节点是叶子节点
		if tree.Left == nil && tree.Right == nil {
			path := generatePathStr(parents, tree.Val)
			res = append(res, path)
			return
		}
		parents = append(parents, tree.Val)
		traverse(tree.Left, parents)
		traverse(tree.Right, parents)
	}
	traverse(root, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreePaths(t *testing.T) {

}
