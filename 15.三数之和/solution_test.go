package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		target := -nums[first]
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}

		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSum(t *testing.T) {
	//res := threeSum([]int{0, 0, 0, 0, 0})
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}
