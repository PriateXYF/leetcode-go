package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 通过上一列生成杨辉三角的某一列
func generateRowII(prevRow []int) (row []int) {
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

// 获取杨辉三角的指定行
func getRow(rowIndex int) (res []int) {
	res = []int{1}
	for range rowIndex {
		res = generateRowII(res)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPascalsTriangleIi(t *testing.T) {

}
