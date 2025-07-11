package leetcode

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func summaryRanges(nums []int) (res []string) {
	start := 0
	for i := 0; i < len(nums); i++ {
		// 如果是最后一位或者碰到中断
		if i == len(nums)-1 || nums[i+1]-nums[i] > 1 {
			if nums[start] == nums[i] {
				res = append(res, fmt.Sprintf("%d", nums[i]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
			}
			start = i + 1
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSummaryRanges(t *testing.T) {

}
