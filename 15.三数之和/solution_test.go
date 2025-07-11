package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) (res [][]int) {
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return
	}
	head, tail := 0, len(nums)-1
	for head <= tail-2 && nums[head] <= 0 && nums[tail] >= 0 {
		diff := nums[tail] + nums[head] // 偏差值
		if diff > 0 {                   //正数偏大，需要两个负数抵消
			for temp := head + 1; temp < tail && nums[temp] <= 0 && nums[temp]+diff <= 0; temp++ {
				if nums[temp]+diff == 0 {
					res = append(res, []int{nums[head], nums[temp], nums[tail]})
					break
				}
			}
			for tail >= 1 && nums[tail-1] == nums[tail] {
				tail--
			}
			tail--
		} else { // 负数偏大，需要寻找正数
			for temp := tail - 1; temp > head && nums[temp] >= 0 && nums[temp]+diff >= 0; temp-- {
				if nums[temp]+diff == 0 {
					res = append(res, []int{nums[head], nums[temp], nums[tail]})
					break
				}
			}
			for head < len(nums)-1 && nums[head+1] == nums[head] {
				head++
			}
			head++
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSum(t *testing.T) {
	res := threeSum([]int{0, 0, 0, 0, 0})
	fmt.Println(res)
}
