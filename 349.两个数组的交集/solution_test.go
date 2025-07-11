package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) (res []int) {
	// 桶排序
	bucket := make([]int, 1001)
	for _, num := range nums1 {
		bucket[num]++
	}
	for _, num := range nums2 {
		if bucket[num] == -1 {
			continue
		}
		if bucket[num] > 0 {
			res = append(res, num)
			bucket[num] = -1
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIntersectionOfTwoArrays(t *testing.T) {

}
