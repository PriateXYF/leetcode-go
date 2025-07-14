package leetcode

import (
	"slices"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	for a := range n - 3 { // 枚举第一个数
		x := nums[a]
		if a > 0 && x == nums[a-1] { // 跳过重复数字
			continue
		}
		if x+nums[a+1]+nums[a+2]+nums[a+3] > target { // 优化一
			break
		}
		if x+nums[n-3]+nums[n-2]+nums[n-1] < target { // 优化二
			continue
		}
		for b := a + 1; b < n-2; b++ { // 枚举第二个数
			y := nums[b]
			if b > a+1 && y == nums[b-1] { // 跳过重复数字
				continue
			}
			if x+y+nums[b+1]+nums[b+2] > target { // 优化一
				break
			}
			if x+y+nums[n-2]+nums[n-1] < target { // 优化二
				continue
			}
			c, d := b+1, n-1
			for c < d { // 双指针枚举第三个数和第四个数
				s := x + y + nums[c] + nums[d] // 四数之和
				if s > target {
					d--
				} else if s < target {
					c++
				} else { // s == target
					ans = append(ans, []int{x, y, nums[c], nums[d]})
					for c++; c < d && nums[c] == nums[c-1]; c++ {
					} // 跳过重复数字
					for d--; d > c && nums[d] == nums[d+1]; d-- {
					} // 跳过重复数字
				}
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFourSum(t *testing.T) {

}
