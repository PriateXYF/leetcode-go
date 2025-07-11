package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	// 头指针
	head := 0
	// 尾指针
	tail := 1
	for tail < len(nums) {
		// 如果头指针等于尾指针
		if nums[head] == nums[tail] {
			tail++
		} else {
			head++
			nums[head] = nums[tail]
			tail++
		}
	}
	return head + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {

}
