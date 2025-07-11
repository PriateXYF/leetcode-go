package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maxCount(m int, n int, ops [][]int) int {
	minRow, minCol := m, n
	for _, op := range ops {
		row, col := op[0], op[1]
		minRow, minCol = min(minRow, row), min(minCol, col)
	}
	return minRow * minCol
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRangeAdditionIi(t *testing.T) {

}
