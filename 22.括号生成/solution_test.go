package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) (res []string) {
	var dfs func(int, int)
	temp := make([]uint8, n*2)
	dfs = func(i int, left int) {
		// 达到边界
		if i == len(temp) {
			res = append(res, string(temp))
		}
		// 先填充左括号
		if left < n {
			temp[i] = '('
			dfs(i+1, left+1)
		}
		if right := i - left; right < left {
			temp[i] = ')'
			dfs(i+1, left)
		}
	}
	dfs(0, 0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGenerateParentheses(t *testing.T) {
	generateParenthesis(3)
}
