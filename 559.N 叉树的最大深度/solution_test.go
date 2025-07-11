package leetcode

import (
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 计算 N 叉数的最大深度
func maxDepthNTree(root *Node) (depth int) {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	for _, node := range root.Children {
		depth = max(maxDepthNTree(node)+1, depth)
	}
	return depth
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumDepthOfNAryTree(t *testing.T) {

}
