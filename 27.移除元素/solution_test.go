package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	head := 0
	tail := 0
	for tail < len(nums) {
		for tail < len(nums) && nums[tail] == val {
			tail++
		}
		if tail < len(nums) {
			nums[head] = nums[tail]
			head++
			tail++
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveElement(t *testing.T) {

}
