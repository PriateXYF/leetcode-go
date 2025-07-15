package leetcode

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isValidSudoku(board [][]byte) bool {
	grids := make([][]int, 3)
	for i := range grids {
		grids[i] = make([]int, 3)
	}
	x, y := make([]int, 9), make([]int, 9)
	for i := range 9 {
		for j := range 9 {
			if board[i][j] == '.' {
				continue
			}
			board[i][j] = board[i][j] - '0'
			mask := 1 << board[i][j]
			// 判断行重复
			if x[i]&mask != 0 {
				return false
			}
			// 判断列重复
			if y[j]&mask != 0 {
				return false
			}
			// 判断格重复
			grid := grids[i/3][j/3]
			if grid&mask != 0 {
				return false
			}
			x[i] |= mask
			y[j] |= mask
			grids[i/3][j/3] |= mask
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidSudoku(t *testing.T) {
	board := [][]byte{{'.', '.', '.', '.', '5', '.', '.', '1', '.'}, {'.', '4', '.', '3', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '3', '.', '.', '1'}, {'8', '.', '.', '.', '.', '.', '.', '2', '.'}, {'.', '.', '2', '.', '7', '.', '.', '.', '.'}, {'.', '1', '5', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '2', '.', '.', '.'}, {'.', '2', '.', '9', '.', '.', '.', '.', '.'}, {'.', '.', '4', '.', '.', '.', '.', '.', '.'}}
	fmt.Println(isValidSudoku(board))
}
