package leetcode

import (
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxAverage(nums []int, k int) float64 {
	head, tail, sum, maxSum := 0, k-1, 0, math.MinInt
	// 设置滑动窗口
	for i := head; i <= tail; i++ {
		sum += nums[i]
	}
	maxSum = sum
	for head, tail = head+1, tail+1; tail < len(nums); head, tail = head+1, tail+1 {
		sum = sum - nums[head-1] + nums[tail]
		maxSum = max(sum, maxSum)
	}
	return float64(maxSum) / float64(k)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumAverageSubarrayI(t *testing.T) {

}
