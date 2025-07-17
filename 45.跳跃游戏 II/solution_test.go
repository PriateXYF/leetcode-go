package leetcode

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	times, index := 1, 0
	for index+nums[index] < len(nums)-1 { // 判断下两跳的最大值
		maxTwice, best := 0, 0
		for i := 0; i <= nums[index]; i++ {
			j := index + i
			once := nums[j]
			twice := once + j
			if twice > maxTwice {
				maxTwice, best = twice, i
			}
		}
		index += best
		times++
	}
	return times
}

//leetcode submit region end(Prohibit modification and deletion)

func TestJumpGameIi(t *testing.T) {
	res := jump([]int{1, 1, 1, 1})
	fmt.Println(res)
}
