package leetcode

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) (res [][]int) {
	var dfs func(int, int)
	var temp []int
	// 选择指定位数
	dfs = func(i int, t int) {
		if i >= len(candidates) || t < 0 { // 超出迭代范围
			return
		}
		if t == 0 { // 达到目标条件
			s := make([]int, len(temp))
			copy(s, temp)
			res = append(res, s)
			return
		}
		// 选择
		temp = append(temp, candidates[i])
		dfs(i, t-candidates[i])
		if len(temp) > 0 {
			temp = temp[:len(temp)-1]
		}
		// 不选择
		dfs(i+1, t)
	}
	dfs(0, target)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCombinationSum(t *testing.T) {
	res := combinationSum([]int{2, 3, 5}, 8)
	fmt.Println(res)
}
