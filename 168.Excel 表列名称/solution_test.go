package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(columnNumber int) string {
	var res []rune
	for columnNumber != 0 {
		num := columnNumber % 26
		if num == 0 {
			num = 26
		}
		columnNumber = (columnNumber - num) / 26
		res = append(res, rune('A'+num-1))
	}
	// 逆转结果
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestExcelSheetColumnTitle(t *testing.T) {

}
