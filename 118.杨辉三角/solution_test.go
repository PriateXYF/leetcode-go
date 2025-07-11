package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 通过上一列生成杨辉三角的某一列
func generateRow(prevRow []int) (row []int) {
	if len(prevRow) == 1 {
		return []int{1, 1}
	}
	row = append(row, 1)
	for i := 0; i < len(prevRow)-1; i++ {
		row = append(row, prevRow[i]+prevRow[i+1])
	}
	row = append(row, 1)
	return row
}

// 生成杨辉三角
func generate(numRows int) (res [][]int) {
	for i := range numRows {
		if i == 0 {
			res = append(res, []int{1})
			continue
		}
		row := generateRow(res[i-1])
		res = append(res, row)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPascalsTriangle(t *testing.T) {

}
