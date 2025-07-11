package leetcode

import (
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findLHS(nums []int) (res int) {
	sort.Ints(nums)
	start := 0
	for i, num := range nums {
		if num-nums[start] > 1 {
			start++
		}
		if num-nums[start] == 1 {
			res = max(res, i-start+1)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestHarmoniousSubsequence(t *testing.T) {

}
