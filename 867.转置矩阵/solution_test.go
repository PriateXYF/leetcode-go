package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		line := make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			line[j] = matrix[j][i]
		}
		res[i] = line
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTransposeMatrix(t *testing.T) {

}
