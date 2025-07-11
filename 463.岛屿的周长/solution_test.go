package leetcode

import (
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)

func isOcean(grid [][]int, row int, col int) int {
	if row >= len(grid) || row < 0 {
		return 1
	} else if col >= len(grid[0]) || col < 0 {
		return 1
	} else if grid[row][col] == 0 {
		return 1
	} else {
		return 0
	}
}

// 计算岛屿周长
func islandPerimeter(grid [][]int) (res int) {
	for row, rows := range grid {
		for col, val := range rows {
			if val == 1 {
				res += isOcean(grid, row-1, col)
				res += isOcean(grid, row, col-1)
				res += isOcean(grid, row+1, col)
				res += isOcean(grid, row, col+1)
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIslandPerimeter(t *testing.T) {

}
