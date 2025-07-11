package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	initColor := image[sr][sc]
	horizontal := []int{0, 0, 1, -1}
	vertical := []int{1, -1, 0, 0}
	var dfs func(int, int)
	dfs = func(row int, col int) {
		// 越界
		if row < 0 || row >= len(image) || col < 0 || col >= len(image[0]) {
			return
		}
		// 不为初始值
		if image[row][col] != initColor {
			return
		}
		// 如果节点已经访问过
		if image[row][col] == color {
			return
		}
		image[row][col] = color
		for i := 0; i < 4; i++ {
			dfs(row+horizontal[i], col+vertical[i])
		}
	}
	dfs(sr, sc)
	return image
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFloodFill(t *testing.T) {

}
