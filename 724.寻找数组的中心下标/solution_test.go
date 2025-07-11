package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	leftSum := 0
	for i, num := range nums {
		if leftSum == sum-leftSum-num {
			return i
		}
		leftSum += num
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindPivotIndex(t *testing.T) {

}
