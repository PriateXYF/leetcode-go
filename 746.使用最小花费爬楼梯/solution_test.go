package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	pre, cur := 0, 0
	for i := 2; i <= len(cost); i++ {
		// cur 是前一个， per 是前两个
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinCostClimbingStairs(t *testing.T) {

}
