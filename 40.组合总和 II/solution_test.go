package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	var dfs func(int, int)
	var temp []int
	dfs = func(i, t int) {
		if t == 0 {
			res = append(res, append([]int(nil), temp...))
			return
		}
		if i >= len(candidates) || t < 0 { // 到达边界条件
			return
		}
		// 小等于目标时才选择
		if candidates[i] <= target {
			temp = append(temp, candidates[i])
			dfs(i+1, t-candidates[i])
			temp = temp[:len(temp)-1]
		} else { // 大于时直接返回，因为之后都不需要选择
			return
		}
		// 不选择
		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
		dfs(i+1, t)
	}
	dfs(0, target)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCombinationSumIi(t *testing.T) {
	res := combinationSum2([]int{29, 19, 14, 33, 11, 5, 9, 23, 23, 33, 12, 9, 25, 25, 12, 21, 14, 11, 20, 30, 17, 19, 5, 6, 6, 5, 5, 11, 12, 25, 31, 28, 31, 33, 27, 7, 33, 31, 17, 13, 21, 24, 17, 12, 6, 16, 20, 16, 22, 5}, 28)
	fmt.Println(res)
}
