package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maximumLength(nums []int, k int) (res int) {
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	for _, x := range nums {
		x %= k
		for y := range k {
			dp[x][y] = dp[y][x] + 1
			res = max(dp[x][y], res)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindTheMaximumLengthOfValidSubsequenceIi(t *testing.T) {

}
