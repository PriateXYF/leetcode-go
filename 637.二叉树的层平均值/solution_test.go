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
func averageOfLevels(root *TreeNode) (res []float64) {
	queue := []*TreeNode{root}
	head := 0
	level, carry, nextCarry, sum := 1, 1, 0, 0.0
	// 层序遍历二叉树
	for head < len(queue) {
		// 取出头节点
		node := queue[head]
		sum += float64(node.Val)
		if node.Left != nil {
			nextCarry++
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			nextCarry++
			queue = append(queue, node.Right)
		}
		head, carry = head+1, carry-1
		if carry == 0 {
			res = append(res, sum/float64(level))
			sum, level, carry, nextCarry = 0, nextCarry, nextCarry, 0
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAverageOfLevelsInBinaryTree(t *testing.T) {

}
