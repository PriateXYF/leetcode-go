package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) (res int) {
	head, tail := 0, len(height)-1
	res = (tail - head) * min(height[head], height[tail])
	for head != tail {
		if height[head] < height[tail] {
			head++
		} else {
			tail--
		}
		area := (tail - head) * min(height[head], height[tail])
		res = max(res, area)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContainerWithMostWater(t *testing.T) {

}
