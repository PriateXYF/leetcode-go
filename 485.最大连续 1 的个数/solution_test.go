package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxConsecutiveOnes(nums []int) (maxTimes int) {
	tempTimes := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 1 {
			tempTimes++
			maxTimes = max(tempTimes, maxTimes)
		} else if nums[i] == 0 {
			tempTimes = 0
		}
	}
	if tempTimes == maxTimes && nums[len(nums)-1] == 1 {
		maxTimes++
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaxConsecutiveOnes(t *testing.T) {

}
