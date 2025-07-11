package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	tail := len(nums1) - 1
	m--
	n--
	for tail >= 0 {
		if m < 0 && n >= 0 {
			nums1[tail] = nums2[n]
			n--
		} else if n < 0 && m >= 0 {
			nums1[tail] = nums1[m]
			m--
		} else if nums2[n] > nums1[m] {
			nums1[tail] = nums2[n]
			n--
		} else if nums1[m] >= nums2[n] {
			nums1[tail] = nums1[m]
			m--
		}
		tail--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeSortedArray(t *testing.T) {

}
