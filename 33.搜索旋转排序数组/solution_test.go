package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for mid := (left + right) / 2; left <= right; mid = (left + right) / 2 {
		if nums[mid] == target {
			return mid
		}
		if nums[right] < nums[left] { // 进入的是异常空间
			if nums[mid] < nums[left] { // 异常空间在 left - mid，有效范围是 mid+1 - right
				if target < nums[mid] || target > nums[right] { // 目标在异常空间
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else { // 异常空间在 mid+1 - right，有效范围是 left - mid
				if target < nums[mid] && target >= nums[left] { // 目标在正常
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		} else { // 进入的是正常空间
			if target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInRotatedSortedArray(t *testing.T) {

}
