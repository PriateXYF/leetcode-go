package leetcode

import (
	"fmt"
	"slices"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var a, b int
	// 找到最后一个顺序排列的
	for a = len(nums) - 2; a >= 0 && nums[a] >= nums[a+1]; a-- {
	}
	if a >= 0 {
		// 找到最小的大于 a 的数字
		for b = len(nums) - 1; nums[b] <= nums[a]; b-- {
		}
		// 交换 a 和 b
		nums[a], nums[b] = nums[b], nums[a]
	}
	// 反转 a 之后的元素
	slices.Reverse(nums[a+1:])
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNextPermutation(t *testing.T) {
	arr := []int{1, 5, 1}
	nextPermutation(arr)
	fmt.Println(arr)
}
