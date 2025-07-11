package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 获取指定行列的周围九个单元格
func getAroundImgs(row int, col int, width int, height int) (res [][]int) {
	positions := [][]int{
		[]int{row - 1, col - 1},
		[]int{row - 1, col},
		[]int{row - 1, col + 1},
		[]int{row, col - 1},
		[]int{row, col},
		[]int{row, col + 1},
		[]int{row + 1, col - 1},
		[]int{row + 1, col},
		[]int{row + 1, col + 1},
	}
	for _, pos := range positions {
		i, j := pos[0], pos[1]
		if i >= 0 && i < height && j >= 0 && j < width {
			res = append(res, pos)
		}
	}
	return
}

// 图片平滑器
func imageSmoother(img [][]int) [][]int {
	width, height := len(img[0]), len(img)
	res := make([][]int, height)
	for i := range res {
		res[i] = make([]int, width)
	}
	for row, cols := range img {
		for col, _ := range cols {
			positions := getAroundImgs(row, col, width, height)
			sum := 0
			for _, pos := range positions {
				sum += img[pos[0]][pos[1]]
			}
			res[row][col] = sum / len(positions)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestImageSmoother(t *testing.T) {

}
