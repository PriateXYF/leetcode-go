package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(nums1 []int, nums2 []int) (res []int) {
	table := make(map[int]int)
	stack, head := make([]int, len(nums2)), -1
	// 对 nums2 建立查询表
	for i := 0; i < len(nums2); i++ {
		// 出栈,并添加索引
		for head != -1 && nums2[i] > stack[head] {
			table[stack[head]] = nums2[i]
			head--
		}
		head++
		stack[head] = nums2[i]
	}
	// 全部出栈,并添加索引
	for head >= 0 {
		table[stack[head]] = -1
		head--
	}
	for _, num := range nums1 {
		nextGreater, _ := table[num]
		res = append(res, nextGreater)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNextGreaterElementI(t *testing.T) {

}
