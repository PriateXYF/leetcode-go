package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	head := 0
	tail := len(nums) - 1
	if target <= nums[head] {
		return head
	} else if target > nums[tail] {
		return tail + 1
	}
	var mid int
	for true {
		mid = (head + tail) / 2
		if nums[mid] > target {
			tail = mid
		} else if nums[mid] < target {
			head = mid
		} else {
			return mid
		}
		if head == tail || tail-head == 1 {
			return tail
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInsertPosition(t *testing.T) {

}
