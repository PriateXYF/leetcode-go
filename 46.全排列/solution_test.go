package leetcode

import (
    "testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) (res [][]int) {
	var dfs func(int)
	dfs = func(n int) {
		if n == len(nums) { // 终止条件
			res = append(res, append([]int(nil), nums...))
			return
		}
		for i := n; i < len(nums); i++ {
			nums[n], nums[i] = nums[i], nums[n] // 固定住前 i 项
			dfs(n + 1)
			nums[n], nums[i] = nums[i], nums[n] // 恢复现场
		}
	}
	dfs(0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPermutations(t *testing.T) {
	permute([]int{1, 2, 3, 4})
}
