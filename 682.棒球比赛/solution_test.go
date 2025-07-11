package leetcode

import (
	"strconv"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func calPoints(operations []string) (res int) {
	stack := make([]int, len(operations))
	point := 0
	for _, operation := range operations {
		switch operation {
		case "C":
			point--
		case "D":
			stack[point] = stack[point-1] * 2
			point++
		case "+":
			stack[point] = stack[point-1] + stack[point-2]
			point++
		default:
			stack[point], _ = strconv.Atoi(operation)
			point++
		}
	}
	for i := 0; i < point; i++ {
		res += stack[i]
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBaseballGame(t *testing.T) {

}
