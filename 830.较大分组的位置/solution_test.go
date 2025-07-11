package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func largeGroupPositions(s string) (res [][]int) {
	head, tail := 0, 0
	for head < len(s) {
		for ; tail < len(s) && s[tail] == s[head]; tail++ {
		}
		if tail-head >= 3 {
			res = append(res, []int{head, tail - 1})
		}
		head = tail
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPositionsOfLargeGroups(t *testing.T) {

}
