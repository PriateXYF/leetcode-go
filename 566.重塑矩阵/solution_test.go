package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	size := len(mat) * len(mat[0])
	if size != r*c {
		return mat
	}
	res := make([][]int, r)
	i, j := 0, 0
	line := make([]int, c)
	for _, row := range mat {
		for _, val := range row {
			line[j%c] = val
			if j++; j%c == 0 {
				res[i], line = line, make([]int, c)
				i++
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReshapeTheMatrix(t *testing.T) {

}
