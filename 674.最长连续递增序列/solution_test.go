package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findLengthOfLCIS(nums []int) (res int) {
	if len(nums) == 1 {
		return 1
	}
	for seq, i := 1, 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			seq++
		} else {
			seq = 1
		}
		res = max(res, seq)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestContinuousIncreasingSubsequence(t *testing.T) {

}
