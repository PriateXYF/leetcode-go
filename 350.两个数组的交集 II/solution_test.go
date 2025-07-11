package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) (res []int) {
	table := map[int]int{}
	for _, num := range nums1 {
		if _, exists := table[num]; exists {
			table[num]++
		} else {
			table[num] = 1
		}
	}
	for _, num := range nums2 {
		if times, exists := table[num]; exists && times > 0 {
			res = append(res, num)
			table[num]--
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIntersectionOfTwoArraysIi(t *testing.T) {

}
