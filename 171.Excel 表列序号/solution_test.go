package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func pow(num int, p int) (res int) {
	res = 1
	for range p {
		res *= num
	}
	return res
}

// 计算 Excel 列名称对应的序号
func titleToNumber(columnTitle string) (res int) {
	digit := len(columnTitle)
	for _, col := range columnTitle {
		digit--
		res += int(col-'A'+1) * pow(26, digit)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestExcelSheetColumnNumber(t *testing.T) {

}
